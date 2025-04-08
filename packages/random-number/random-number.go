package random_number

import (
	"context"
	"crypto/rand"
	"log"
	"math"
	"math/big"
)

type Server struct{}

func (s *Server) mustEmbedUnimplementedRandomNumberServer() {
	//TODO implement me
	panic("implement me")
}

func (s *Server) GetRandomNumber(ctx context.Context, message *ReqMessage) (*ResMessage, error) {
	log.Printf("GetRandomNumber %v", message.ReqId)
	rn, err := rand.Int(rand.Reader, big.NewInt(math.MaxInt32))
	if err != nil {
		return nil, err
	}

	return &ResMessage{
		RandNum: rn.Int64(),
	}, nil
}
