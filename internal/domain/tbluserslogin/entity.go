package tbluserslogin

import "time"

type TblUserLogin struct {
	AuthToken       string    `db:"auth_token" json:"auth_token"`
	UserEid         string    `db:"user_eid" json:"user_eid"`
	RoleAccess      string    `db:"role_access" json:"role_access"`
	LoginDatetime   time.Time `db:"login_datetime" json:"login_datetime"`
	LoginDateExpire time.Time `db:"login_date_expire" json:"login_date_expire"`
	CreatedAt       string    `db:"created_at" json:"created_at"`
	CreatedBy       string    `db:"created_by" json:"created_by"`
	UpdatedAt       string    `db:"updated_at" json:"updated_at"`
	UpdatedBy       string    `db:"updated_by" json:"updated_by"`
}

type TblUserLoginRequest struct {
	AuthToken       string    `json:"auth_token"`
	UserEid         string    `json:"user_eid"`
	RoleAccess      string    `json:"role_access"`
	LoginDatetime   time.Time `json:"login_datetime"`
	LoginDateExpire time.Time `json:"login_date_expire"`
	CreatedBy       string    `json:"created_by"`
	UpdatedBy       string    `json:"updated_by"`
}

type LoginRequest struct {
	EID      string `json:"eid" validate:"required"`
	Password string `json:"password" validate:"required"`
}

type LoginResponse struct {
	AuthToken string `json:"auth_token"`
	Eid       string `json:"eid"`
	Nam       string `json:"nam"`
	Dept      string `json:"dept"`
	Sid       string `json:"sid"`
	Slevel    string `json:"slevel"`
	Groups    int    `json:"groups"`
}
