package save

import (
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"net/http"
	"os"
	"path/filepath"

	"gopi/internal/lib/response"
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
	response.Response
	UUID string `json:"uuid,omitempty"`
	Path string `json:"path,omitempty"`
}

type GifSaver interface {
	SaveGif(uuid string, path string) (int64, error)
}

func New(gifSaver GifSaver) gin.HandlerFunc {
	return func(c *gin.Context) {
		if c.ContentType() != "application/json" {
			c.JSON(
				BadRequest,
				response.Error("Invalid Content-Type. Expected application/json"),
			)
			return
		}

		var jsonRequest Request
		if err := c.ShouldBindJSON(&jsonRequest); err != nil {
			c.JSON(
				BadRequest,
				response.Error("Invalid JSON: "+err.Error()),
			)
			return
		}

		filePath := jsonRequest.Path
		if _, err := os.Stat(filePath); os.IsNotExist(err) {
			c.JSON(
				BadRequest,
				response.Error("File does not exist at the provided path"),
			)
			return
		}

		newGifUUID := uuid.New().String()

		serverFilePath := filepath.Join("gifs", newGifUUID+".gif")

		err := os.MkdirAll(filepath.Dir(serverFilePath), os.ModePerm)
		if err != nil {
			c.JSON(
				ServerError,
				response.Error("Failed to create directory: "+err.Error()),
			)
			return
		}

		inputFile, err := os.Open(filePath)
		if err != nil {
			c.JSON(
				ServerError,
				response.Error("Failed to open file: "+err.Error()),
			)
			return
		}
		defer inputFile.Close()

		outputFile, err := os.Create(serverFilePath)
		if err != nil {
			c.JSON(
				ServerError,
				response.Error("Failed to create output file: "+err.Error()),
			)
			return
		}
		defer outputFile.Close()

		_, err = outputFile.ReadFrom(inputFile)
		if err != nil {
			c.JSON(
				ServerError,
				response.Error("Failed to save file: "+err.Error()),
			)
			return
		}

		_, err = gifSaver.SaveGif(newGifUUID, serverFilePath)
		if err != nil {
			c.JSON(
				ServerError,
				response.Error("Failed to save GIF: "+err.Error()),
			)
			return
		}

		c.JSON(http.StatusOK, Response{
			Response: response.OK(),
			UUID:     newGifUUID,
			Path:     serverFilePath,
		})
	}
}
