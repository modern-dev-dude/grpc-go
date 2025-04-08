package renderer

import (
	"context"
	"errors"
	v8 "rogchap.com/v8go"
	"sync"
)

type Server struct{}

func (s *Server) RenderPage(ctx context.Context, message *ReqMessage) (*ResMessage, error) {
	// setup v8 runtime
	node := NewNodeCtx()

	meta := message.Metadata
	if meta == nil {
		return nil, errors.New("no metadata")
	}

	return &ResMessage{
		Data: "",
	}, nil
}

type NodeCtx *v8.Context

var once sync.Once
var nodeInstance NodeCtx

func NewNodeCtx() NodeCtx {
	once.Do(func() {
		nodeInstance = v8.NewContext()
	})

	return nodeInstance
}
