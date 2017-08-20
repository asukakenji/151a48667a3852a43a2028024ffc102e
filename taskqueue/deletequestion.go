package taskqueue

import (
	"github.com/asukakenji/151a48667a3852a43a2028024ffc102e/common"
	"github.com/golang/glog"
	"github.com/kr/beanstalk"
)

func DeleteJob(conn *Connection, jid uint64) common.Error {
	_err := conn.Conn.Delete(
		jid, // id
	)
	if _err != nil {
		hash := common.NewToken()
		if cerr, ok := _err.(beanstalk.ConnError); !ok {
			glog.Errorf("[%s] DeleteJob: Non-ConnError", hash)
			return NewUnexpectedError(_err, hash)
		} else if cerr.Err == beanstalk.ErrNotFound {
			glog.Errorf("[%s] DeleteJob: Not found", hash)
			return NewNotFoundError(_err, hash)
		}
		glog.Errorf("[%s] DeleteJob: Unknown ConnError", hash)
		return NewUnexpectedError(_err, hash)
	}
	glog.Infof("DeleteJob: jid: %d", jid)
	return nil
}
