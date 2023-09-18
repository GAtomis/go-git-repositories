package service

import (
	"context"
	"fmt"
	"github.com/go-kratos/kratos/v2/metadata"
	"go-git-repositories/helper"
	"go-git-repositories/models"

	pb "go-git-repositories/api/git"
)

type UserService struct {
	pb.UnimplementedUserServer
}

func NewUserService() *UserService {
	return &UserService{}
}

func (s *UserService) Login(ctx context.Context, req *pb.LoginRequest) (*pb.LoginReply, error) {
	var username, identity string
	if md, ok := metadata.FromServerContext(ctx); ok {
		username = md.Get("username")
		identity = md.Get("identity")
		fmt.Println(username)
		fmt.Println(identity)
	}
	ub := new(models.UserBasic)
	db := models.GAA_SQL.GetDB()
	err := db.Where("username=? And password = ?", req.Username, helper.GetMd5(req.Password)).First(ub).Error
	fmt.Println(&ub)
	if err != nil {
		return nil, err
	}
	fmt.Println(ub.Identity, ub.Username)
	token, err := helper.GenerateToken(ub.Identity, ub.Username)
	if err != nil {
		return nil, err
	}
	return &pb.LoginReply{
		Token: token,
	}, nil
}

func (s *UserService) CreateUser(ctx context.Context, req *pb.CreateUserRequest) (*pb.CreateUserReply, error) {
	ub := new(models.UserBasic)
	ub.Username = req.Username
	ub.Password = helper.GetMd5(req.Password)
	fmt.Println(ub)
	db := models.GAA_SQL.GetDB()
	err := db.Create(ub).Error
	if err != nil {
		return nil, err
	}
	return &pb.CreateUserReply{
		Msg: "成功",
	}, nil

}

func (s *UserService) GetInfo(ctx context.Context, req *pb.InfoRequest) (*pb.InfoReply, error) {
	return &pb.InfoReply{Msg: "Hello " + req.Name}, nil
}
