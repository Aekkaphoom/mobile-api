package authtoken

import "time"

type AuthToken struct {
	Token string `json:"auth_token"`
	Eid   string `json:"eid"`
}

type TokenExpireResponse struct {
	TokenDateExpire time.Time
}
