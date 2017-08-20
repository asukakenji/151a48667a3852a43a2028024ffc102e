package taskqueue

import (
	"bytes"
	"encoding/json"

	"github.com/asukakenji/151a48667a3852a43a2028024ffc102e/common"
	"github.com/golang/glog"
	"github.com/kr/beanstalk"
)

// TODO: Error handling not yet updated
func GetAnswer2(conn *Connection, token string) (id uint64, a *Answer, err common.Error) {
	tube := beanstalk.Tube{
		Conn: conn.Conn,
		Name: token,
	}

	var body []byte
	id, body, _err := tube.PeekReady()
	if _err != nil {
		hash := common.NewToken()
		if cerr, ok := _err.(beanstalk.ConnError); !ok {
			glog.Errorf("[%s] GetAnswer: Non-ConnError", hash)
			return 0, nil, err
		} else if cerr.Err == beanstalk.ErrNotFound {
			glog.Infof("[%s] GetAnswer: Not found", hash)
			return 0, nil, err
		}
		glog.Errorf("[%s] GetAnswer: Unknown error", hash)
		return 0, nil, err
	}
	glog.Infof("GetAnswer: id: %d", id)

	_err = json.NewDecoder(bytes.NewReader(body)).Decode(a)
	if err != nil {
		glog.Errorf("GetAnswer: Decode JSON: %#v", err)
		return 0, nil, err
	}

	return id, a, nil
}

func GetAnswer1(conn *Connection, token string) (id uint64, dr *common.DrivingRoute, err common.Error) {
	id, a, err := GetAnswer2(conn, token)
	if err != nil {
		return 0, nil, err
	}
	return id, a.DrivingRoute, nil
}
