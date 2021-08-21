package services

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"net/http"

	firebase "firebase.google.com/go"
	FirebaseAuth "firebase.google.com/go/auth"
	"github.com/aivative/login-dev/config"
	"github.com/aivative/login-dev/repository/model"
)

func NewFirebaseAuthService(ctx context.Context) (authSvc *FirebaseAuthSvc, err error) {
	authSvc = new(FirebaseAuthSvc)

	app, err := firebase.NewApp(ctx, nil)
	if err != nil {
		return nil, fmt.Errorf("NewApp: error initializing app %v", err)
	}

	if authSvc.client, err = app.Auth(ctx); err != nil {
		return nil, fmt.Errorf("authSvc: error getting auth client %v", err)
	}

	return
}

type FirebaseAuthSvc struct {
	client *FirebaseAuth.Client
}

func (a *FirebaseAuthSvc) C() *FirebaseAuth.Client {
	return a.client
}

func (a *FirebaseAuthSvc) CreateToken(ctx context.Context, uid, credentialId string) (token model.TFirebaseExchangeTokenResp, err error) {
	// Generate Custom Token
	customToken, err := a.client.CustomTokenWithClaims(ctx, uid, map[string]interface{}{"credential_id": credentialId})

	// Exchange Custom Token with ID Token + Refresh Token
	token, err = a.ExchangeCustomToken(customToken)
	if err != nil {
		return
	}

	return
}

func (a *FirebaseAuthSvc) RefreshToken(refreshToken string) (token model.TFirebaseRefreshTokenResp, err error) {

	// TODO: create http caller, or find existing package that can solve these overhead
	key := config.APIKeyConf["default"].Key

	idTokenReq := model.TFirebaseRefreshTokenReq{
		GrantType:    "refresh_token",
		RefreshToken: refreshToken,
	}

	reqBody, err := json.Marshal(idTokenReq)
	if err != nil {
		return
	}

	strReqBody := bytes.NewBuffer(reqBody)

	bodyResp, err := http.Post(
		fmt.Sprintf("https://securetoken.googleapis.com/v1/token?key=%v", key),
		"application/json",
		strReqBody,
	)
	if err != nil {
		return
	}

	if bodyResp.StatusCode != http.StatusOK {
		return model.TFirebaseRefreshTokenResp{}, fmt.Errorf("invalid token")
	}

	if err = json.NewDecoder(bodyResp.Body).Decode(&token); err != nil {
		return
	}

	return

}

func (a *FirebaseAuthSvc) RevokeRefreshTokens(ctx context.Context, idToken string) error {
	return a.client.RevokeRefreshTokens(ctx, idToken)
}

func (a *FirebaseAuthSvc) ExchangeCustomToken(customToken string) (token model.TFirebaseExchangeTokenResp, err error) {

	key := config.APIKeyConf["default"].Key

	idTokenReq := model.TFirebaseExchangeTokenReq{
		Token:             customToken,
		ReturnSecureToken: true,
	}

	reqBody, err := json.Marshal(idTokenReq)
	if err != nil {
		return
	}

	strReqBody := bytes.NewBuffer(reqBody)

	bodyResp, err := http.Post(
		fmt.Sprintf("https://identitytoolkit.googleapis.com/v1/accounts:signInWithCustomToken?key=%v", key),
		"application/json",
		strReqBody,
	)
	if err != nil {
		return
	}

	if bodyResp.StatusCode != http.StatusOK {
		return model.TFirebaseExchangeTokenResp{}, fmt.Errorf("invalid")
	}

	if err = json.NewDecoder(bodyResp.Body).Decode(&token); err != nil {
		return
	}

	return
}

func (a *FirebaseAuthSvc) GetUserRecord(ctx context.Context, uid string) (*FirebaseAuth.UserRecord, error) {
	userRecord, err := a.client.GetUser(ctx, uid)
	if err != nil {
		return nil, err
	}

	return userRecord, nil
}
