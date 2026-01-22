package tbluserslogin

type TblUsersLoginUseCase interface {
	AddUserLogin(loginReq *TblUserLoginRequest) error
	DeleteUserLoginByEid(userEid string) error
	GetUserLogin(eid, token string) (*LoginResponse, error)
	GetLogin(eid, pass string) (*LoginResponse, error)
}
