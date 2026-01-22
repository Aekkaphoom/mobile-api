package authtoken

import (
	"grouplease.co.th/mobile_api/internal/domain"
	"grouplease.co.th/mobile_api/internal/domain/authtoken"
	"grouplease.co.th/mobile_api/internal/domain/tblusers"
	"grouplease.co.th/mobile_api/internal/domain/tbluserslogin"
)

type authtokenUsecase struct {
	authRepo  authtoken.AuthtokenRepository
	userRepo  tblusers.TblUsersRepository
	loginRepo tbluserslogin.TblUsersLoginRepository
}

func NewAuthtokenUsecase(
	authRepo authtoken.AuthtokenRepository,
	userRepo tblusers.TblUsersRepository,
	loginRepo tbluserslogin.TblUsersLoginRepository,
) authtoken.AuthtokenUsecase {
	return &authtokenUsecase{
		authRepo:  authRepo,
		userRepo:  userRepo,
		loginRepo: loginRepo,
	}
}

// == Implement interface ===

func (r authtokenUsecase) GetDatetimeLoginExpireByAuth(req authtoken.AuthToken) (*authtoken.TokenExpireResponse, error) {
	// == Check req value ==
	if req.Eid == "" || req.Token == "" {
		return nil, domain.ErrAuthTokenMissing
	}

	// == Get Token from login ==
	response, err := r.authRepo.GetLoginExpireByAuth(req)
	if err != nil {
		return nil, err
	}

	return response, nil
}

func (r authtokenUsecase) GetExpireAccess(req authtoken.AuthToken) (bool, error) {
	if req.Eid == "" || req.Token == "" {
		return false, domain.ErrAuthTokenMissing
	}

	tokenExpire, err := r.authRepo.FindLoginExpireByQuery(req)
	if err != nil {
		return false, err
	}

	if !tokenExpire {
		return false, domain.ErrAuthTokenExpire
	}

	return true, nil
}
