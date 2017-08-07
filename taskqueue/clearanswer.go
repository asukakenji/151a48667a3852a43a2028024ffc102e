package taskqueue

import (
	"github.com/golang/glog"
	"github.com/kr/beanstalk"
)

func ClearAnswer(conn *beanstalk.Conn, token string) (err error) {
	tube := beanstalk.Tube{
		Conn: conn,
		Name: token,
	}
	for {
		var id uint64
		id, _, err = tube.PeekReady()
		if err != nil {
			if cerr, ok := err.(beanstalk.ConnError); !ok {
				glog.Errorf("ClearAnswer: Non-ConnError: %#v", err)
				return err
			} else if cerr.Err == beanstalk.ErrNotFound {
				return nil
			}
			glog.Errorf("ClearAnswer: Unknown error: %#v", err)
			return err
		}
		glog.Infof("ClearAnswer: Peek Ready: id: %d", id)

		err = conn.Delete(
			id, // id
		)
		if err != nil {
			if cerr, ok := err.(beanstalk.ConnError); !ok {
				glog.Errorf("ClearAnswer: Non-ConnError: %#v", err)
				return err
			} else if cerr.Err == beanstalk.ErrNotFound {
				glog.Errorf("ClearAnswer: Not found")
				return err
			}
			glog.Errorf("ClearAnswer: Unknown error: %#v", err)
			return err
		}
		glog.Infof("ClearAnswer: Deleted: id: %d", id)
	}
}
