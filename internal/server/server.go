package server

import (
	"context"
	"fmt"

	"github.com/twitchtv/twirp"
	hd "github.com/wung-s/hellotwirp/rpc/haberdasher"
	hw "github.com/wung-s/hellotwirp/rpc/helloworld"
)

// Server implements the Haberdasher service
type Server struct{}

func (s *Server) MakeHat(ctx context.Context, size *hd.Size) (hat *hd.Hat, err error) {
	if size.Inches <= 0 {
		return nil, twirp.InvalidArgumentError("inches", "I can't make a hat that small!")
	}
	return &hd.Hat{Inches: size.Inches + 2}, nil
}

func (s *Server) Hello(ctx context.Context, req *hw.HelloReq) (resp *hw.HelloResp, err error) {
	return &hw.HelloResp{Text: fmt.Sprintf("Subject: %s, this is service B !", req.Subject)}, nil
}
