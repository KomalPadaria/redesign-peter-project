package service

import (
	"context"

	"github.com/nurdsoft/redesign-grp-trust-portal-api/internal/file_converter/entities"
	"github.com/nurdsoft/redesign-grp-trust-portal-api/shared/converter"
)

type Service interface {
	ConvertDocxToHtml(ctx context.Context, req *entities.UploadedFile) (*entities.DocxToHTMLResponse, error)
	ConvertHtmlToDocx(ctx context.Context, req *entities.UploadedFile) (interface{}, error)
}

type service struct {
	file converter.File
}

func (s *service) ConvertDocxToHtml(ctx context.Context, req *entities.UploadedFile) (*entities.DocxToHTMLResponse, error) {
	s.file.Data = req.File
	s.file.From = converter.DOCX
	s.file.To = converter.HTML

	html, err := s.file.DocxToHTML()
	if err != nil {
		return nil, err
	}
	if len(html) > 0 {
		return &entities.DocxToHTMLResponse{HTML: html}, nil
	}
	return nil, nil
}

func (s *service) ConvertHtmlToDocx(ctx context.Context, req *entities.UploadedFile) (interface{}, error) {
	s.file.Data = req.File
	s.file.From = converter.HTML
	s.file.To = converter.DOCX

	docx, err := s.file.HTMLToDocx()
	if err != nil {
		return nil, err
	}
	return docx, nil
}

func New(file converter.File) Service {
	return &service{
		file: file,
	}
}
