package save

import (
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"net/http"
	"os"
	"path/filepath"

	r "gopi/internal/lib/response"
)

const (
	BadRequest  = http.StatusBadRequest
	ServerError = http.StatusInternalServerError
)

type Request struct {
	Path string `json:"path" validate:"required"`
	UUID string `json:"uuid,omitempty"`
}

type Response struct {
	r.Response
	UUID string `json:"uuid,omitempty"`
	Path string `json:"path,omitempty"`
}

type GifSaver interface {
	SaveGif(uuid string, path string) (int64, error)
}

func New(gifSaver GifSaver) gin.HandlerFunc {
	return func(c *gin.Context) {
		contentType := c.ContentType()
		if contentType != "multipart/form-data" && contentType != "application/octet-stream" {
			c.JSON(BadRequest, r.InvalidContentType)
			return
		}

		gifUUID := uuid.New().String()
		serverFilePath := filepath.Join("gifs", gifUUID+".gif")

		if err := os.MkdirAll(filepath.Dir(serverFilePath), os.ModePerm); err != nil {
			c.JSON(ServerError, r.DirectoryCreationFailed)
			return
		}

		outputFile, err := os.Create(serverFilePath)
		if err != nil {
			c.JSON(ServerError, r.FileCreationFailed)
			return
		}
		defer outputFile.Close()

		if _, err := outputFile.ReadFrom(c.Request.Body); err != nil {
			c.JSON(ServerError, r.FileSaveFailed)
			return
		}

		if _, err := gifSaver.SaveGif(gifUUID, serverFilePath); err != nil {
			c.JSON(ServerError, r.DatabaseSaveFailed)
			return
		}

		c.JSON(http.StatusOK, Response{
			Response: r.OK(),
			UUID:     gifUUID,
			Path:     serverFilePath,
		})
	}
}
