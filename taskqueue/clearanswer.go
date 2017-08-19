package taskqueue

import (
	"bytes"
	"encoding/json"

	"github.com/golang/glog"
	"github.com/kr/beanstalk"
)

func ClearAnswer(conn *Connection, token string) (id uint64, a *Answer, err error) {
	tube := beanstalk.Tube{
		Conn: conn.Conn,
		Name: token,
	}

	var lastBody []byte
	for {
		var body []byte
		id, body, err = tube.PeekReady()
		if err != nil {
			if cerr, ok := err.(beanstalk.ConnError); !ok {
				glog.Errorf("ClearAnswer: Non-ConnError: %#v", err)
				return 0, nil, err
			} else if cerr.Err == beanstalk.ErrNotFound {
				break
			}
			glog.Errorf("ClearAnswer: Unknown error: %#v", err)
			return 0, nil, err
		}
		glog.Infof("ClearAnswer: Peek Ready: id: %d", id)

		err = conn.Conn.Delete(
			id, // id
		)
		if err != nil {
			if cerr, ok := err.(beanstalk.ConnError); !ok {
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

	err = json.NewDecoder(bytes.NewReader(lastBody)).Decode(a)
	if err != nil {
		glog.Errorf("ClearAnswer: Decode JSON: %#v", err)
		return 0, nil, err
	}

	return id, a, nil
}
