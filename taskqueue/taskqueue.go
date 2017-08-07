// Tier 1:
// - Handle Route:
//   - USE "default"
//   - PUT
//     - body: token + locations
//     - pri: timestamp
//     - delay: 0 seconds
//     - ttr: 3 seconds
// - Handle Route Token:
//   - USE "<token>"
//   - RESERVE
//   - Return result

// Tier 2:
// - Handle Task:
//   - USE "default"
//   - RESERVE
//   - USE "<token>"
//   - Loop
//     - RESERVE
//     - DELETE
//       - id: id of job just reserved
//   - Depending on the result:
//     - Not Found:
//       - PUT
//         - body: "in progress" + trialCount (= 1)
//         - pri: (current) timestamp
//         - delay: 0 seconds
//         - ttr: 0 seconds
//     - trialCount != maxTrialCount
//       - PUT
//         - body: "in progress" + trialCount (+= 1)
//         - pri: (current) timestamp
//         - delay: 0 seconds
//         - ttr: 0 seconds
//     - trialCount == maxTrialCount
//       - PUT
//         - body: "failure" + trialCount (= maxTrialCount)
//         - pri: (current) timestamp
//         - delay: 0 seconds
//         - ttr: 0 seconds
//       - Return
//   - Google Maps
//     - Use goroutine to parallel the queries
//     - Or see whether the API allows many questions in one trip
//   - Travelling Salesman
//   - PUT
//     - body: "success" + path + other results
//     - pri: (current) timestamp
//     - delay: 0 seconds
//     - ttr: 0 seconds
//   - USE "garbage"
//   - PUT
//     - body: token
//     - pri: (current) timestamp
//     - delay: 600 seconds
//     - ttr: 0 seconds
//   - USE "default"
//   - DELETE
//     - id: id of the job

// Tier C:
// - Loop
//   - USE "garbage"
//   - RESERVE
//   - USE "<token>"
//     - Loop
//       - RESERVE
//       - DELETE
//         - id: id of the job just reserved

package taskqueue

import (
	"math"
	"time"

	"github.com/golang/glog"
	"github.com/kr/beanstalk"
)

const (
	// TimeForever represents a very large duration.
	TimeForever = math.MaxUint32 * time.Second
)

func WithConnection(addr string, do func(*beanstalk.Conn) error) error {
	conn, err := beanstalk.Dial("tcp", addr)
	if err != nil {
		glog.Errorf("WithConnection: Dial error: %#v", err)
		return err
	}
	defer conn.Close()
	return do(conn)
}
