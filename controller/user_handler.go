package controller

import (
	"fmt"

	"github.com/aivative/login-dev/repository/model"
	"github.com/aivative/login-dev/utils"
	"github.com/gin-gonic/gin"
)

const KIND = "user"

func (ctrl *Controller) Login(c *gin.Context) {
	ctx := c.Request.Context()

	kind := KIND + "#login"
	req := utils.NewReqRespStruct(c, kind)

	// parse
	var loginReq model.TLoginReq
	if err := utils.ParseKindAndBody(c, kind, &loginReq); err != nil {
		req.BadReq("didn't parse correctly", err)
		return
	}

	// 1. Get user by email
	var user model.TUser
	if err := ctrl.UserSVC.GetUserLoginInfo(ctx, loginReq, &user); err != nil && utils.NotFound(err) {
		req.BadReq("user not found", err)
		return
	} else if err != nil {
		req.BadReq("unknown error", err)
		return
	}

	// 2. Validate password
	if err := ctrl.AuthSVC.ValidatePassword(ctx, loginReq.Password, user.PasswordHash); err != nil {
		req.BadReq("invalid", err)
		return
	}

	// 3. Create ID idToken
	idToken, err := ctrl.FirebaseAuthSVC.CreateToken(ctx, user.UserID, user.CredentialID)
	if err != nil {
		req.Unauthorized("can't create token", err)
		return
	}

	req.Ok(idToken)

	// // get jwt secret from db
	// jwtSecret, err := ctrl.AuthSVC.GetUserAccountSecret(ctx)
	// if err != nil {
	// 	req.BadReq("can't get jwt secret", err)
	// 	return
	// }
	//
	// // create token
	// jwtObj := utils.NewJWTObj(jwtSecret)
	// token, err := jwtObj.Sign(map[string]interface{}{
	// 	"credential_id": user.CredentialID,
	// 	"user_id":       user.UserID,
	// }, 3600)
	// if err != nil {
	// 	req.InternalErr("can't create jwt token", err)
	// 	return
	// }
	//
	// newH, err := hasher.GenerateSecret(32)
	// if err != nil {
	// 	req.BadReq("can't generate refresh secret", err)
	// 	return
	// }
	//
	// refreshToken, err := hasher.LoadSecret(newH).Hash(token)
	// if err != nil {
	// 	req.BadReq("can't generate refresh token", err)
	// 	return
	// }
	//
	// // authHeader := strings.SplitAfter(c.GetHeader("Authorization"), "Bearer ")
	// // if len(authHeader) != 2 {
	// // 	req.BadReq("bad authorization header", nil)
	// // 	return
	// // }
	//
	// // TODO: Create jwt token verifier service and repo, also create RefreshToken saver in it
	// // Create token and refresh token, save refresh token on db
	//
	// // idToken, err := ctrl.AuthSVC.C().VerifyIDTokenAndCheckRevoked(ctx, authHeader[1])
	// // if err != nil {
	// // 	req.BadReq("bad req", err)
	// // 	return
	// // }
	// //
	// // logrus.Info(token)
	// //
	// // // TODO: create mew user session repo for managing refresh token
	// var resp model.TLoginResp
	// resp.idToken = token
	// resp.RefreshToken = refreshToken
}

func (ctrl *Controller) RefreshToken(c *gin.Context) {
	// ctx := c.Request.Context()

	kind := KIND + "#refreshtoken"
	req := utils.NewReqRespStruct(c, kind)

	// parse
	var bodyReq model.TRefreshTokenReq
	if err := utils.ParseKindAndBody(c, kind, &bodyReq); err != nil {
		req.BadReq("didn't parse correctly", err)
		return
	}

	// refresh token
	newRefreshToken, err := ctrl.FirebaseAuthSVC.RefreshToken(bodyReq.RefreshToken)
	if err != nil {
		req.Unauthorized("can't refresh token", err)
		return
	}

	req.Ok(newRefreshToken)
}

func (ctrl *Controller) RevokeToken(c *gin.Context) {
	uid := c.Param("uid")
	kind := KIND + "#revoketoken"
	resp := utils.NewReqRespStruct(c, kind)

	if err := ctrl.FirebaseAuthSVC.RevokeRefreshTokens(c.Request.Context(), uid); err != nil {
		resp.BadReq(fmt.Sprintf("can't revoke token to uid %v", uid), err)
		return
	}
	resp.Ok(nil)
}

func (ctrl *Controller) Check(c *gin.Context) {
	c.Header("Authorization", "1234567654321")
	// c.Header("X-Auth-Request-Redirect", "https://google.com")
	c.JSON(200, gin.H{"status": "ok"})
}
