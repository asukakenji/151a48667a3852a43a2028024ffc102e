package taskqueue

import "github.com/kr/beanstalk"

func RegisterGarbage(conn *beanstalk.Conn, qid uint64, token string) (id uint64, err error) {
	// TODO: Write this!
	return 0, nil
}
