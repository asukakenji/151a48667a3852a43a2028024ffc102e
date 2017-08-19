package taskqueue

import (
	"time"

	"github.com/kr/beanstalk"
)

func RegisterGarbage(conn *beanstalk.Conn, token string, qid uint64) (id uint64, err error) {
	now := time.Now()
	_ = now
	// TODO: Write this!
	return 0, nil
}
