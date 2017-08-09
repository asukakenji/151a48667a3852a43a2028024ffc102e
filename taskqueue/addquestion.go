package taskqueue

import (
	"bytes"
	"time"

	"github.com/asukakenji/151a48667a3852a43a2028024ffc102e/common"
	"github.com/golang/glog"
	"github.com/kr/beanstalk"
)

// timeLimit = execution time limit
func AddQuestion(conn *beanstalk.Conn, token string, locs common.Locations, timeLimit time.Duration) (id uint64, err error) {
	buf := new(bytes.Buffer)
	q := Question{
		Token:     token,
		Locations: locs,
	}
	err = q.ToJSON(buf)
	if err != nil {
		glog.Errorf("AddQuestion: Encode JSON: %#v", err)
		return 0, err
	}

	id, err = conn.Put(
		buf.Bytes(),               // body
		uint32(time.Now().Unix()), // pri
		time.Duration(0),          // delay: immediately ready
		timeLimit,                 // ttr: let caller set how long it is allowed to run
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
