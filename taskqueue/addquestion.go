package taskqueue

import (
	"time"

	"github.com/asukakenji/151a48667a3852a43a2028024ffc102e/common"
	"github.com/golang/glog"
	"github.com/kr/beanstalk"
)

// timeLimit = execution time limit
// ttr = time to run
func AddQuestion(conn *Connection, token string, locs common.Locations, ttr time.Duration) (qid uint64, err common.Error) {
	now := time.Now()
	q, err := Question{
		Timestamp: now,
		Token:     token,
		Locations: locs,
	}.ToJSONBytes()
	if err != nil {
		glog.Errorf("[%s] AddQuestion: Encode to JSON", err.Hash())
		return 0, err
	}

	qid, _err := conn.Conn.Put(
		q,                  // body
		uint32(now.Unix()), // pri
		time.Duration(0),   // delay: immediately ready
		ttr,                // ttr: let caller set how long it is allowed to run
	)
	if _err != nil {
		hash := common.NewToken()
		if cerr, ok := _err.(beanstalk.ConnError); !ok {
			glog.Errorf("[%s] AddQuestion: Non-ConnError", hash)
			return 0, NewUnexpectedError(_err, hash)
		} else if cerr.Err == beanstalk.ErrBuried {
			glog.Errorf("[%s] AddQuestion: Burried", hash)
			return 0, NewUnexpectedError(_err, hash)
		} else if cerr.Err == beanstalk.ErrNoCRLF {
			glog.Errorf("[%s] AddQuestion: Expected CRLF", hash)
			return 0, NewUnexpectedError(_err, hash)
		} else if cerr.Err == beanstalk.ErrJobTooBig {
			glog.Errorf("[%s] AddQuestion: Job Too Big", hash)
			return 0, NewJobTooBigError(_err, hash)
		} else if cerr.Err == beanstalk.ErrDraining {
			glog.Errorf("[%s] AddQuestion: Draining", hash)
			return 0, NewUnexpectedError(_err, hash)
		}
		glog.Errorf("[%s] AddQuestion: Unknown ConnError", hash)
		return 0, NewUnexpectedError(_err, hash)
	}
	glog.Infof("AddQuestion: token: %q, qid: %d", token, qid)
	return qid, nil
}
