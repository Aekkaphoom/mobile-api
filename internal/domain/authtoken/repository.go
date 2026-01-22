package authtoken

type AuthtokenRepository interface {
	GetLoginExpireByAuth(req AuthToken) (*TokenExpireResponse, error)
	FindLoginExpireByQuery(req AuthToken) (bool, error)
}
