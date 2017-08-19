package main

import (
	"flag"

	"github.com/asukakenji/151a48667a3852a43a2028024ffc102e/taskqueue"
	"github.com/golang/glog"
	"github.com/kr/beanstalk"
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
		err := taskqueue.WithConnection(addr, func(conn *beanstalk.Conn) error {
			for {
				gid, g, _err := taskqueue.ReserveGarbage(conn)
				if _err != nil {
					glog.Errorf("main: cannot reserve garbage")
					return _err
				}
				glog.Infof("gid: %d", gid)

				token := g.Token
				aid, a, _err := taskqueue.GetAnswer2(conn, token)
				if _err != nil {
					glog.Errorf("main: cannot get answer")
					return _err
				}
				glog.Infof("aid: %d", aid)

				// TODO: Compare the timestamps:
				// - g.Timestamp
				// - a.Timestamp
				// - time.Now()
				glog.Infof("main: a.Timestamp: %d", a.Timestamp)

				// TODO: Move the logic here
				taskqueue.ClearAnswer(conn, token)

				q, _err := taskqueue.GetQuestion(conn, g.QuestionID)
				if _err != nil {
					glog.Errorf("main: cannot get question")
					return _err
				}
				glog.Infof("main: q.Timestamp: %d", q.Timestamp)

				// TODO: Compare the timestamps:
				// - g.Timestamp
				// - q.Timestamp
				// - time.Now()
				_err = taskqueue.DeleteQuestion(conn, g.QuestionID)
				if _err != nil {
					glog.Errorf("main: cannot delete question")
					return _err
				}
			}
		})
		if err != nil {
			glog.Errorf("%v", err)
		}
	}
}
