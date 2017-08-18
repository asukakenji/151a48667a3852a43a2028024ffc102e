package taskqueue

import "github.com/kr/beanstalk"

func RegisterGarbage(conn *beanstalk.Conn, token string, qid uint64) (id uint64, err error) {
	// TODO: Write this!
	return 0, nil
}
