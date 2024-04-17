package ping

import (
	"github.com/EricWvi/atto/handler"
	"github.com/gin-gonic/gin"
)

type Base struct{}

func DefaultHandler(c *gin.Context) {
	handler.Dispatch(c, Base{})
}

func (b Base) Ping(c *gin.Context, req *PingRequest) *PingResponse {
	return &PingResponse{
		Value: "Pong",
	}
}

type PingRequest struct {
}

type PingResponse struct {
	Value string `json:"value"`
}
