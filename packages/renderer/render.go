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

	tmplName := "shell"
	tmpl, err := renderTemplate(tmplName)
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

func renderTemplate(tmplName string) (*template.Template, error) {
	dir, err := os.Getwd()
	if err != nil {
		return nil, err
	}

	pathToTemplate := filepath.Join(dir, "packages", "renderer", "templates", tmplName+".html")

	tmpl, err := template.ParseFiles(pathToTemplate)
	if err != nil {
		log.Printf("render template: %v\n", pathToTemplate)

		return nil, err
	}
	return tmpl, nil
}
