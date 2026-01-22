package tbluserslogin_usecase

import (
	"strings"

	"grouplease.co.th/mobile_api/internal/domain"
	"grouplease.co.th/mobile_api/internal/domain/tbluserslogin"
)

type tblUsersLoginUsecase struct {
	loginRepo tbluserslogin.TblUsersLoginRepository
}

func NewUsersLoginUsecase(loginRepo tbluserslogin.TblUsersLoginRepository) tbluserslogin.TblUsersLoginUseCase {
	return &tblUsersLoginUsecase{
		loginRepo: loginRepo,
	}
}

// === Implement TblUsersLoginUsecase interface ===

func (u *tblUsersLoginUsecase) AddUserLogin(userLogin *tbluserslogin.TblUserLoginRequest) error {
	return u.loginRepo.AddUserLogin(userLogin)
}

func (u *tblUsersLoginUsecase) DeleteUserLoginByEid(userEid string) error {
	return u.loginRepo.DeleteUserLoginByEid(userEid)
}

func (u *tblUsersLoginUsecase) GetUserLogin(eid, token string) (*tbluserslogin.LoginResponse, error) {
	if strings.TrimSpace(eid) == "" || strings.TrimSpace(token) == "" {
		return nil, domain.ErrNotFound
	}

	user, err := u.loginRepo.GetUserLogin(eid, token)
	if err != nil {
		return nil, err
	}

	return user, nil
}

func (u *tblUsersLoginUsecase) GetLogin(eid, pass string) (*tbluserslogin.LoginResponse, error) {
	if strings.TrimSpace(eid) == "" || strings.TrimSpace(pass) == "" {
		return nil, domain.ErrNotFound
	}

	login, err := u.loginRepo.GetLogin(eid, pass)
	if err != nil {
		return nil, err
	}

	if login != nil {
		userLogin, err := u.loginRepo.GetUserLogin(login.UserEid, login.AuthToken)
		if err != nil {
			return nil, err
		}
		return userLogin, nil
	}

	return nil, nil
}
