package events

import (
	"testing"
	"time"

	"github.com/shelbyramsey/golang_interfaces/pkg/dynamo"
)

// TestErrBuff makes sure that if an error happens from Dynamo that
// the data is not lost and is in fact in the Queue.buff to be processed
// at a later time
//
func TestErrBuff(t *testing.T) {
	Dynamo = &dynamo.Dynamo{}
	Queue = NewQueue()
	ev := &Event{
		Timestamp: time.Now().Unix(),
		Data: []*Data{
			&Data{
				Key: "uuid-1",
				Val: "some-val",
			},
		},
	}
	Queue.Rcv <- ev
	ev = &Event{
		Timestamp: time.Now().Unix(),
		Data: []*Data{
			&Data{
				Key: "uuid-2",
				Val: "some-val",
			},
		},
	}
	Queue.Rcv <- ev
	// let chans clear
	time.Sleep(10 * time.Millisecond)
	if len(Queue.buff) != 2 {
		t.Fail()
	}
}
