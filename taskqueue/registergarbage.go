package taskqueue

import "time"

func RegisterGarbage(conn *Connection, token string, qid uint64) (id uint64, err error) {
	now := time.Now()
	_ = now
	// TODO: Write this!
	return 0, nil
}
