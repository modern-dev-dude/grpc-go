package random_number

import (
	"context"
	"crypto/rand"
	"math"
	"math/big"
)

type Server struct{}

func (s *Server) mustEmbedUnimplementedRenderingEngineServer() {
	//TODO implement me
	panic("implement me")
}

func (s *Server) RenderPage(ctx context.Context, message *ReqMessage) (*ResMessage, error) {
	rn, err := rand.Int(rand.Reader, big.NewInt(math.MaxInt32))
	if err != nil {
		return nil, err
	}

	return &ResMessage{
		RandNum: rn.Int64(),
	}, nil
}
