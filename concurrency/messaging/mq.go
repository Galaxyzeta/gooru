package mq

// MQ is an abbr for message queue, which is a map with key string and a generic channel as value.
type MQ struct {
	mailBox map[string]chan interface{}
}

// NewMQ returns a new MQ.
func NewMQ() MQ {
	return NewMQ()
}

// Send message to message queue.
func (mq MQ) Send(label string, msg interface{}) {
	mq.mailBox[label] <- msg
}

// ReceiveBlocking blocks until it received message from a labelled message queue.
func (mq MQ) ReceiveBlocking(label string) interface{} {
	return <-mq.mailBox[label]
}

// Receive receives message from a labelled message queue. Will return nil if it has not received message yet.
func (mq MQ) Receive(label string) (interface{}, bool) {
	select {
	case ret := <-mq.mailBox[label]:
		return ret, true
	default:
		return nil, false
	}
}

// ReadonlyChannel returns a channel that can only be read.
func (mq MQ) ReadonlyChannel(label string) <-chan interface{} {
	return mq.mailBox[label]
}
