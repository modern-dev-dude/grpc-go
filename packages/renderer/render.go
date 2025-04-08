package renderer

import (
	"bytes"
	"context"
	"errors"
	"html/template"
	"log"
	"os"
	"path/filepath"
)

type Server struct{}

func (s *Server) mustEmbedUnimplementedRenderingEngineServer() {
	//TODO implement me
	panic("implement me")
}

func (s *Server) RenderPage(ctx context.Context, message *ReqMessage) (*ResMessage, error) {
	meta := message.Metadata
	if meta == nil {
		return nil, errors.New("no metadata")
	}

	tmpl, err := renderTemplate()
	if err != nil {
		log.Printf("render template err: %v\n", err)
		return nil, err
	}
	buf := bytes.NewBuffer(nil)
	if err := tmpl.Execute(buf, meta); err != nil {
		log.Printf("err: %v", err)
		return nil, err
	}

	return &ResMessage{
		Markup: buf.String(),
	}, nil
}

func constructTemplateFilePath(tmplName string) (string, error) {
	dir, err := os.Getwd()
	if err != nil {
		return "", err
	}
	return filepath.Join(dir, "packages", "renderer", "templates", tmplName+".html"), nil
}

func renderTemplate() (*template.Template, error) {
	shellPath, err := constructTemplateFilePath("shell")
	if err != nil {
		return nil, err
	}
	innerPath, err := constructTemplateFilePath("inner")
	if err != nil {
		return nil, err
	}

	templateList := []string{shellPath, innerPath}

	tmpl, err := template.ParseFiles(templateList...)
	if err != nil {
		log.Printf("render template: %v\n", templateList)
		return nil, err
	}

	return tmpl, nil
}
