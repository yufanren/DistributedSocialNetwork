package etcdclient

import (
	"context"
	"cs9223-final-project/auth/authpb"
	. "cs9223-final-project/etcdclient/etcdclientpb"
	// "errors"
)

type etcdService struct {
	UnimplementedEtcdclientServer
}

func newEtcdService() *etcdService {
	s := etcdService{}
	return &s
}

func (s *etcdService) Login(ctx context.Context, req *AuthRequest) (*AuthResponse, error) {

	err := Login_(req.Username, req.Password)
	if err != nil {
		return &AuthResponse{Token: ""}, err
	}
	token, _ := authClient.GetToken(ctx, &authpb.GetTokenRequest{Username: req.Username})
	return &AuthResponse{Token: token.Token}, nil
}

func (s *etcdService) Register(ctx context.Context, req *AuthRequest) (*AuthResponse, error) {

	if err := Register_(req.Username, req.Password); err != nil {
		return &AuthResponse{Token: ""}, err
	}
	token, _ := authClient.GetToken(ctx, &authpb.GetTokenRequest{Username: req.Username})
	return &AuthResponse{Token: token.Token}, nil
}

func (s *etcdService) Follow(ctx context.Context, req *FollowRequest) (*Response, error) {
	err := Follow_(req.Parent, req.Follower)
	return &Response{}, err
}

func (s *etcdService) IsFollow(ctx context.Context, req *FollowRequest) (*CheckResponse, error) {
	return &CheckResponse{Res: Isfollow_(req.Parent, req.Follower)}, nil
}

func (s *etcdService) GetUserFollowed(ctx context.Context, req *UserRequest) (*ListUserResponse, error) {
	total, users := Getuserfollowed_(string(req.Username))
	return &ListUserResponse{Total: int32(total), Parents: users}, nil
}

func (s *etcdService) GetUserFollower(ctx context.Context, req *UserRequest) (*ListUserResponse, error) {
	total, users := Getuserfollower_(string(req.Username))
	return &ListUserResponse{Total: int32(total), Parents: users}, nil
}


func (s *etcdService) ListBlog(ctx context.Context, req *ListBlogRequest) (*ListBlogResponse, error) {
	total, blogs := Listblog_(int(req.PageNum), int(req.PageSize))
	return &ListBlogResponse{Total: int32(total), Blogs: blogs}, nil
}

func (s *etcdService) AddBlog(ctx context.Context, req *AddBlogRequest) (*AddBlogResponse, error) {
	blogId, err := Addblog_(req.Blog)
	return &AddBlogResponse{BlogId: int32(blogId)}, err
}

func (s *etcdService) GetBlog(ctx context.Context, req *GetBlogRequest) (*GetBlogResponse, error) {
	blog, err := Getblog_(int(req.BlogId))
	if err != nil {
		return &GetBlogResponse{Blog: nil}, err
	}
	return &GetBlogResponse{Blog: blog}, nil
}

func (s *etcdService) GetUserBlog(ctx context.Context, req *UserBlogRequest) (*ListBlogResponse, error) {
	total, blogs := Getuserblog_(string(req.Username), int(req.PageNum), int(req.PageSize))
	return &ListBlogResponse{Total: int32(total), Blogs: blogs}, nil
}


func (s *etcdService) GetFollowBlog(ctx context.Context, req *UserBlogRequest) (*ListBlogResponse, error) {
	total, blogs := Getfollowblog_(string(req.Username), int(req.PageNum), int(req.PageSize))
	return &ListBlogResponse{Total: int32(total), Blogs: blogs}, nil
}