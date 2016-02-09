package events

import (
	"bytes"
	"encoding/json"
	"log"
	"strconv"
	"time"

	"github.com/aws/aws-sdk-go/service/dynamodb"
	"github.com/aws/aws-sdk-go/service/dynamodb/dynamodbiface"
	"github.com/satori/go.uuid"
	"golang.org/x/net/websocket"
)

var (
	Queue       *queue
	Dynamo      dynamodbiface.DynamoDBAPI
	DynamoTable string
)

// Data holds a key:value pair
type Data struct {
	Key string `json:"key"`
	Val string `json:"val"`
}

// Event holds a []*Data and a Timestamp
type Event struct {
	Timestamp int64
	Data      []*Data
}

// String returns the string representation of the Event or an error
func (ev *Event) String() (str string, err error) {
	b, err := json.Marshal(ev)
	if err != nil {
		return "", err
	}
	return bytes.NewBuffer(b).String(), nil
}

// GetEvent returns an Event or an error from a []byte
func GetEvent(b []byte) (ev *Event, err error) {
	ev = &Event{}
	if err = json.Unmarshal(b, ev); err != nil {
		return nil, err
	}
	return ev, nil
}

// WebsocketHandler is a simple handler that reads messages sent
// and in the event that an error occurs closes the websocket connection
func WebsocketHandler(ws *websocket.Conn) {
	defer ws.Close()
	var n int
	var err error
	var ev *Event
	for {
		msg := make([]byte, 512)
		n, err = ws.Read(msg)
		// if err occurs just close websocket
		if err != nil {
			log.Println("WebsocketHandler: got err reading:", err)
			break
		}
		// if err occurs just close websocket
		ev, err = GetEvent(msg[0:n])
		if err != nil {
			log.Println("WebsocketHandler: got err getting event:", err)
			break
		}
		Queue.Rcv <- ev
	}
}

// queue is a buffer to put data into Dynamo.  New events should be sent
// to the Rcv channel and buff can be used to ensure that if any errors
// occur while sending to Dynamo that data will not get lost.  This is
// helpful to deal with intermittent, short term network problems (common
// on AWS and other cloud services).
//
type queue struct {
	Rcv  chan *Event
	buff chan *dynamodb.PutItemInput
	t    *time.Timer
}

// Run clears the Rcv channel
func (q *queue) Run() {
	for {
		ev, ok := <-q.Rcv
		if !ok {
			panic("queue.Run: Rcv chan is closed")
		}
		q.put(ev)
	}
}

// put places events into Dynamo
func (q *queue) put(ev *Event) {
	ts := strconv.FormatInt(ev.Timestamp, 10)
	id := uuid.NewV4().String()
	str, err := ev.String()
	if err != nil {
		panic("queue.put: got err getting event string")
	}
	item := map[string]*dynamodb.AttributeValue{
		"id":   {S: &id},
		"ts":   {N: &ts},
		"data": {S: &str},
	}
	input := &dynamodb.PutItemInput{
		Item:      item,
		TableName: &DynamoTable,
	}
	_, err = Dynamo.PutItem(input)
	if err != nil {
		log.Println("queue.put: got err from dynamo:", err)
		q.buff <- input
	}
}

// chk clears the buff or breaks if an error occurs
func (q *queue) chk() {
	var err error
	for {
		input, ok := <-q.buff
		if !ok {
			panic("queue.chk: buff chan is closed")
		}
		_, err = Dynamo.PutItem(input)
		if err != nil {
			q.buff <- input
			break
		}
	}
}

// runChk listens to the timer and executes chk
func (q *queue) runChk() {
	for {
		<-q.t.C
		q.chk()
		q.t.Reset(60 * time.Second)
	}
}

// NewQueue returns a queue and starts the appropriate services
// on the queue (Run and runChk)
func NewQueue() *queue {
	q := &queue{
		Rcv:  make(chan *Event, 100),
		buff: make(chan *dynamodb.PutItemInput, 200),
		t:    time.NewTimer(60 * time.Second),
	}
	go q.Run()
	go q.runChk()
	return q
}
