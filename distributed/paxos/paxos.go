// Package Paxos is used to provide consensus algorithm for distributed system.

package paxos

import (
	"context"
	"errors"
	"time"

	mq "github.com/galaxyzeta/concurrency/messaging"
)

// Proposer represents the character who propose a decision to be discussed.
type Proposer struct {
	label             string
	acceptors         []Acceptor
	network           IPaxosNetwork
	proposalID        int
	timeout           time.Duration
	retry             int
	maxretry          int
	seenProposalValue string
}

// Acceptor accepts decisions and decide to accept or reject a proposal.
type Acceptor struct {
	recentProposalID int
	proposalValue    string
	label            string
	network          IPaxosNetwork
}

type status int8

const (
	ack  status = iota
	nack        = iota
)

// ErrInconsistentProposal indicates that the proposer has received different seenProposals from acceptors.
var ErrInconsistentProposal error = errors.New("inconsistent proposal value received from acceptors")

type message struct {
	messageID  int
	from       string
	to         string
	statusCode status
	body       string
}

// IPaxosNetwork is an abstraction for transmission between endpoints.
type IPaxosNetwork interface {
	SendMessage(msg message)
	ReadMessage(label string) (message, bool)
	ReadMessageBlocking(label string) message
}

type paxosNetwork struct {
	mq mq.MQ
}

func (n *paxosNetwork) SendMessage(msg message) {
	n.mq.Send(msg.to, msg)
}

func (n *paxosNetwork) ReadMessage(label string) (message, bool) {
	ret, success := n.mq.Receive(label)
	return ret.(message), success
}

func (n *paxosNetwork) ReadMessageBlocking(label string) message {
	return n.mq.ReceiveBlocking(label).(message)
}

// Reset proposer to its most initial status.
func (p *Proposer) Reset() {
	p.retry = p.maxretry
}

// AddAcceptors to the proposer.
func (p *Proposer) AddAcceptors(acceptors ...Acceptor) {
	p.acceptors = append(p.acceptors, acceptors...)
}

// Prepare sends a preparation to all acceptors, blocks until half acceptances are reached.
func (p *Proposer) Prepare() error {
	p.Reset()
	msg := message{from: p.label}
	for ; p.retry > 0; p.retry-- {
		// Send message to all acceptors
		acceptorNumbers := len(p.acceptors)
		msg.messageID = p.proposalID
		for _, acceptor := range p.acceptors {
			msg.to = acceptor.label
			p.network.SendMessage(msg)
		}
		// Try to read message from channel.
		timeoutCtx, cancel := context.WithTimeout(context.Background(), p.timeout)
		receivedCount := 0
		for {
			select {
			// If already done.
			case <-timeoutCtx.Done():
				break
			}
			// If not done, try to read from all acceptors
			for i := 0; i < len(p.acceptors); i++ {
				if msg, success := p.network.ReadMessage(p.label); success == true {
					if msg.statusCode == ack {
						receivedCount++
						if p.seenProposalValue == "" {
							p.seenProposalValue = msg.body
						} else if p.seenProposalValue != msg.body {
							cancel()
							return ErrInconsistentProposal
						}
					}
				}
			}
			// Already received half acknowledgement, cancel timeout and break.
			if receivedCount >= acceptorNumbers/2 {
				cancel()
				break
			}
		}
		// Check timeout.
		if err := timeoutCtx.Err(); err != context.Canceled {
			break
		}
		// Else, failed to retrieve half accepts. Start another round.
	}
	return nil
}

func (p *Proposer) Propose(val string) {

}

// AcceptPrepare wait for preparations to be received, and send back a response to message sender.
func (a *Acceptor) AcceptPrepare() {
	msg := a.network.ReadMessageBlocking(a.label)
	msg.body = ""
	if msg.messageID > a.recentProposalID {
		msg.statusCode = ack
		msg.body = a.proposalValue
	} else {
		msg.statusCode = nack
	}
	// Send back response
	msg.to = msg.from
	msg.from = a.label
	a.network.SendMessage(msg)
}
