package server

import (
	"database/sql"
	"errors"
	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
	"github.com/google/uuid"
	"gopi/internal/config"
	"log"
	"net/http"
	"os"
	"path/filepath"
)

var DB *sql.DB

type Gif struct {
	ID   int    `json:"id"`
	UUID string `json:"uuid"`
	Path string `json:"path"`
}

func init() {
	cfg := config.LoadConfig()
	dsn := cfg.Database.Dsn

	var err error
	DB, err = sql.Open("mysql", dsn)
	if err != nil {
		log.Fatal("Error opening database: ", err)
	}

	err = DB.Ping()
	if err != nil {
		log.Fatal("Error pinging database: ", err)
	}

	DB.SetMaxOpenConns(10)
	DB.SetMaxIdleConns(5)
	DB.SetConnMaxLifetime(0)
}

func getGifs(c *gin.Context) {
	rows, err := DB.Query("SELECT id, uuid, path FROM gifs")
	if err != nil {
		log.Println("Error reading data:", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error reading data"})
		return
	}
	defer rows.Close()

	var gifs []Gif
	for rows.Next() {
		var gif Gif
		if err := rows.Scan(&gif.ID, &gif.UUID, &gif.Path); err != nil {
			log.Println("Error processing data:", err)
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Error processing data"})
			return
		}
		gifs = append(gifs, gif)
	}

	if err := rows.Err(); err != nil {
		log.Println("Error with data:", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error with data"})
		return
	}

	c.JSON(http.StatusOK, gifs)
}

func createGif(c *gin.Context) {
	file, err := c.FormFile("gif")
	if err != nil {
		log.Println("Error uploading file:", err)
		c.JSON(http.StatusBadRequest, gin.H{"error": "Error uploading file"})
		return
	}

	newGif := Gif{
		UUID: uuid.New().String(),
	}

	filePath := filepath.Join("gifs", newGif.UUID+".gif")

	err = os.MkdirAll(filepath.Dir(filePath), os.ModePerm)
	if err != nil {
		log.Println("Error creating directory:", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error saving file"})
		return
	}

	err = c.SaveUploadedFile(file, filePath)
	if err != nil {
		log.Println("Error saving file:", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error saving file"})
		return
	}

	_, err = DB.Exec("INSERT INTO gifs (uuid, path) VALUES (?, ?)", newGif.UUID, filePath)
	if err != nil {
		log.Println("Error adding gif:", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error adding gif"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "GIF added", "uuid": newGif.UUID, "path": filePath})
}

func getGifFile(c *gin.Context) {
	gifUUID := c.Param("uuid")

	var gif Gif
	err := DB.QueryRow("SELECT uuid, path FROM gifs WHERE uuid = ?", gifUUID).Scan(&gif.UUID, &gif.Path)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			c.JSON(http.StatusNotFound, gin.H{"error": "GIF not found"})
		} else {
			log.Println("Error retrieving gif data:", err)
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Error retrieving gif data"})
		}
		return
	}

	if _, err := os.Stat(gif.Path); os.IsNotExist(err) {
		c.JSON(http.StatusNotFound, gin.H{"error": "File not found"})
		return
	}

	c.File(gif.Path)
}

func NewServer() *gin.Engine {
	r := gin.Default()
	r.GET("/gifs", getGifs)
	r.POST("/create", createGif)
	r.GET("/files/:uuid", getGifFile)
	return r
}
