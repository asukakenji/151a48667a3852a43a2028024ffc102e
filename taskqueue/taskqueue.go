// Tier 1:
// - Handle Route:
//   - USE "default"
//   - PUT
//     - body: token + locations
//     - pri: timestamp
//     - delay: 0 seconds
//     - ttr: 3 seconds (customizable)
// - Handle Route Token:
//   - USE "<token>"
//   - PEEK-READY
//   - Return result

// Tier 2:
// - Handle Task:
//   - USE "default"
//   - RESERVE
//   - USE "garbage"
//   - PUT
//     - body: token
//     - pri: (current) timestamp
//     - delay: 600 seconds (customizable)
//     - ttr: 3 seconds (customizable)
//   - USE "<token>"
//   - PEEK-READY
//   - Depending on the result:
//     - Not Found:
//       - PUT
//         - body: "in progress" + trialCount (= 1)
//         - pri: max uint32 - (current) timestamp
//         - delay: 0 seconds
//         - ttr: 0 seconds
//     - trialCount != maxTrialCount
//       - PUT
//         - body: "in progress" + trialCount (+= 1)
//         - pri: max uint32 - (current) timestamp
//         - delay: 0 seconds
//         - ttr: 0 seconds
//     - trialCount == maxTrialCount
//       - PUT
//         - body: "failure" + trialCount (= maxTrialCount)
//         - pri: max uint32 - (current) timestamp
//         - delay: 0 seconds
//         - ttr: 0 seconds
//       - Return
//     -  Not Matched:
//        - TODO: Write this!
//   - Google Maps
//     - Get Distance Matrix
//   - Travelling Salesman
//   - PUT
//     - body: "success" + path + other results
//     - pri: max uint32 - (current) timestamp
//     - delay: 0 seconds
//     - ttr: 0 seconds
//   - DELETE
//     - id: id of the job

// Tier C:
// - Loop
//   - USE "garbage"
//   - RESERVE
//   - USE "<token>"
//     - Loop
//       - PEEK-READY
//       - DELETE
//         - id: id of the job just reserved

package taskqueue

import (
	"math"
	"time"

	"github.com/asukakenji/151a48667a3852a43a2028024ffc102e/common"
	"github.com/golang/glog"
	"github.com/kr/beanstalk"
)

const (
	// TimeForever represents a very large duration.
	TimeForever = math.MaxUint32 * time.Second
)

func WithConnection(addr string, do func(*Connection) error) error {
	conn, err := beanstalk.Dial("tcp", addr)
	if err != nil {
		hash := common.NewToken()
		glog.Errorf("WithConnection: cannot dial (%s)", hash)
		return NewConnectionError(err, hash)
	}
	defer conn.Close()
	return do(&Connection{conn})
}
