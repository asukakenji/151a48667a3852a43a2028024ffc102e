package taskqueue

import (
	"bytes"
	"encoding/json"
	"time"

	"github.com/asukakenji/151a48667a3852a43a2028024ffc102e/common"
	"github.com/golang/glog"
	"github.com/kr/beanstalk"
)

func FetchQuestion(conn *beanstalk.Conn, timeout time.Duration) (id uint64, q *common.Question, err error) {
	for {
		var body []byte
		id, body, err = conn.Reserve(
			timeout, // timeout
		)
		if err != nil {
			if cerr, ok := err.(beanstalk.ConnError); !ok {
				glog.Errorf("FetchQuestion: Non-ConnError: %#v", err)
				return 0, nil, err
			} else if cerr.Err == beanstalk.ErrTimeout {
				glog.Infof("FetchQuestion: Timeout")
				continue
			} else if cerr.Err == beanstalk.ErrDeadline {
				glog.Infof("FetchQuestion: Deadline soon")
				time.Sleep(1 * time.Second)
				continue
			}
			glog.Errorf("FetchQuestion: Unknown error: %#v", err)
			return 0, nil, err
		}
		glog.Infof("FetchQuestion: id: %d", id)

		err = json.NewDecoder(bytes.NewReader(body)).Decode(q)
		if err != nil {
			glog.Errorf("FetchQuestion: Decode JSON: %#v", err)
			return 0, nil, err
		}
		return id, q, err
	}
}
