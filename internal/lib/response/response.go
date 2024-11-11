package response

type Response struct {
	Status string `json:"status"`
	Error  string `json:"error,omitempty"`
}

const (
	StatusOK    = "OK"
	StatusError = "Error"
)

func OK() Response {
	return Response{
		Status: StatusOK,
	}
}

func Error(msg string) Response {
	return Response{
		Status: StatusError,
		Error:  msg,
	}
}

var (
	InvalidContentType      = Error("Invalid Content-Type. Expected application/json")
	InvalidJSON             = Error("Invalid JSON format")
	FileNotFound            = Error("File does not exist at the provided path")
	DirectoryCreationFailed = Error("Failed to create directory")
	FileOpenFailed          = Error("Failed to open file")
	FileCreationFailed      = Error("Failed to create output file")
	FileSaveFailed          = Error("Failed to save file")
	DatabaseSaveFailed      = Error("Failed to save GIF to database")
)
