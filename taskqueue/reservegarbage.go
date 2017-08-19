package taskqueue

import "github.com/kr/beanstalk"

func ReserveGarbage(conn *beanstalk.Conn) (id uint64, g *Garbage, err error) {
	// TODO: Write this!
	return 0, nil, nil
}
