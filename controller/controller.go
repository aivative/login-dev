package controller

import (
	"context"

	"github.com/aivative/login-dev/config"
	"github.com/aivative/login-dev/services"
)

// Controller is a controller object that have services
type Controller struct {
	// mongoClient *mongo.Client
	UserSVC         *services.UserSvc
	AuthSVC         *services.AuthSVC
	FirebaseAuthSVC *services.FirebaseAuthSvc
}

// New generates new controller object
func New(ctx context.Context) (ctrl *Controller, err error) {
	ctrl = new(Controller)

	uri := config.MongoConf["mongo-user-auth"].URI
	if ctrl.AuthSVC, err = services.NewAuthService(ctx, uri); err != nil {
		return
	}

	uri = config.MongoConf["mongo-user"].URI
	if ctrl.UserSVC, err = services.NewUserService(ctx, uri); err != nil {
		return
	}

	if ctrl.FirebaseAuthSVC, err = services.NewFirebaseAuthService(ctx); err != nil {
		return
	}

	return
}
