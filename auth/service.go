package auth

import (
	"context"
	proto "cs9223-final-project/auth/authpb"
)

type authService struct {
	proto.UnimplementedAuthServer
}

func (s *authService) GetToken(ctx context.Context, req *proto.GetTokenRequest) (*proto.GetTokenResponse, error) {
	token, err := GetToken(req.GetUsername())
	if err != nil {
		return nil, err
	}
	return &proto.GetTokenResponse{Token: token}, nil
}

func (s *authService) GetUsername(ctx context.Context, req *proto.GetUsernameRequest) (*proto.GetUsernameResponse, error) {
	username, err := GetUsername(req.Token)
	if err != nil {
		return nil, err
	}
	return &proto.GetUsernameResponse{Username: username}, nil
}
