package taskqueue

import (
	"time"

	"github.com/asukakenji/151a48667a3852a43a2028024ffc102e/common"
	"github.com/golang/glog"
	"github.com/kr/beanstalk"
)

func RegisterGarbage(conn *Connection, token string, qid uint64, delay time.Duration, ttr time.Duration) (id uint64, err error) {
	now := time.Now()
	g, err := Garbage{
		Timestamp:  now,
		Token:      token,
		QuestionID: qid,
	}.ToJSONBytes()
	if err != nil {
		err := err.(common.MyError)
		glog.Errorf("[%s] RegisterGarbage: Encode to JSON", err.Hash())
		return 0, err
	}

	tube := beanstalk.Tube{
		Conn: conn.Conn,
		Name: "garbage",
	}
	id, err = tube.Put(
		g,                  // body
		uint32(now.Unix()), // pri
		delay,              // delay: let caller set how long to wait until it appears in the ready queue
		ttr,                // ttr: let caller set how long it is allowed to run
	)
	if err != nil {
		hash := common.NewToken()
		if cerr, ok := err.(beanstalk.ConnError); !ok {
			glog.Errorf("[%s] RegisterGarbage: Non-ConnError", hash)
			return 0, NewUnexpectedError(err, hash)
		} else if cerr.Err == beanstalk.ErrBuried {
			glog.Errorf("[%s] RegisterGarbage: Burried", hash)
			return 0, NewUnexpectedError(err, hash)
		} else if cerr.Err == beanstalk.ErrNoCRLF {
			glog.Errorf("[%s] RegisterGarbage: Expected CRLF", hash)
			return 0, NewUnexpectedError(err, hash)
		} else if cerr.Err == beanstalk.ErrJobTooBig {
			glog.Errorf("[%s] RegisterGarbage: Job Too Big", hash)
			return 0, NewJobTooBigError(err, hash)
		} else if cerr.Err == beanstalk.ErrDraining {
			glog.Errorf("[%s] RegisterGarbage: Draining", hash)
			return 0, NewUnexpectedError(err, hash)
		}
		glog.Errorf("[%s] RegisterGarbage: Unknown ConnError", hash)
		return 0, NewUnexpectedError(err, hash)
	}
	glog.Infof("RegisterGarbage: token: %q, qid: %d, gid: %d", token, id)
	return id, nil
}
