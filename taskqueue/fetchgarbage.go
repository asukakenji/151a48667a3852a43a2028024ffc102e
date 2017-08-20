package taskqueue

import (
	"time"

	"github.com/asukakenji/151a48667a3852a43a2028024ffc102e/common"
	"github.com/golang/glog"
	"github.com/kr/beanstalk"
)

func FetchGarbage(conn *Connection) (gid uint64, g *Garbage, err common.Error) {
	tubeSet := beanstalk.NewTubeSet(
		conn.Conn,
		"garbage",
	)

	for {
		gid, body, _err := tubeSet.Reserve(
			TimeForever, // timeout
		)
		if _err != nil {
			hash := common.NewToken()
			if cerr, ok := _err.(beanstalk.ConnError); !ok {
				glog.Errorf("[%s] FetchGarbage: Non-ConnError", hash)
				return 0, nil, NewUnexpectedError(_err, hash)
			} else if cerr.Err == beanstalk.ErrTimeout {
				glog.Infof("[%s] FetchGarbage: Timeout", hash)
				continue
			} else if cerr.Err == beanstalk.ErrDeadline {
				glog.Infof("[%s] FetchGarbage: Deadline Soon", hash)
				time.Sleep(1 * time.Second)
				continue
			}
			glog.Errorf("[%s] FetchGarbage: Unknown ConnError", hash)
			return 0, nil, NewUnexpectedError(_err, hash)
		}
		glog.Infof("FetchGarbage: gid: %d", gid)

		g, err := GarbageFromJSONBytes(body)
		if err != nil {
			glog.Errorf("[%s] FetchGarbage: Decode from JSON", err.Hash())
			return 0, nil, err
		}

		return gid, g, nil
	}
}
