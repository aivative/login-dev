package model

// TPasswordHashSalt is password hash structure data for user
type TPasswordHashSalt struct {
	PasswordHash string `json:"password_hash" bson:"password_hash"`
	PasswordSalt string `json:"password_salt" bson:"password_salt"`
}

// TUserProfile is user profile structure data for user
type TUserProfile struct {
	Name       string    `json:"name" bson:"name"`
	Picture    string    `json:"picture" bson:"picture"`
	DistrictID string    `json:"district_id" bson:"district_id"`
	UserType   TUserType `json:"user_type" bson:"user_type"`
	Email      string    `json:"email" bson:"email"`
}

// TUserID is user id structure data for user
type TUserID struct {
	UserID string `json:"user_id" bson:"user_id"`
}

// TUser is base structure data for user
type TUser struct {
	CredentialID string `json:"credential_id" bson:"credential_id"`

	TUserID           `json:",inline" bson:",inline"`
	TUserProfile      `json:",inline" bson:",inline"`
	TPasswordHashSalt `json:",inline" bson:",inline"`
	TTimeAttributes   `json:",inline" bson:",inline"`
}

// ### LOGIN USER ###

// TLoginReq is  login request structure data
type TLoginReq struct {
	Email    string `json:"email" bson:"email" validate:"required"`
	Password string `json:"password" validate:"required"`
}

// TLoginResp is a login response structure data
type TLoginResp struct {
	Token        string `json:"token" bson:"token"`
	RefreshToken string `json:"refresh_token" bson:"refresh_token"`
}

// ### REFRESH TOKEN ###

// TRefreshTokenReq is Refresh token request structure data
type TRefreshTokenReq struct {
	RefreshToken string `json:"refresh_token" bson:"refresh_token" binding:"required"`
}

// TRefreshTokenResp is Refresh token response structure data
type TRefreshTokenResp struct {
	Token        string `json:"token" bson:"token"`
	RefreshToken string `json:"refresh_token" bson:"refresh_token"`
}

// ### CREATE USER ###

// TCreateUserReq is create user request structure data
type TCreateUserReq struct {
	Name       string    `json:"name" bson:"name" validate:"required"`
	Picture    string    `json:"picture" bson:"picture" validate:"required"`
	DistrictID string    `json:"district_id" bson:"district_id" validate:"required"`
	UserType   TUserType `json:"user_type" bson:"user_type" validate:"required"`
	Email      string    `json:"email" bson:"email" validate:"required"`
}

type TCreateUserQuery struct {
	TUserID      `json:",inline" bson:",inline"`
	TUserProfile `json:",inline" bson:",inline"`
}

// TCreateUserResp is  create user response structure data
type TCreateUserResp struct {
	TUserID      `json:",inline" bson:",inline"`
	TUserProfile `json:",inline" bson:",inline"`
}

// ### GET USER ###

type TGetUserResp struct {
	TUserID      `json:",inline" bson:",inline"`
	TUserProfile `json:",inline" bson:",inline"`
}

// ### UPDATE USER ###

type TUpdateUserReq struct {
	Name       *string    `json:"name" bson:"name"`
	Picture    *string    `json:"picture" bson:"picture"`
	DistrictID *string    `json:"district_id" bson:"district_id"`
	UserType   *TUserType `json:"user_type" bson:"user_type"`
	Email      *string    `json:"email" bson:"email"`
}
