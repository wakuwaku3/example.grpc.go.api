package main

import (
	"context"
	"errors"

	"example.grpc.go.api/cat"
)

// MyCatService です
type MyCatService struct {
}

// GetMyCat です
func (s *MyCatService) GetMyCat(ctx context.Context, message *cat.GetMyCatMessage) (*cat.MyCatResponse, error) {
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
