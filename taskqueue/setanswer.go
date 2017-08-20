package taskqueue

import (
	"math"
	"time"

	"github.com/asukakenji/151a48667a3852a43a2028024ffc102e/common"
	"github.com/golang/glog"
	"github.com/kr/beanstalk"
)

func SetAnswer(conn *Connection, token string, qid uint64, rc int, dr *common.DrivingRoute) (aid uint64, err common.Error) {
	now := time.Now()
	a, err := Answer{
		Timestamp:    now,
		QuestionID:   qid,
		RetryCount:   rc,
		DrivingRoute: dr,
	}.ToJSONBytes()
	if err != nil {
		glog.Errorf("[%s] SetAnswer: Encode to JSON", err.Hash())
		return 0, err
	}

	tube := beanstalk.Tube{
		Conn: conn.Conn,
		Name: token,
	}
	pri := math.MaxUint32 - uint32(now.Unix())
	aid, _err := tube.Put(
		a,                // body
		pri,              // pri
		time.Duration(0), // delay: immediately ready
		time.Duration(0), // ttr: zero as answers are never reserved
	)
	if _err != nil {
		hash := common.NewToken()
		if cerr, ok := _err.(beanstalk.ConnError); !ok {
			glog.Errorf("[%s] SetAnswer: Non-ConnError", hash)
			return 0, NewUnexpectedError(_err, hash)
		} else if cerr.Err == beanstalk.ErrBuried {
			glog.Errorf("[%s] SetAnswer: Burried", hash)
			return 0, NewUnexpectedError(_err, hash)
		} else if cerr.Err == beanstalk.ErrNoCRLF {
			glog.Errorf("[%s] SetAnswer: Expected CRLF", hash)
			return 0, NewUnexpectedError(_err, hash)
		} else if cerr.Err == beanstalk.ErrJobTooBig {
			glog.Errorf("[%s] SetAnswer: Job Too Big", hash)
			return 0, NewJobTooBigError(_err, hash)
		} else if cerr.Err == beanstalk.ErrDraining {
			glog.Errorf("[%s] SetAnswer: Draining", hash)
			return 0, NewUnexpectedError(_err, hash)
		}
		glog.Errorf("[%s] SetAnswer: Unknown ConnError", hash)
		return 0, NewUnexpectedError(_err, hash)
	}
	glog.Infof("SetAnswer: token: %q, qid: %d, rc: %d, aid: %d", token, qid, rc, aid)
	return aid, nil
}
