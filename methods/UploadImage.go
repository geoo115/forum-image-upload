package forum

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"path/filepath"
)

// UploadImage handles the upload of images
func UploadImage(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Uploading File\n")

	// 1.Parse Input
	if err := r.ParseMultipartForm(10 << 20); err != nil {
		// Handle the error
		fmt.Println("Error parsing multipart form:", err)
	}

	// 2.Retrive File
	file, handler, err := r.FormFile("myImage")
	if err != nil {
		fmt.Println("Error Retrieving the File")
		fmt.Println(err)
		return
	}
	defer file.Close()

	// Check the allowed file types
	allowedTypes := map[string]bool{
		".jpg":  true,
		".jpeg": true,
		".png":  true,
		".gif":  true,
		// Add more allowed types as needed
	}
	contentType := handler.Header.Get("Content-Type")
	if !allowedTypes[contentType] {
		fmt.Println("Error: Unsupported file type")
		http.Error(w, "Unsupported file type", http.StatusBadRequest)
		return
	}

	fileExt := filepath.Ext(handler.Filename)
	if !allowedTypes[fileExt] {
		fmt.Println("Error: Unsupported file type")
		http.Error(w, "Unsupported file type", http.StatusBadRequest)
		return
	}

	fmt.Printf("Uploaded File: %+v\n", handler.Filename)
	fmt.Printf("Size: %+v\n", handler.Size)
	fmt.Printf("MIME Header: %+v\n", handler.Header)

	// 3.Write Temporary File
	tempFile, err := os.CreateTemp("uploads", "upload-*.png")
	if err != nil {
		fmt.Println("Error creating temporary file:", err)
		return
	}
	defer os.Remove(tempFile.Name())
	fmt.Println("Temporary file:", tempFile.Name())

	// Use io.ReadAll to read the entire contents of the file
	fileBytes, err := io.ReadAll(file)
	if err != nil {
		fmt.Println(err)
		return
	}
	// Use fileBytes as needed
	tempFile.Write(fileBytes)

	// 4.Return whether it was successful or not
	fmt.Printf("Successfully Uploaded File: %v \n", handler.Filename)

}
