package sqlserver

import (
	"database/sql"

	"grouplease.co.th/mobile_api/internal/domain/tbluserslogin"
)

type tblUsersLoginRepository struct {
	db *sql.DB
}

func NewTblUsersLoginRepository(db *sql.DB) tbluserslogin.TblUsersLoginRepository {
	return &tblUsersLoginRepository{
		db: db,
	}
}

// === Implement TblUsersLoginRepository interface ===

func (r *tblUsersLoginRepository) AddUserLogin(userLogin *tbluserslogin.TblUserLoginRequest) error {

	query := `
		INSERT INTO tbl_users_login 
		(auth_token, user_eid, role_access, login_datetime, login_date_expire, created_by, updated_by)
		VALUES 
		(@p1, @p2, @p3, @p4, @p5, @p6, @p7)
	`

	_, err := r.db.Exec(
		query,
		&userLogin.AuthToken,
		&userLogin.UserEid,
		&userLogin.RoleAccess,
		&userLogin.LoginDatetime,
		&userLogin.LoginDateExpire,
		&userLogin.CreatedBy,
		&userLogin.UpdatedBy,
	)
	return err
}

func (r *tblUsersLoginRepository) GetLogin(eid, pass string) (*tbluserslogin.TblUserLogin, error) {
	query := `
		SELECT 
			ul.auth_token, ul.user_eid, ul.role_access, ul.login_datetime, ul.login_date_expire, ul.created_at, ul.created_by, ul.updated_at, ul.updated_by
		FROM tbl_users_login ul
		INNER JOIN tbl_users u on ul.user_eid = u.eid 
		WHERE ul.user_eid = @p1
		AND u.PWORD1 = @p2
	`
	row := r.db.QueryRow(query, &eid, &pass)

	var login tbluserslogin.TblUserLogin
	err := row.Scan(
		&login.AuthToken,
		&login.UserEid,
		&login.RoleAccess,
		&login.LoginDatetime,
		&login.LoginDateExpire,
		&login.CreatedAt,
		&login.CreatedBy,
		&login.UpdatedAt,
		&login.UpdatedBy,
	)

	return &login, err
}

func (r *tblUsersLoginRepository) GetUserLogin(eid, authToken string) (*tbluserslogin.LoginResponse, error) {
	query := `
		SELECT 
			ul.auth_token, u.EID, u.NAM, u.DEPT, u.SID, u.GROUPS, u.SLEVEL
		FROM tbl_users u
		INNER JOIN tbl_users_login ul ON u.EID = ul.user_eid
		WHERE ul.user_eid = @p1
		AND ul.auth_token = @p2
	`
	row := r.db.QueryRow(query, &eid, &authToken)
	var user tbluserslogin.LoginResponse
	err := row.Scan(
		&user.AuthToken,
		&user.Eid,
		&user.Nam,
		&user.Dept,
		&user.Sid,
		&user.Groups,
		&user.Slevel,
	)
	if err != nil {
		return nil, err
	}

	return &user, nil
}

func (r *tblUsersLoginRepository) DeleteUserLoginByEid(userEid string) error {
	query := `
		DELETE FROM tbl_users_login 
		WHERE user_eid = @p1
	`
	_, err := r.db.Exec(query, &userEid)
	return err
}
