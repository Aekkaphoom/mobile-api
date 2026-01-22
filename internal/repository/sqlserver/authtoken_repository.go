package sqlserver

import (
	"database/sql"
	"time"

	"grouplease.co.th/mobile_api/internal/domain/authtoken"
)

type authtokenRepository struct {
	db *sql.DB
}

func NewAuthtokenRepository(db *sql.DB) authtoken.AuthtokenRepository {
	return &authtokenRepository{
		db: db,
	}
}

// == Implement interface ===

func (r *authtokenRepository) GetLoginExpireByAuth(req authtoken.AuthToken) (*authtoken.TokenExpireResponse, error) {
	query := `
		SELECT login_date_expire
		FROM tbl_users_login
		WHERE auth_token = @p1
		AND user_eid = @p2
	`
	row := r.db.QueryRow(
		query,
		&req.Token,
		&req.Eid,
	)

	var dateExpire time.Time
	parseErr := row.Scan(&dateExpire)
	if parseErr != nil {
		return nil, parseErr
	}

	var response authtoken.TokenExpireResponse
	response.TokenDateExpire = dateExpire

	return &response, nil
}

func (r *authtokenRepository) FindLoginExpireByQuery(req authtoken.AuthToken) (bool, error) {
	query := `
	SELECT 
		CASE WHEN getdate() > login_date_expire
		THEN 0
		ELSE 1
		END AS 'chkLoginExpire'
	FROM tbl_users_login
	WHERE auth_token = @p1
	AND user_eid = @p2
	`
	row := r.db.QueryRow(
		query,
		&req.Token,
		&req.Eid,
	)
	var expire int
	err := row.Scan(&expire)
	if err != nil {
		return false, err
	}

	// ** expire get 0
	return expire > 0, nil

}
