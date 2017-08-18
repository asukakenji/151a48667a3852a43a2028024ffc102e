package taskqueue

import (
	"encoding/json"
	"io"

	"github.com/asukakenji/151a48667a3852a43a2028024ffc102e/common"
)

type Question struct {
	Token     string           `json:"t"`
	Locations common.Locations `json:"l"`
}

func QuestionFromJSON(r io.Reader) (q *Question, err error) {
	_err := json.NewDecoder(r).Decode(q)
	if _err != nil {
		return nil, common.NewJSONDecodeError(_err)
	}
	return q, nil
}

func (q *Question) ToJSON(w io.Writer) error {
	_err := json.NewEncoder(w).Encode(q)
	if _err != nil {
		return common.NewJSONEncodeError(_err)
	}
	return nil
}

type Answer struct {
	QuestionID   uint64               `json:"q"`
	TrialCount   int                  `json:"t"`
	DrivingRoute *common.DrivingRoute `json:"d"`
}
