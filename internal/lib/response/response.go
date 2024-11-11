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
	InvalidContentType      = Error("Invalid Content-Type. Expected multipart/form-data or application/octet-stream")
	DirectoryCreationFailed = Error("Failed to create directory for saving the GIF")
	FileCreationFailed      = Error("Failed to create output file for GIF")
	FileSaveFailed          = Error("Failed to save binary file content")
	DatabaseSaveFailed      = Error("Failed to save GIF metadata to the database")
	FileTooLarge            = Error("File is too large. Maximum size is 100 MB")
	InvalidFileFormat       = Error("Invalid file format. Only .gif files are allowed")
	FileNotFound            = Error("File not found")
)
