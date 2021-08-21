package services

import (
	"context"
	"fmt"

	"github.com/aivative/login-dev/repository"
	"github.com/aivative/login-dev/repository/model"
	"github.com/aivative/login-dev/utils"
)

type UserSvc struct {
	repo *repository.UserMongoRepo
}

// NewUserService User service object initializer
func NewUserService(ctx context.Context, uri string) (userSvc *UserSvc, err error) {
	userSvc = new(UserSvc)

	mongoClient, err := utils.NewMongoConnection(ctx, uri)
	if err != nil {
		return nil, fmt.Errorf("mongoConnection: error connecting to the server %v", err)
	}

	userSvc.repo = repository.NewUserRepo(mongoClient)
	return
}

func (us *UserSvc) GetUserLoginInfo(ctx context.Context, loginReq model.TLoginReq, userResult *model.TUser) (err error) {
	return us.repo.GetUserLoginInfo(ctx, loginReq, userResult)
}

// CRUD

func (us *UserSvc) CreateUser(ctx context.Context, uid string, userPayload model.TCreateUserReq) error {
	return us.repo.CreateUser(ctx, model.TCreateUserQuery{
		TUserID: model.TUserID{UserID: uid},
		TUserProfile: model.TUserProfile{
			Name:       userPayload.Name,
			Picture:    userPayload.Picture,
			DistrictID: userPayload.DistrictID,
			UserType:   userPayload.UserType,
			Email:      userPayload.Email,
		},
	})
}

func (us *UserSvc) GetUser(ctx context.Context, uid string) (model.TGetUserResp, error) {
	return us.repo.GetUser(ctx, uid)
}

func (us *UserSvc) DeleteUser(ctx context.Context, uid string) error {
	return us.repo.DeleteUser(ctx, uid)
}

func (us *UserSvc) UpdateUser(ctx context.Context,uid string, profile model.TUpdateUserReq) error {
	return us.repo.UpdateUser(ctx, uid, profile)
}
