package save

import (
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	r "gopi/internal/lib/response"
	"net/http"
	"os"
	"path/filepath"
	"strings"
)

const (
	BadRequest        = http.StatusBadRequest
	ServerError       = http.StatusInternalServerError
	MaxFileSize int64 = 100 * 1024 * 1024 // 100 MB
)

type Response struct {
	r.Response
	Alias string `json:"alias,omitempty"`
	UUID  string `json:"uuid,omitempty"`
	Path  string `json:"path,omitempty"`
}

type GifSaver interface {
	SaveGif(uuid string, path string) (int64, string, error)
}

func New(gifSaver GifSaver) gin.HandlerFunc {
	return func(c *gin.Context) {
		contentType := c.ContentType()
		if contentType != "multipart/form-data" && contentType != "application/octet-stream" {
			c.JSON(BadRequest, r.InvalidContentType)
			return
		}

		if c.Request.ContentLength > MaxFileSize {
			c.JSON(BadRequest, r.FileTooLarge)
			return
		}

		file, fileHeader, err := c.Request.FormFile("file")
		if err != nil {
			c.JSON(BadRequest, r.FileNotFound)
			return
		}
		defer file.Close()

		if !strings.HasSuffix(strings.ToLower(fileHeader.Filename), ".gif") {
			c.JSON(BadRequest, r.InvalidFileFormat)
			return
		}

		gifUUID := uuid.New().String()
		filePath := filepath.Join("gifs", gifUUID+".gif")

		if err := os.MkdirAll(filepath.Dir(filePath), os.ModePerm); err != nil {
			c.JSON(ServerError, r.DirectoryCreationFailed)
			return
		}

		outputFile, err := os.Create(filePath)
		if err != nil {
			c.JSON(ServerError, r.FileCreationFailed)
			return
		}
		defer outputFile.Close()

		if _, err := outputFile.ReadFrom(c.Request.Body); err != nil {
			c.JSON(ServerError, r.FileSaveFailed)
			return
		}

		_, alias, err := gifSaver.SaveGif(gifUUID, filePath)
		if err != nil {
			c.JSON(ServerError, r.DatabaseSaveFailed)
			return
		}

		c.JSON(http.StatusOK, Response{
			Response: r.OK(),
			Alias:    alias,
			UUID:     gifUUID,
			Path:     filePath,
		})
	}
}
