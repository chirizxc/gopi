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
		if c.ContentType() != "application/json" {
			c.JSON(BadRequest, r.InvalidContentType)
			return
		}

		var jsonRequest Request
		if err := c.ShouldBindJSON(&jsonRequest); err != nil {
			c.JSON(BadRequest, r.InvalidJSON)
			return
		}

		filePath := jsonRequest.Path
		if _, err := os.Stat(filePath); os.IsNotExist(err) {
			c.JSON(BadRequest, r.FileNotFound)
			return
		}

		newGifUUID := uuid.New().String()
		serverFilePath := filepath.Join("gifs", newGifUUID+".gif")

		if err := os.MkdirAll(filepath.Dir(serverFilePath), os.ModePerm); err != nil {
			c.JSON(ServerError, r.DirectoryCreationFailed)
			return
		}

		inputFile, err := os.Open(filePath)
		if err != nil {
			c.JSON(ServerError, r.FileOpenFailed)
			return
		}
		defer inputFile.Close()

		outputFile, err := os.Create(serverFilePath)
		if err != nil {
			c.JSON(ServerError, r.FileCreationFailed)
			return
		}
		defer outputFile.Close()

		if _, err := outputFile.ReadFrom(inputFile); err != nil {
			c.JSON(ServerError, r.FileSaveFailed)
			return
		}

		if _, err := gifSaver.SaveGif(newGifUUID, serverFilePath); err != nil {
			c.JSON(ServerError, r.DatabaseSaveFailed)
			return
		}

		c.JSON(http.StatusOK, Response{
			Response: r.OK(),
			UUID:     newGifUUID,
			Path:     serverFilePath,
		})
	}
}
