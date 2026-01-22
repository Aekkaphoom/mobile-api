package sqlserver

import (
	"context"
	"database/sql"

	"grouplease.co.th/mobile_api/internal/domain/tblusers"
)

type tblUsersRepository struct {
	db *sql.DB
}

func NewTblUsersRepository(db *sql.DB) tblusers.TblUsersRepository {
	return &tblUsersRepository{
		db: db,
	}
}

// === Implement TblUsersRepository interface ===

func (r *tblUsersRepository) FindByEIDAndPass(eid string, pass string) (*tblusers.TblUsers, error) {
	query := `
		SELECT 
			EID, NAM, PWORD1, DEPT, SID, GROUPS, SLEVEL, USER_STATUS, SID2 
		FROM tbl_users 
		WHERE 1=1
		AND EID = @p1 
		AND PWORD1 = @p2
	`
	row := r.db.QueryRow(query, &eid, &pass)

	var user tblusers.TblUsers
	err := row.Scan(&user.EID, &user.NAM, &user.PWORD1, &user.DEPT, &user.SID, &user.GROUPS, &user.SLEVEL, &user.USER_STATUS, &user.SID2)
	if err != nil {
		return nil, err
	}
	return &user, nil
}

func (r *tblUsersRepository) UpdateSidByEid(req tblusers.UpdateSidRequestResponse) error {
	query := `
		UPDATE tbl_users 
		SET SID = @p1 
		WHERE EID = @p2
	`
	_, err := r.db.Exec(query, &req.Sid, &req.Eid)
	return err
}

func (r *tblUsersRepository) FindByEid(eic string) (*tblusers.TblUsers, error) {
	query := `
		SELECT 
			EID, NAM, PWORD1, DEPT, SID, GROUPS, SLEVEL, USER_STATUS, SID2 
		FROM tbl_users 
		WHERE EID = @p1
	`
	row := r.db.QueryRow(query, &eic)

	var user tblusers.TblUsers
	err := row.Scan(&user.EID, &user.NAM, &user.PWORD1, &user.DEPT, &user.SID, &user.GROUPS, &user.SLEVEL, &user.USER_STATUS, &user.SID2)
	if err != nil {
		return nil, err
	}
	return &user, nil
}

func (r *tblUsersRepository) FindSidByEid(ctx context.Context, eid string) (string, error) {
	query := `
		SELECT SID 
		FROM tbl_users 
		WHERE EID = @p1
	`
	row := r.db.QueryRow(query, &eid)

	var sid string
	err := row.Scan(&sid)
	if err != nil {
		return "", err
	}
	return sid, nil
}

func (r *tblUsersRepository) Create(ctx context.Context, user *tblusers.TblUsers) error {
	query := `
		INSERT INTO tbl_users (EID, NAM, PWORD1, DEPT, SID, GROUPS, SLEVEL, USER_STATUS, SID2)
		VALUES (@p1, @p2, @p3, @p4, @p5, @p6, @p7, @p8, @p9)
	`
	_, err := r.db.Exec(query,
		user.EID,
		user.NAM,
		user.PWORD1,
		user.DEPT,
		user.SID,
		user.GROUPS,
		user.SLEVEL,
		user.USER_STATUS,
		user.SID2,
	)
	return err
}

func (r *tblUsersRepository) GetSID(eid, sid string) (string, error) {
	query := `
		SELECT 
			SID
		FROM tbl_users
		WHERE EID = @p1
		AND (SID = @p2 OR charindex(@p2, SID) > 0)
	`
	row := r.db.QueryRow(query, &eid, &sid)

	var result string
	err := row.Scan(&sid)
	if err != nil {
		return "", err
	}
	return result, nil
}
