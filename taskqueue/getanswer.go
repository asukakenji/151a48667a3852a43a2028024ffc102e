package taskqueue

import (
	"github.com/asukakenji/151a48667a3852a43a2028024ffc102e/common"
	"github.com/golang/glog"
	"github.com/kr/beanstalk"
)

func GetAnswer1(conn *Connection, token string) (aid uint64, dr *common.DrivingRoute, err common.Error) {
	aid, a, err := GetAnswer2(conn, token)
	if err != nil {
		return 0, nil, err
	}
	return aid, a.DrivingRoute, nil
}

func GetAnswer2(conn *Connection, token string) (aid uint64, a *Answer, err common.Error) {
	tube := beanstalk.Tube{
		Conn: conn.Conn,
		Name: token,
	}

	aid, body, _err := tube.PeekReady()
	if _err != nil {
		hash := common.NewToken()
		if cerr, ok := _err.(beanstalk.ConnError); !ok {
			glog.Errorf("[%s] GetAnswer: Non-ConnError", hash)
			return 0, nil, NewUnexpectedError(_err, hash)
		} else if cerr.Err == beanstalk.ErrNotFound {
			glog.Infof("[%s] GetAnswer: Not found", hash)
			return 0, nil, NewNotFoundError(_err, hash)
		}
		glog.Errorf("[%s] GetAnswer: Unknown error", hash)
		return 0, nil, NewUnexpectedError(_err, hash)
	}
	glog.Infof("GetAnswer: aid: %d", aid)

	a, err = AnswerFromJSONBytes(body)
	if err != nil {
		hash := common.NewToken()
		glog.Errorf("[%s] GetAnswer: Decode from JSON", hash)
		return 0, nil, err
	}

	return aid, a, nil
}
