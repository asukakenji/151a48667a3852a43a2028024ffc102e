package taskqueue

import (
	"time"

	"github.com/asukakenji/151a48667a3852a43a2028024ffc102e/common"
	"github.com/golang/glog"
	"github.com/kr/beanstalk"
)

func FetchQuestion(conn *Connection) (qid uint64, q *Question, err common.Error) {
	for {
		qid, body, _err := conn.Conn.Reserve(
			TimeForever, // timeout
		)
		if _err != nil {
			hash := common.NewToken()
			if cerr, ok := _err.(beanstalk.ConnError); !ok {
				glog.Errorf("[%s] FetchQuestion: Non-ConnError", hash)
				return 0, nil, NewUnexpectedError(_err, hash)
			} else if cerr.Err == beanstalk.ErrTimeout {
				glog.Infof("[%s] FetchQuestion: Timeout", hash)
				continue
			} else if cerr.Err == beanstalk.ErrDeadline {
				glog.Infof("[%s] FetchQuestion: Deadline Soon", hash)
				time.Sleep(1 * time.Second)
				continue
			}
			glog.Errorf("[%s] FetchQuestion: Unknown ConnError", hash)
			return 0, nil, NewUnexpectedError(_err, hash)
		}
		glog.Infof("FetchQuestion: qid: %d", qid)

		q, err := QuestionFromJSONBytes(body)
		if err != nil {
			glog.Errorf("[%s] FetchQuestion: Decode from JSON", err.Hash())
			return 0, nil, err
		}

		return qid, q, nil
	}
}
