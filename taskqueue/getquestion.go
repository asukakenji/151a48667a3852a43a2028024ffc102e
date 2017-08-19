package taskqueue

import "github.com/kr/beanstalk"

func GetQuestion(conn *beanstalk.Conn, qid uint64) (q *Question, err error) {
	return nil, nil
}
