// Author: @ Galaxyzeta
// BioServer is a simple implementation of HttpServer. Including:
// 1. HttpRequest parsing.
// 2. HashMap router.
// 3. HttpResponse parsing.

package galaxyserver

import (
	"fmt"
	"net"
	"reflect"
	"strconv"
	"strings"
)

var maxBufferSize int32 = 1024

var cutset = " \r\n"
var methodList = []string{"GET", "POST"}

// Router : Type alias for router
type Router = map[string]map[string]func(*Context) interface{}

// GalaxyServer : Abstract representation of my server.
type GalaxyServer struct {
	ip     string
	port   int
	router Router
}

// Context : Abstraction of a http session.
type Context struct {
	request    *HTTPRequest
	response   *HTTPResponse
	viewObject interface{}
}

// HTTPRequest : Inbound request.
type HTTPRequest struct {
	url         string
	version     string
	method      string
	headers     map[string]string
	requestBody string
}

// HTTPResponse : Outbound response.
type HTTPResponse struct {
	code         int
	status       string
	version      string
	headers      map[string]string
	responseBody string
}

var mockResp string = "http/1.1 200 OK\r\nToken: 123456\r\nContent-Length: 12\r\n\r\nthis is body\r\n\r\n"

// New : Create a new GalaxyServer
func New() GalaxyServer {
	ret := GalaxyServer{router: Router{}}
	for _, v := range methodList {
		ret.router[v] = map[string]func(*Context) interface{}{}
	}
	return ret
}

// Start : start my server
func (server *GalaxyServer) Start(ip string, port int) {
	server.ip = ip
	server.port = port
	serve(server, buildAddress(ip, port))
}

// RespNotFound : Get a not found 404 response
func RespNotFound() *HTTPResponse {
	resp := RespOK()
	resp.status = "NOT FOUND"
	resp.code = 404
	return resp
}

// RespInternalServerError : Get an Internal Server Error 500 response
func RespInternalServerError() *HTTPResponse {
	resp := RespOK()
	resp.status = "INTERNAL SERVER ERROR"
	resp.code = 500
	return resp
}

// RespOK : Get a 200 OK blank response
func RespOK() *HTTPResponse {
	ret := &HTTPResponse{
		code:         200,
		status:       "OK",
		version:      "http/1.1",
		headers:      map[string]string{},
		responseBody: "",
	}
	ret.headers["Server"] = "Galaxy Server"
	ret.headers["Content-Length"] = "0"
	return ret
}

// ParseRequest : convert bare string into http request.
func ParseRequest(str string) *HTTPRequest {

	defer panicHandler()

	var request *HTTPRequest = &HTTPRequest{headers: map[string]string{}}
	lines := strings.Split(str, "\n")
	httpHeads := strings.Split(lines[0], " ")
	if len(httpHeads) != 3 {
		return (*HTTPRequest)(nil)
	}
	request.method = httpHeads[0]
	request.url = httpHeads[1]
	request.version = httpHeads[2]

	index := 1

	// Handle Http Headers
	for ; index < len(lines); index++ {
		if strings.Trim(lines[index], cutset) != "" {
			kv := strings.Split(lines[index], ":")
			request.headers[strings.Trim(kv[0], " ")] = strings.Trim(kv[1], " ")
		} else {
			index++
			break
		}
	}

	// Handle request body
	lines = lines[index:] // re-cut
	request.requestBody = strings.Join(lines, "\n")

	return request
}

// StructToString : Println from struct to stdout.
func StructToString(obj interface{}) string {
	refval := reflect.ValueOf(obj)
	reftype := reflect.TypeOf(obj)
	kv := make(map[string]interface{})
	for i := 0; i < refval.NumField(); i++ {
		kv[reftype.Field(i).Name] = refval.Field(i)
	}
	var str string = ""
	for k, v := range kv {
		str += fmt.Sprintf("%s: %s\n", k, v)
	}
	return str
}

// ToString : Convert HTTPResponse to string.
func (resp *HTTPResponse) ToString() string {
	str := strings.Join([]string{resp.version, strconv.Itoa(resp.code), resp.status}, " ")
	str += "\r\n"
	for k, v := range resp.headers {
		str += fmt.Sprintf("%s: %s\r\n", k, v)
	}
	str += "\r\n"
	str += resp.responseBody
	str += "\r\n"
	return str
}

// ==== Server Registration ====

// GET : Register a get mapping to the register.
func (server *GalaxyServer) GET(url string, method func(*Context) interface{}) {
	server.router["GET"][url] = method
}

// POST : Register a get mapping to the register.
func (server *GalaxyServer) POST(url string, method func(*Context) interface{}) {
	server.router["POST"][url] = method
}

func buildAddress(ip string, port int) string {
	return ip + ":" + strconv.Itoa(port)
}

func serve(server *GalaxyServer, address string) {
	tcphandler, err := net.Listen("tcp", address)
	// defer tcphandler.Close()
	if err != nil {
		panic(err)
	}
	// Endless server loop
	for {
		socket, err := tcphandler.Accept()
		if err != nil {
			panic(err)
		}
		fmt.Printf("Received connection from %s\n", socket.RemoteAddr().String())
		go httpHandler(server, socket)
	}
}

func httpHandler(server *GalaxyServer, socket net.Conn) {
	// Prevent Resource Leak
	defer func() {
		socket.Close()
		fmt.Print("Connection closed\n")
	}()
	buf := make([]byte, maxBufferSize)
	// Prepare context
	var session *Context = &Context{}
	for {
		var tmp int = -1
		var err error
		var str string = ""
		if tmp != 0 {
			tmp, err = socket.Read(buf)
			if err != nil {
				fmt.Printf("Error : %s\n", err)
				return
			}
			str += string(buf)
		} else {
			break
		}
		// Parse Http Request
		session.request = ParseRequest(str)
		if session.request == nil {
			fmt.Println("Not a valid http request")
		}
		fmt.Printf("Received message: %s\n", StructToString(*session.request))

		// Handler Process
		method := session.request.method
		url := session.request.url
		function := server.router[method][url]
		if function == nil {
			session.response = RespNotFound()
		} else {
			session.response = RespOK()
			session.viewObject = function(session)
			// test--
			session.response.responseBody = session.viewObject.(string)
			// end test--
			session.response.headers["Content-Length"] = strconv.Itoa(len(session.response.responseBody))
		}

		// Send request
		fmt.Println(session.response.ToString())
		socket.Write([]byte(session.response.ToString()))
	}
}

func panicHandler() interface{} {
	err := recover()
	if err != nil {
		fmt.Println(err)
		return err
	}
	return nil
}
