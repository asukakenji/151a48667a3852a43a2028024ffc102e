package taskqueue

import (
	"bytes"
	"encoding/json"
	"math"
	"time"

	"github.com/asukakenji/151a48667a3852a43a2028024ffc102e/common"
	"github.com/golang/glog"
	"github.com/kr/beanstalk"
)

func SetAnswer(conn *beanstalk.Conn, token string, qid uint64, tc int, dr *common.DrivingRoute) (id uint64, err error) {
	buf := new(bytes.Buffer)
	err = json.NewEncoder(buf).Encode(Answer{
		QuestionID:   qid,
		TrialCount:   tc,
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
		buf.Bytes(),                              // body
		math.MaxUint32-uint32(time.Now().Unix()), // pri
		time.Duration(0),                         // delay: immediately ready
		time.Duration(0),                         // ttr: zero as answers are never reserved
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
