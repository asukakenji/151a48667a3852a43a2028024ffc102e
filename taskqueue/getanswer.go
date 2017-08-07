package taskqueue

import (
	"bytes"
	"encoding/json"

	"github.com/asukakenji/151a48667a3852a43a2028024ffc102e/common"
	"github.com/golang/glog"
	"github.com/kr/beanstalk"
)

func GetAnswer(conn *beanstalk.Conn, token string) (id uint64, a *common.Answer, err error) {
	tube := beanstalk.Tube{
		Conn: conn,
		Name: token,
	}

	var body []byte
	id, body, err = tube.PeekReady()
	if err != nil {
		if cerr, ok := err.(beanstalk.ConnError); !ok {
			glog.Errorf("GetAnswer: Non-ConnError: %#v", err)
			return 0, nil, err
		} else if cerr.Err == beanstalk.ErrNotFound {
			glog.Infof("GetAnswer: Not found")
			return 0, nil, err
		}
		glog.Errorf("GetAnswer: Unknown error: %#v", err)
		return 0, nil, err
	}
	glog.Infof("GetAnswer: id: %d", id)

	err = json.NewDecoder(bytes.NewReader(body)).Decode(a)
	if err != nil {
		glog.Errorf("GetAnswer: Decode JSON: %#v", err)
		return 0, nil, err
	}

	return id, a, nil
}
