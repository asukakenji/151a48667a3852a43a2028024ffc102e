package taskqueue

import (
	"github.com/asukakenji/151a48667a3852a43a2028024ffc102e/common"
)

type Question struct {
	Token     string           `json:"token"`
	Locations common.Locations `json:"locations"`
}

type Answer struct {
	Timestamp    int64                `json:"timestamp"`
	DrivingRoute *common.DrivingRoute `json:"driving_route"`
}
