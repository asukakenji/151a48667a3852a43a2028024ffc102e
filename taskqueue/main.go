// Tier 1:
// - Handle Route:
//   - USE "default"
//   - PUT
//     - body: token + locations
//     - pri: timestamp
//     - delay: 0 seconds
//     - ttr: 3 seconds
// - Handle Route Token:
//   - USE "<token>"
//   - RESERVE
//   - Return result

// Tier 2:
// - Handle Task:
//   - USE "default"
//   - RESERVE
//   - USE "<token>"
//   - Loop
//     - RESERVE
//     - DELETE
//       - id: id of job just reserved
//   - Depending on the result:
//     - Not Found:
//       - PUT
//         - body: "in progress" + trialCount (= 1)
//         - pri: (current) timestamp
//         - delay: 0 seconds
//         - ttr: 0 seconds
//     - trialCount != maxTrialCount
//       - PUT
//         - body: "in progress" + trialCount (+= 1)
//         - pri: (current) timestamp
//         - delay: 0 seconds
//         - ttr: 0 seconds
//     - trialCount == maxTrialCount
//       - PUT
//         - body: "failure" + trialCount (= maxTrialCount)
//         - pri: (current) timestamp
//         - delay: 0 seconds
//         - ttr: 0 seconds
//       - Return
//   - Google Maps
//     - Use goroutine to parallel the queries
//     - Or see whether the API allows many questions in one trip
//   - Travelling Salesman
//   - PUT
//     - body: "success" + path + other results
//     - pri: (current) timestamp
//     - delay: 0 seconds
//     - ttr: 0 seconds
//   - USE "garbage"
//   - PUT
//     - body: token
//     - pri: (current) timestamp
//     - delay: 600 seconds
//     - ttr: 0 seconds
//   - USE "default"
//   - DELETE
//     - id: id of the job

// Tier C:
// - Loop
//   - USE "garbage"
//   - RESERVE
//   - USE "<token>"
//     - Loop
//       - RESERVE
//       - DELETE
//         - id: id of the job just reserved

package beanstalk

import (
	"bytes"
	"encoding/json"
	"math"
	"time"

	"github.com/asukakenji/151a48667a3852a43a2028024ffc102e/common"
	"github.com/golang/glog"
	"github.com/kr/beanstalk"
)

const (
	// TimeForever represents a very large duration.
	TimeForever = math.MaxUint32 * time.Second
)

func WithConnection(addr string, do func(*beanstalk.Conn) error) error {
	conn, err := beanstalk.Dial("tcp", addr)
	if err != nil {
		glog.Errorf("WithConnection: Dial error: %#v", err)
		return err
	}
	defer conn.Close()
	return do(conn)
}

func AddQuestion(conn *beanstalk.Conn, token string, locations common.Locations, executionTimeLimit time.Duration) (id uint64, err error) {
	buf := new(bytes.Buffer)
	err = json.NewEncoder(buf).Encode(common.Question{
		Token:     token,
		Locations: locations,
	})
	if err != nil {
		glog.Errorf("AddQuestion: Encode JSON: %#v", err)
		return 0, err
	}

	id, err = conn.Put(
		buf.Bytes(),               // body
		uint32(time.Now().Unix()), // pri
		time.Duration(0),          // delay
		executionTimeLimit,        // ttr
	)
	if err != nil {
		if cerr, ok := err.(beanstalk.ConnError); !ok {
			glog.Errorf("AddQuestion: Non-ConnError: %#v", err)
			return 0, err
		} else if cerr.Err == beanstalk.ErrBuried {
			glog.Errorf("AddQuestion: Buried: %#v", err)
			return 0, err
		} else if cerr.Err == beanstalk.ErrNoCRLF {
			glog.Errorf("AddQuestion: Expected CRLF: %#v", err)
			return 0, err
		} else if cerr.Err == beanstalk.ErrJobTooBig {
			glog.Errorf("AddQuestion: Job too big: %#v", err)
			return 0, err
		} else if cerr.Err == beanstalk.ErrDraining {
			glog.Errorf("AddQuestion: Draining: %#v", err)
			return 0, err
		}
		glog.Errorf("AddQuestion: Unknown error: %#v", err)
		return 0, err
	}
	glog.Infof("AddQuestion: token: %q, id: %d", token, id)
	return id, nil
}

