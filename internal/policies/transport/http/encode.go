package http

import (
	"bytes"
	"context"
	"encoding/binary"
	"fmt"
	"io"
	"net/http"
	"strconv"

	"github.com/nurdsoft/redesign-grp-trust-portal-api/internal/policies/entities"
)

func encodeGetDocumentResponse(ctx context.Context, w http.ResponseWriter, response interface{}) error {
	docx_document := response.(*entities.GetDocumentResponse)

	if docx_document != nil {
		docx_size := fmt.Sprintf("%d", binary.Size(docx_document.Document))

		w.Header().Set("Content-Transfer-Encoding", "binary")
		w.Header().Set("Content-Disposition", "attachment; filename="+strconv.Quote(docx_document.Name))
		w.Header().Set("Content-Type", "application/vnd.openxmlformats-officedocument.wordprocessingml.document")
		if docx_size != "" {
			w.Header().Set("content-length", docx_size)
		}
		if docx_document.Document != nil {
			_, err := io.Copy(w, bytes.NewReader(docx_document.Document))
			if err != nil {
				return err
			}
		}
		return nil
	}
	return nil
}
