package converter

import (
	"bytes"
	"io"
	"os"
	"strings"
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

func Test_Docx2Html(t *testing.T) {
	expected := "<p>this is a <strong>paragraph</strong></p>\n"

	Convey("Given a docx file", t, func() {
		docx_file, _ := os.ReadFile("sample.docx")
		f := &File{
			From: DOCX,
			To:   HTML,
			Data: io.NopCloser(bytes.NewReader(docx_file)),
		}
		Convey("Call the DocxToHTML method", func() {
			html, err := f.DocxToHTML()
			Convey("The html returned should be valid and expected", func() {
				So(err, ShouldBeNil)
				So(html, ShouldNotBeNil)
				So(html, ShouldResemble, expected)
			})
		})
	})
}

func Test_Html2Docx(t *testing.T) {
	Convey("Given a HTML string", t, func() {
		html := "<p>this is a <strong>paragraph</strong></p>\n"
		f := &File{
			From: HTML,
			To:   DOCX,
			Data: io.NopCloser(strings.NewReader(html)),
		}
		Convey("Call HTML2Docx method", func() {
			docx, err := f.HTMLToDocx()
			// write data to a file
			_ = os.WriteFile("data.docx", docx, 0644)
			// re-read it again
			docx_file, _ := os.ReadFile("data.docx")

			Convey("The docx returned should be valid and expected", func() {
				So(err, ShouldBeNil)
				So(docx, ShouldNotBeNil)
				So(docx, ShouldResemble, docx_file)
				os.Remove("data.docx")
			})
		})
	})
}
