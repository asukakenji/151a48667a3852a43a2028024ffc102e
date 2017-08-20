package main

import (
	"flag"

	"github.com/asukakenji/151a48667a3852a43a2028024ffc102e/common"
	"github.com/asukakenji/151a48667a3852a43a2028024ffc102e/taskqueue"
	"github.com/golang/glog"
)

func main() {
	// Check whether "-logtostderr=true" or "-logtostderr=false" is provided in
	// command line. If yes, obey the command line option. Otherwise, use the
	// default, "true".
	isLogToStderrProvided := false
	flag.Parse()
	flag.Visit(func(f *flag.Flag) {
		if f.Name == "logtostderr" {
			isLogToStderrProvided = true
		}
	})
	if !isLogToStderrProvided {
		flag.Set("logtostderr", "true")
	}

	addr := "127.0.0.1:11300" // TODO: Customize: addr

	for {
		err := taskqueue.WithConnection(addr, func(conn *taskqueue.Connection) common.Error {
			for {
				gid, g, err2 := taskqueue.FetchGarbage(conn)
				if err2 != nil {
					glog.Errorf("[%s] main: cannot reserve garbage", err2.Hash())
					return err2
				}
				glog.Infof("gid: %d", gid)

				token := g.Token
				aid, a, err2 := taskqueue.GetAnswer2(conn, token)
				if err2 != nil {
					glog.Errorf("[%s] main: cannot get answer", err2.Hash())
					return err2
				}
				glog.Infof("aid: %d", aid)

				// TODO: Compare the timestamps:
				// - g.Timestamp
				// - a.Timestamp
				// - time.Now()
				glog.Infof("main: a.Timestamp: %d", a.Timestamp)

				err2 = taskqueue.DeleteJob(conn, aid)
				if err2 != nil {
					// TODO: Handle error
					return err2
				}

				q, err2 := taskqueue.GetQuestion(conn, g.QuestionID)
				if err2 != nil {
					glog.Errorf("[%s] main: cannot get question", err2.Hash())
					return err2
				}
				glog.Infof("main: q.Timestamp: %d", q.Timestamp)

				// TODO: Compare the timestamps:
				// - g.Timestamp
				// - q.Timestamp
				// - time.Now()
				err2 = taskqueue.DeleteJob(conn, g.QuestionID)
				if err2 != nil {
					glog.Errorf("[%s] main: cannot delete question", err2.Hash())
					return err2
				}
			}
		})
		if err != nil {
			glog.Errorf("[%s] main: %v", err.Hash(), err)
		}
	}
}
