package http

import (
	"bytes"
	"context"
	"encoding/binary"
	"fmt"
	"io"
	"net/http"
	"strconv"
	"strings"

	"github.com/nurdsoft/redesign-grp-trust-portal-api/shared/auth"
)

func encodeHtml2DocxResponse(ctx context.Context, w http.ResponseWriter, response interface{}) error {
	ioi := response.([]byte)

	w.Header().Set("Access-Control-Allow-Headers", strings.Join([]string{
		"content-type",
		string(auth.AuthorizationKey),
		string(auth.RedesignTokenKey),
		string(auth.Access),
		"Host",
		"Origin",
	}, ","))
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Content-Transfer-Encoding", "binary")
	w.Header().Set("Content-Disposition", "attachment; filename="+strconv.Quote("fileName.docx"))
	w.Header().Set("Content-Type", "application/vnd.openxmlformats-officedocument.wordprocessingml.document")
	w.Header().Set("content-length", fmt.Sprintf("%d", binary.Size(ioi)))
	_, err := io.Copy(w, bytes.NewReader(ioi))
	if err != nil {
		return err
	}
	return nil
}
