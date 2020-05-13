package main

import (
	"context"
	"errors"
	"log"

	"github.com/wakuwaku3/example.grpc.go.api/cat"
	"google.golang.org/grpc/metadata"
)

// MyCatService です
type MyCatService struct {
}

// GetMyCat です
func (s *MyCatService) GetMyCat(ctx context.Context, message *cat.GetMyCatMessage) (*cat.MyCatResponse, error) {
	if md, ok := metadata.FromIncomingContext(ctx); ok {
		log.Print(md)
	}
	switch message.TargetCat {
	case "tama":
		//たまはメインクーン
		return &cat.MyCatResponse{
			Name: "tama",
			Kind: "mainecoon",
		}, nil
	case "mike":
		//ミケはノルウェージャンフォレストキャット
		return &cat.MyCatResponse{
			Name: "mike",
			Kind: "Norwegian Forest Cat",
		}, nil
	}
	return nil, errors.New("Not Found YourCat")
}
