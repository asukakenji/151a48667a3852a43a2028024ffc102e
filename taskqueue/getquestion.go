package taskqueue

import (
	"github.com/asukakenji/151a48667a3852a43a2028024ffc102e/common"
	"github.com/golang/glog"
	"github.com/kr/beanstalk"
)

func GetQuestion(conn *Connection, qid uint64) (q *Question, err common.Error) {
	body, _err := conn.Conn.Peek(
		qid, // id
	)
	if _err != nil {
		hash := common.NewToken()
		if cerr, ok := _err.(beanstalk.ConnError); !ok {
			glog.Errorf("[%s] GetQuestion: Non-ConnError", hash)
			return nil, NewUnexpectedError(_err, hash)
		} else if cerr.Err == beanstalk.ErrNotFound {
			glog.Infof("[%s] GetQuestion: Not found", hash)
			return nil, NewNotFoundError(_err, hash)
		}
		glog.Errorf("[%s] GetQuestion: Unknown error", hash)
		return nil, NewUnexpectedError(_err, hash)
	}
	glog.Info("GetQuestion: qid: %d", qid)

	q, err = QuestionFromJSONBytes(body)
	if err != nil {
		hash := common.NewToken()
		glog.Errorf("[%s] GetQuestion: Decode from JSON", hash)
		return nil, err
	}

	return q, nil
}
