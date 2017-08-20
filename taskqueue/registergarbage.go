package taskqueue

import (
	"time"

	"github.com/asukakenji/151a48667a3852a43a2028024ffc102e/common"
	"github.com/golang/glog"
	"github.com/kr/beanstalk"
)

func RegisterGarbage(conn *Connection, token string, qid uint64, delay time.Duration, ttr time.Duration) (gid uint64, err common.Error) {
	now := time.Now()
	g, err := Garbage{
		Timestamp:  now,
		Token:      token,
		QuestionID: qid,
	}.ToJSONBytes()
	if err != nil {
		glog.Errorf("[%s] RegisterGarbage: Encode to JSON", err.Hash())
		return 0, err
	}

	tube := beanstalk.Tube{
		Conn: conn.Conn,
		Name: "garbage",
	}
	gid, _err := tube.Put(
		g,                  // body
		uint32(now.Unix()), // pri
		delay,              // delay: let caller set how long to wait until it appears in the ready queue
		ttr,                // ttr: let caller set how long it is allowed to run
	)
	if _err != nil {
		hash := common.NewToken()
		if cerr, ok := _err.(beanstalk.ConnError); !ok {
			glog.Errorf("[%s] RegisterGarbage: Non-ConnError", hash)
			return 0, NewUnexpectedError(_err, hash)
		} else if cerr.Err == beanstalk.ErrBuried {
			glog.Errorf("[%s] RegisterGarbage: Burried", hash)
			return 0, NewUnexpectedError(_err, hash)
		} else if cerr.Err == beanstalk.ErrNoCRLF {
			glog.Errorf("[%s] RegisterGarbage: Expected CRLF", hash)
			return 0, NewUnexpectedError(_err, hash)
		} else if cerr.Err == beanstalk.ErrJobTooBig {
			glog.Errorf("[%s] RegisterGarbage: Job Too Big", hash)
			return 0, NewJobTooBigError(_err, hash)
		} else if cerr.Err == beanstalk.ErrDraining {
			glog.Errorf("[%s] RegisterGarbage: Draining", hash)
			return 0, NewUnexpectedError(_err, hash)
		}
		glog.Errorf("[%s] RegisterGarbage: Unknown ConnError", hash)
		return 0, NewUnexpectedError(_err, hash)
	}
	glog.Infof("RegisterGarbage: token: %q, qid: %d, gid: %d", token, gid)
	return gid, nil
}
