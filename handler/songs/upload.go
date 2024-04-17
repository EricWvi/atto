package songs

import (
	"github.com/EricWvi/atto/service"
	"github.com/gin-gonic/gin"
)

func (b Base) Upload(c *gin.Context, body []byte) *UploadResponse {
	name := c.Request.URL.Query().Get("Name")

	link, err := service.UploadFile(name, body, c.ContentType())
	if err != nil {
		return nil
	}

	return &UploadResponse{
		Link: link,
	}
}

type UploadRequest struct {
	Name string `json:"name"`
}

type UploadResponse struct {
	Link string `json:"link"`
}
