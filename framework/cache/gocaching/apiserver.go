package gocaching

import (
	"context"
	"io/ioutil"
	"net/http"
	"time"

	"github.com/galaxyzeta/concurrency/singleflight"
	"github.com/galaxyzeta/consistenthash"
	"github.com/galaxyzeta/logger"
	"github.com/galaxyzeta/util/common"
)

// HandlerFunc works as the handler for http request.
type HandlerFunc func(http.ResponseWriter, *http.Request)

// APIServer provides general front-end web api.
type APIServer struct {
	host          string
	c             *consistenthash.ConsistentHash
	peers         map[string]struct{}
	inactivePeers map[string]struct{}
	getter        LocalGetter
	router        map[string]HandlerFunc
	logger        *logger.Logger
	hbTicker      *time.Ticker // HeartBeatTicker for handling heart-beat send event.
	sf            *singleflight.SingleFlight
}

var heartBeatInterval time.Duration = time.Second * 5

// RunAPIServer boots up the API front-end server.
func (g *GoCaching) RunAPIServer(host string, peers ...string) {
	api := &APIServer{
		host:          host,
		router:        make(map[string]HandlerFunc),
		c:             consistenthash.NewConsistentHash(),
		logger:        logger.New("API Server"),
		hbTicker:      time.NewTicker(heartBeatInterval),
		peers:         make(map[string]struct{}),
		inactivePeers: make(map[string]struct{}),
		sf:            singleflight.NewSingleFlight(),
	}
	// Router register.
	api.router["/getkey"] = api.getFromRemote
	// Peers register.
	
	for _, peer := range peers {
		api.peers[peer] = struct{}{}
		api.c.Add(peer)
	}
	// Start a heart beat chrono job.
	go api.heartBeatRoutine()
	http.ListenAndServe(api.host, api)
}

// heartBeatRoutine sends heart beat to all caching server in certain frequency.
func (g *APIServer) heartBeatRoutine() {
	for {
		<-g.hbTicker.C
		g.logger.Debugf("Sending heart-beat to all servers.")
		for peer := range g.peers {
			go g.sendHeartBeat(peer, true)
		}
		// Probing inactive.
		for peer := range g.inactivePeers {
			go g.sendHeartBeat(peer, false)
		}
	}
}

// sendHeartBeat to a remote node.
func (g *APIServer) sendHeartBeat(peer string, isActive bool) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	go func() {
		_, err := http.Get(common.BuildString("http://", peer, heartBeatRoute))
		if err != nil {
			if isActive == true {
				g.logger.Errorf("An error occured while probing heart-beat to remote node %v: %v", peer, err.Error())
				return
			}
			g.logger.Errorf("No response from DEAD node %v", peer)

		} else if isActive == false {
			// Inactive node revive.
			g.logger.Warnf("Re-activating node %v", peer)
			g.peerActivate(peer)
		}
		cancel()
	}()

	<-ctx.Done() // Wait a job to be done (cancel manually), or DLE.
	if ctx.Err() == context.DeadlineExceeded {
		if isActive {
			// If isActive and Timeout, call peerTimeout to move the active peer to inactive set.
			g.logger.Warnf("Lost connection to remote server %s", peer)
			g.peerInactivate(peer)
		}
	}
}

// peerInactivate moves a peer to inactive set.
func (g *APIServer) peerInactivate(peer string) {
	delete(g.peers, peer)
	g.inactivePeers[peer] = struct{}{}
	g.c.Remove(peer)
}

// peerActivate moves a peer from inactive set to active.
func (g *APIServer) peerActivate(peer string) {
	delete(g.inactivePeers, peer)
	g.peers[peer] = struct{}{}
	g.c.Add(peer)
}

// ServeHTTP implements http.Handler.
func (g *APIServer) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	handler := g.router[r.URL.Path]
	if handler != nil {
		handler(w, r)
	} else {
		write404(w, "Handler not found !")
	}
}

// getFromRemote retrieves key from remote.
func (g *APIServer) getFromRemote(w http.ResponseWriter, r *http.Request) {
	queries := r.URL.Query()
	key := queries.Get("key")
	if key == "" {
		write404(w, "Do not have [KEY] query! ")
		return
	}
	// Retrieve consistent hash peer.
	remoteAddr := g.c.Get(key)
	g.logger.Debugf("Received Key %v, RemoteAddr should be %v.", key, remoteAddr)
	url := "http://" + remoteAddr + "/" + key

	// Prevent violent, flooding requests by using singleflight.
	res, err := g.sf.Do(url, func() (interface{}, error) {
		g.logger.Debugf("Sending request to %s", url)
		resp, err := http.Get(url)
		defer resp.Body.Close()
		if err != nil {
			g.logger.Fatalf(err.Error())
			return nil, err
		}
		res, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			g.logger.Fatalf(err.Error())
			return nil, err
		}
		return res, err
	})
	if err != nil {
		write404(w, err.Error())
		return
	}
	write200(w, string(res.([]byte)))
}
