package entities

import (
	"io"
)

type UploadedFile struct {
	File io.ReadCloser
}

type DocxToHTMLResponse struct {
	HTML string `json:"html"`
}

type HTMLToDocxResponse struct {
	DOCX []byte
}
