package tbluserslogin

type TblUsersLoginRepository interface {
	AddUserLogin(userLogin *TblUserLoginRequest) error
	GetLogin(eid, pass string) (*TblUserLogin, error)
	GetUserLogin(eid, authToken string) (*LoginResponse, error)
	DeleteUserLoginByEid(userEid string) error
}
