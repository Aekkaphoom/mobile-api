package authtoken

type AuthtokenUsecase interface {
	GetDatetimeLoginExpireByAuth(req AuthToken) (*TokenExpireResponse, error)
	GetExpireAccess(req AuthToken) (bool, error)
}
