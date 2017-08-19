package taskqueue

import (
	"bytes"
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

func QuestionFromJSON(r io.Reader) (q *Question, err common.MyError) {
	_err := json.NewDecoder(r).Decode(q)
	if _err != nil {
		hash := common.NewToken()
		return nil, common.NewJSONDecodeError(_err, hash)
	}
	return q, nil
}

func QuestionFromJSONBytes(bs []byte) (q *Question, err common.MyError) {
	return QuestionFromJSON(bytes.NewReader(bs))
}

func (q Question) ToJSON(w io.Writer) common.MyError {
	_err := json.NewEncoder(w).Encode(q)
	if _err != nil {
		hash := common.NewToken()
		return common.NewJSONEncodeError(_err, hash)
	}
	return nil
}

func (q Question) ToJSONBytes() ([]byte, common.MyError) {
	w := new(bytes.Buffer)
	err := q.ToJSON(w)
	if err != nil {
		return nil, err
	}
	return w.Bytes(), nil
}

type Answer struct {
	Timestamp    time.Time            `json:"t"`
	QuestionID   uint64               `json:"q"`
	RetryCount   int                  `json:"r"`
	DrivingRoute *common.DrivingRoute `json:"d"`
}

func AnswerFromJSON(r io.Reader) (a *Answer, err common.MyError) {
	_err := json.NewDecoder(r).Decode(a)
	if _err != nil {
		hash := common.NewToken()
		return nil, common.NewJSONDecodeError(_err, hash)
	}
	return a, nil
}

func AnswerFromJSONBytes(bs []byte) (q *Answer, err common.MyError) {
	return AnswerFromJSON(bytes.NewReader(bs))
}

func (a Answer) ToJSON(w io.Writer) common.MyError {
	_err := json.NewEncoder(w).Encode(a)
	if _err != nil {
		hash := common.NewToken()
		return common.NewJSONEncodeError(_err, hash)
	}
	return nil
}

func (a Answer) ToJSONBytes() ([]byte, common.MyError) {
	w := new(bytes.Buffer)
	err := a.ToJSON(w)
	if err != nil {
		return nil, err
	}
	return w.Bytes(), nil
}

type Garbage struct {
	Timestamp  time.Time `json:"t"`
	Token      string    `json:"x"`
	QuestionID uint64    `json:"q"`
}

func GarbageFromJSON(r io.Reader) (g *Garbage, err common.MyError) {
	_err := json.NewDecoder(r).Decode(g)
	if _err != nil {
		hash := common.NewToken()
		return nil, common.NewJSONDecodeError(_err, hash)
	}
	return g, nil
}

func GarbageFromJSONBytes(bs []byte) (g *Garbage, err common.MyError) {
	return GarbageFromJSON(bytes.NewReader(bs))
}

func (g Garbage) ToJSON(w io.Writer) common.MyError {
	_err := json.NewEncoder(w).Encode(g)
	if _err != nil {
		hash := common.NewToken()
		return common.NewJSONEncodeError(_err, hash)
	}
	return nil
}

func (g Garbage) ToJSONBytes() ([]byte, common.MyError) {
	buf := new(bytes.Buffer)
	err := g.ToJSON(buf)
	if err != nil {
		return nil, err
	}
	return buf.Bytes(), nil
}
