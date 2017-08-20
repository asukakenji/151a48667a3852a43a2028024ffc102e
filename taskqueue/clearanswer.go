package taskqueue

import (
	"github.com/asukakenji/151a48667a3852a43a2028024ffc102e/common"
	"github.com/golang/glog"
	"github.com/kr/beanstalk"
)

// TODO: Error handling not yet updated
func ClearAnswer(conn *Connection, token string) (id uint64, a *Answer, err common.Error) {
	tube := beanstalk.Tube{
		Conn: conn.Conn,
		Name: token,
	}

	var lastBody []byte
	for {
		var body []byte
		id, body, _err := tube.PeekReady()
		if err != nil {
			if cerr, ok := _err.(beanstalk.ConnError); !ok {
				glog.Errorf("ClearAnswer: Non-ConnError: %#v", err)
				return 0, nil, err
			} else if cerr.Err == beanstalk.ErrNotFound {
				break
			}
			glog.Errorf("ClearAnswer: Unknown error: %#v", err)
			return 0, nil, err
		}
		glog.Infof("ClearAnswer: Peek Ready: id: %d", id)

		_err = conn.Conn.Delete(
			id, // id
		)
		if _err != nil {
			if cerr, ok := _err.(beanstalk.ConnError); !ok {
				glog.Errorf("ClearAnswer: Non-ConnError: %#v", err)
				return 0, nil, err
			} else if cerr.Err == beanstalk.ErrNotFound {
				glog.Errorf("ClearAnswer: Not found")
				return 0, nil, err
			}
			glog.Errorf("ClearAnswer: Unknown error: %#v", err)
			return 0, nil, err
		}
		glog.Infof("ClearAnswer: Deleted: id: %d", id)
		lastBody = body
	}

	a, err = AnswerFromJSONBytes(lastBody)
	if err != nil {
		glog.Errorf("ClearAnswer: Decode JSON: %#v", err)
		return 0, nil, err
	}

	return id, a, nil
}
