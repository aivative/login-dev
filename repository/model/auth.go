package model

type TFirebaseExchangeTokenReq struct {
	Token             string `json:"token"`
	ReturnSecureToken bool   `json:"returnSecureToken"`
}

type TFirebaseExchangeTokenResp struct {
	Token        string `json:"idToken"`
	RefreshToken string `json:"refreshToken"`
}

type TFirebaseRefreshTokenReq struct {
	GrantType    string `json:"grant_type"`
	RefreshToken string `json:"refresh_token"`
}

type TFirebaseRefreshTokenResp struct {
	Token        string `json:"id_token"`
	RefreshToken string `json:"refresh_token"`
}
