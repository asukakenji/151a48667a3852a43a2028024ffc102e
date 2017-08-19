package taskqueue

import "github.com/kr/beanstalk"

func DeleteQuestion(conn *beanstalk.Conn, qid uint64) error {
	return nil
}