func FetchQuestion(conn *beanstalk.Conn, timeout time.Duration) (id uint64, q *common.Question, err error) {
	for {
		var body []byte
		id, body, err = conn.Reserve(
			timeout, // timeout
		)
		if err != nil {
			if cerr, ok := err.(beanstalk.ConnError); !ok {
				glog.Errorf("FetchQuestion: Non-ConnError: %#v", err)
				return 0, nil, err
			} else if cerr.Err == beanstalk.ErrTimeout {
				glog.Infof("FetchQuestion: Timeout")
				continue
			} else if cerr.Err == beanstalk.ErrDeadline {
				glog.Infof("FetchQuestion: Deadline soon")
				time.Sleep(1 * time.Second)
				continue
			}
			glog.Errorf("FetchQuestion: Unknown error: %#v", err)
			return 0, nil, err
		}
		glog.Infof("FetchQuestion: id: %d", id)

		err = json.NewDecoder(bytes.NewReader(body)).Decode(q)
		if err != nil {
			glog.Errorf("FetchQuestion: Decode JSON: %#v", err)
			return 0, nil, err
		}
		return id, q, err
	}
}

func ClearAnswer(conn *beanstalk.Conn, token string) (err error) {
	tube := beanstalk.Tube{
		Conn: conn,
		Name: token,
	}
	for {
		var id uint64
		id, _, err = tube.PeekReady()
		if err != nil {
			if cerr, ok := err.(beanstalk.ConnError); !ok {
				glog.Errorf("ClearAnswer: Non-ConnError: %#v", err)
				return err
			} else if cerr.Err == beanstalk.ErrNotFound {
				return nil
			}
			glog.Errorf("ClearAnswer: Unknown error: %#v", err)
			return err
		}
		glog.Infof("ClearAnswer: Peek Ready: id: %d", id)

		err = conn.Delete(
			id, // id
		)
		if err != nil {
			if cerr, ok := err.(beanstalk.ConnError); !ok {
				glog.Errorf("ClearAnswer: Non-ConnError: %#v", err)
				return err
			} else if cerr.Err == beanstalk.ErrNotFound {
				glog.Errorf("ClearAnswer: Not found")
				return err
			}
			glog.Errorf("ClearAnswer: Unknown error: %#v", err)
			return err
		}
		glog.Infof("ClearAnswer: Deleted: id: %d", id)
	}
}

func SetAnswer(conn *beanstalk.Conn, token string, dr *common.DrivingRoute, pri uint32, delay, ttr time.Duration) (id uint64, err error) {
	err = ClearAnswer(conn, token)
	if err != nil {
		glog.Errorf("SetAnswer: Error from ClearAnswer: %#v", err)
		return 0, err
	}

	buf := new(bytes.Buffer)
	err = json.NewEncoder(buf).Encode(common.Answer{
		Timestamp:    time.Now().Unix(),
		DrivingRoute: dr,
	})
	if err != nil {
		glog.Errorf("SetAnswer: Encode JSON: %#v", err)
		return 0, err
	}

	tube := beanstalk.Tube{
		Conn: conn,
		Name: token,
	}
	id, err = tube.Put(
		buf.Bytes(),               // body
		uint32(time.Now().Unix()), // pri
		time.Duration(0),          // delay
		5*time.Second,             // ttr
	)
	if err != nil {
		if cerr, ok := err.(beanstalk.ConnError); !ok {
			glog.Errorf("SetAnswer: Non-ConnError: %#v", err)
			return 0, err
		} else if cerr.Err == beanstalk.ErrBuried {
			glog.Errorf("SetAnswer: Buried: %#v", err)
			return 0, err
		} else if cerr.Err == beanstalk.ErrNoCRLF {
			glog.Errorf("SetAnswer: Expected CRLF: %#v", err)
			return 0, err
		} else if cerr.Err == beanstalk.ErrJobTooBig {
			glog.Errorf("SetAnswer: Job too big: %#v", err)
			return 0, err
		} else if cerr.Err == beanstalk.ErrDraining {
			glog.Errorf("SetAnswer: Draining: %#v", err)
			return 0, err
		}
		glog.Errorf("SetAnswer: Unknown error: %#v", err)
		return 0, err
	}
	glog.Infof("SetAnswer: token: %q, id: %d", token, id)
	return id, nil
}

func GetAnswer(conn *beanstalk.Conn, token string) (id uint64, a *common.Answer, err error) {
	tube := beanstalk.Tube{
		Conn: conn,
		Name: token,
	}

	var body []byte
	id, body, err = tube.PeekReady()
	if err != nil {
		if cerr, ok := err.(beanstalk.ConnError); !ok {
			glog.Errorf("GetAnswer: Non-ConnError: %#v", err)
			return 0, nil, err
		} else if cerr.Err == beanstalk.ErrNotFound {
			glog.Infof("GetAnswer: Not found")
			return 0, nil, err
		}
		glog.Errorf("GetAnswer: Unknown error: %#v", err)
		return 0, nil, err
	}
	glog.Infof("GetAnswer: id: %d", id)

	err = json.NewDecoder(bytes.NewReader(body)).Decode(a)
	if err != nil {
		glog.Errorf("GetAnswer: Decode JSON: %#v", err)
		return 0, nil, err
	}

	return id, a, err
}
