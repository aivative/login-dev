package services

import (
	"context"
	"fmt"

	"github.com/aivative/login-dev/repository"
	"github.com/aivative/login-dev/utils"
)

type AuthSVC struct {
	repo *repository.AuthMongoRepo
}

// NewAuthService Auth service object initializer
func NewAuthService(ctx context.Context, uri string) (userSvc *AuthSVC, err error) {
	userSvc = new(AuthSVC)

	mongoClient, err := utils.NewMongoConnection(ctx, uri)
	if err != nil {
		return nil, fmt.Errorf("mongoConnection: error connecting to the server %v", err)
	}

	userSvc.repo = repository.NewAuthRepo(mongoClient)
	return
}

func (as *AuthSVC) GetUserPasswordSecret(ctx context.Context) (string, error) {
	return as.repo.GetUserPasswordSecret(ctx)
}

func (as *AuthSVC) ValidatePassword(ctx context.Context, password, hashedPassword string) error {
	passwordSecret, err := as.GetUserPasswordSecret(ctx)
	if err != nil {
		return err
	}

	// init hasher
	hasher := utils.NewHasher()

	// validate password TODO: 1. create dummy user with PasswordHash and PasswordSalt [done]
	//                         2. then test it
	if !hasher.LoadSecret(passwordSecret).IsMatch(password, hashedPassword) {
		return fmt.Errorf("invalid password")
	}

	return err
}

func (as *AuthSVC) GetUserAccountSecret(ctx context.Context) (string, error) {
	return as.repo.GetUserAccountSecret(ctx)
}

// func (s *AuthSVC) CreateUser(ctx context.Context, payload model.TCreateUserReq) (*FirebaseAuth.UserRecord, string, error) {
// 	var createPayload FirebaseAuth.UserToCreate
//
// 	if payload.Email != "" {
// 		createPayload.Email(payload.Email)
// 	}
// 	if payload.Picture != "" {
// 		createPayload.PhotoURL(payload.Picture)
// 	}
// 	if payload.EmailVerified != false {
// 		createPayload.EmailVerified(payload.EmailVerified)
// 	}
// 	if payload.Name != "" {
// 		createPayload.DisplayName(payload.Name)
// 	}
//
// 	password := utils.RandStr(14)
//
// 	createPayload.Password(password)
//
// 	userRecord, err := s.client.CreateUser(ctx, &createPayload)
// 	if err != nil {
// 		return userRecord, "", err
// 	}
//
// 	return userRecord, password, nil
// }

// func (s *AuthSVC) SetUserClaim(ctx context.Context, uid, cid string) error {
// 	payload := make(map[string]interface{})
// 	payload["credential_id"] = cid
//
// 	if err := s.client.SetCustomUserClaims(ctx, uid, payload); err != nil {
// 		return err
// 	}
//
// 	return nil
// }
//
// func (s *AuthSVC) UpdateUser(ctx context.Context, uid string, payload model.TUpdateUserReq) (*FirebaseAuth.UserRecord, error) {
// 	if uid == "" {
// 		return nil, utils.WErr("need uid")
// 	}
//
// 	s.client.GetUser()
//
// 	var updateFirebase FirebaseAuth.UserToUpdate
// 	if payload.Name != nil {
// 		updateFirebase.DisplayName(*payload.Name)
// 	}
// 	if payload.EmailVerified != nil {
// 		updateFirebase.EmailVerified(*payload.EmailVerified)
// 	}
// 	if payload.Picture != nil {
// 		updateFirebase.PhotoURL(*payload.Picture)
// 	}
// 	if payload.Password != nil {
// 		updateFirebase.Password(*payload.Password)
// 	}
//
// 	userRecord, err := s.client.UpdateUser(ctx, uid, &updateFirebase)
// 	if err != nil {
// 		return nil, err
// 	}
//
// 	return userRecord, nil
// }
//
// func (s *AuthSVC) DeleteUser(ctx context.Context, uid string) error {
// 	if uid == "" {
// 		return utils.WErr("need uid")
// 	}
//
// 	if err := s.client.DeleteUser(ctx, uid); err != nil {
// 		return err
// 	}
//
// 	return nil
// }
