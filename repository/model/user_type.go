package model

type TUserType string

func (ut *TUserType) SetToGarbo() *TUserType {
	*ut = "GARBO"
	return ut
}

func (ut *TUserType) SetToAdmin() *TUserType {
	*ut = "ADMIN"
	return ut
}

func (ut *TUserType) IsGarbo() bool {
	if *ut == "GARBO" {
		return true
	}
	return false
}

func (ut *TUserType) IsAdmin() bool {
	if *ut == "ADMIN" {
		return true
	}
	return false
}
