// Package http for conx.
package http

import (
	"context"
	"net/http"

	"github.com/nurdsoft/redesign-grp-trust-portal-api/internal/file_converter/entities"
)

type RequestBodyType interface {
	entities.UploadedFile
}

func decodeDocx2Html(_ context.Context, r *http.Request) (interface{}, error) {
	return &entities.UploadedFile{File: r.Body}, nil
}

func decodeHtml2Docx(_ context.Context, r *http.Request) (interface{}, error) {
	return &entities.UploadedFile{File: r.Body}, nil
}
