package taskqueue

import (
	"encoding/json"
	"io"
	"time"

	"github.com/asukakenji/151a48667a3852a43a2028024ffc102e/common"
	"github.com/kr/beanstalk"
)

type Connection struct {
	Conn *beanstalk.Conn
}

type Question struct {
	Timestamp time.Time        `json:"t"`
	Token     string           `json:"x"`
	Locations common.Locations `json:"l"`
}

func QuestionFromJSON(r io.Reader) (q *Question, err error) {
	_err := json.NewDecoder(r).Decode(q)
	if _err != nil {
		hash := common.NewToken()
		return nil, common.NewJSONDecodeError(_err, hash)
	}
	return q, nil
}

func (q *Question) ToJSON(w io.Writer) error {
	_err := json.NewEncoder(w).Encode(q)
	if _err != nil {
		hash := common.NewToken()
		return common.NewJSONEncodeError(_err, hash)
	}
	return nil
}

type Answer struct {
	Timestamp    time.Time            `json:"t"`
	QuestionID   uint64               `json:"q"`
	RetryCount   int                  `json:"r"`
	DrivingRoute *common.DrivingRoute `json:"d"`
}

type Garbage struct {
	Timestamp  time.Time `json:"t"`
	Token      string    `json:"x"`
	QuestionID uint64    `json:"q"`
}
