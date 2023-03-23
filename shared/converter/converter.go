package converter

import (
	"fmt"
	"io"
	"log"
	"os"
	"os/exec"
)

const (
	DOCX               = ".docx"
	HTML               = ".html"
	TEMPORARY_FILENAME = "temp"
)

var TEMPORARY_DOCX_FILE string = fmt.Sprintf("temp%s", DOCX)
var TEMPORARY_HTML_FILE string = fmt.Sprintf("temp%s", HTML)

type File struct {
	From string
	To   string
	Data io.ReadCloser
}

func removeFile(path string) {
	if _, err := os.Stat(path); err == nil {
		// File exists, let's remove it
		err := os.Remove(path)
		if err != nil {
			log.Println(err)
		}
	}
}

// Remove temporary files used for conversion
func Cleanup() {
	removeFile(TEMPORARY_DOCX_FILE)
	removeFile(TEMPORARY_HTML_FILE)
}

func (f *File) DocxToHTML() (string, error) {
	// Create a temporary file with uploaded file
	out, err := os.Create(fmt.Sprintf("temp%s", f.From))
	if err != nil {
		return "", err
	}
	defer out.Close()
	_, err = io.Copy(out, f.Data)
	if err != nil {
		return "", err
	}

	// invoke pandoc to do conversion
	cmd := exec.Command("pandoc", fmt.Sprintf("temp%s", f.From), "-o", fmt.Sprintf("temp%s", f.To))
	err = cmd.Run()
	if err != nil {
		return "", err
	}

	dat, err := os.ReadFile(fmt.Sprintf("temp%s", f.To))
	if err != nil {
		return "", err
	}
	defer Cleanup()
	return string(dat), nil
}

// Convert a HTML file to Word Docx format
func (f *File) HTMLToDocx() ([]byte, error) {
	// Create a temporary file with uploaded file
	out, err := os.Create(fmt.Sprintf("temp%s", f.From))
	if err != nil {
		return nil, err
	}
	defer out.Close()
	_, err = io.Copy(out, f.Data)
	if err != nil {
		return nil, err
	}
	// invoke pandoc to do conversion
	cmd := exec.Command("pandoc", fmt.Sprintf("temp%s", f.From), "-o", fmt.Sprintf("temp%s", f.To))
	err = cmd.Run()
	if err != nil {
		return nil, err
	}

	data, err := os.ReadFile(fmt.Sprintf("temp%s", f.To))
	if err != nil {
		return nil, err
	}
	defer Cleanup()
	return data, nil
}
