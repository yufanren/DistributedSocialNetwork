package backend

import (

"context"
"cs9223-final-project/auth/authpb"
"errors"
"final-project-0b5a2e16-ly2062-yufanren/backend/backendpb"
)

type backendService struct {
	backendpb.UnimplementedBackendServer
	db *DB
}

func newBackendService() *backendService {
	s := backendService{}
	s.db = NewDB()
	return &s
}

func (s *backendService) Login(ctx context.Context, req *backendpb.AuthRequest) (*backendpb.AuthResponse, error) {
	if !s.db.HasUser(req.Username) {
		return &backendpb.AuthResponse{Token: ""}, errors.New("user does not exist")
	}
	if !s.db.ValidUser(req.Username, req.Password) {
		return &backendpb.AuthResponse{Token: ""}, errors.New("wrong password")
	}
	token, _ := authClient.GetToken(ctx, &authpb.GetTokenRequest{Username: req.Username})
	return &backendpb.AuthResponse{Token: token.Token}, nil
}

func (s *backendService) Register(ctx context.Context, req *backendpb.AuthRequest) (*backendpb.AuthResponse, error) {
	if err := s.db.AddUser(req.Username, req.Password); err != nil {
		return &backendpb.AuthResponse{Token: ""}, err
	}
	token, _ := authClient.GetToken(ctx, &authpb.GetTokenRequest{Username: req.Username})
	return &backendpb.AuthResponse{Token: token.Token}, nil
}

func (s *backendService) ListBlog(ctx context.Context, req *backendpb.ListBlogRequest) (*backendpb.ListBlogResponse, error) {
	total, blogs := s.db.ListBlogs(int(req.PageNum), int(req.PageSize))
	return &backendpb.ListBlogResponse{Total: int32(total), Blogs: blogs}, nil
}

func (s *backendService) AddBlog(ctx context.Context, req *backendpb.AddBlogRequest) (*backendpb.AddBlogResponse, error) {
	return &backendpb.AddBlogResponse{BlogId: int32(s.db.AddBlog(req.Blog))}, nil
}

func (s *backendService) GetBlog(ctx context.Context, req *backendpb.GetBlogRequest) (*backendpb.GetBlogResponse, error) {
	return &backendpb.GetBlogResponse{Blog: s.db.GetBlog(int(req.BlogId))}, nil
}

func (s *backendService) Follow(ctx context.Context, req *backendpb.FollowRequest) (*backendpb.Response, error) {
	s.db.FollowUser(req.Parent, req.Follower)
	return &backendpb.Response{}, nil
}

func (s *backendService) IsFollow(ctx context.Context, req *backendpb.FollowRequest) (*backendpb.CheckResponse, error) {
	return &backendpb.CheckResponse{Res: s.db.isFollow(req.Parent, req.Follower)}, nil
}

func (s *backendService) GetUserBlog(ctx context.Context, req *backendpb.UserBlogRequest) (*backendpb.ListBlogResponse, error) {
	total, blogs := s.db.ListBlogsFromUser(string(req.Username), int(req.PageNum), int(req.PageSize))
	return &backendpb.ListBlogResponse{Total: int32(total), Blogs: blogs}, nil
}

func (s *backendService) GetUserFollowed(ctx context.Context, req *backendpb.UserRequest) (*backendpb.ListUserResponse, error) {
	total, users := s.db.ListFollowedFromUser(string(req.Username))
	return &backendpb.ListUserResponse{Total: int32(total), Parents: users}, nil
}

func (s *backendService) GetUserFollower(ctx context.Context, req *backendpb.UserRequest) (*backendpb.ListUserResponse, error) {
	total, users := s.db.ListFollowerFromUser(string(req.Username))
	return &backendpb.ListUserResponse{Total: int32(total), Parents: users}, nil
}

func (s *backendService) GetFollowBlog(ctx context.Context, req *backendpb.UserBlogRequest) (*backendpb.ListBlogResponse, error) {
	total, blogs := s.db.ListFollowBlog(string(req.Username), int(req.PageNum), int(req.PageSize))
	return &backendpb.ListBlogResponse{Total: int32(total), Blogs: blogs}, nil
}