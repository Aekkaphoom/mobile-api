package sqlserver

import (
	"database/sql"

	"grouplease.co.th/mobile_api/internal/domain/tblcustomers"
)

type tblCustomersRepository struct {
	db *sql.DB
}

func NewTblCustomersRepository(db *sql.DB) tblcustomers.TblCustomersRepository {
	return &tblCustomersRepository{
		db: db,
	}
}

// === Implement Repo interface ===

func (r *tblCustomersRepository) Add(customer tblcustomers.CustomerRequest) error {
	query := `
	INSERT INTO GLMOBILE_DB.dbo.tbl_customers
		(PROFILE_ID, EID, CUSTOMER_NAME, CONTACT_ID, TEXT_NOTE, GPS_LOCAL, RESULT_CODE, INST_DATE, flgsts, box_no)
	VALUES
		(@p1, @p2, @p3, @p4, @p5, @p6, @p7, @p8, @p9, @p10)
		;
	`
	_, err := r.db.Exec(query,
		customer.ProfileId,
		customer.Eid,
		customer.CustomerName,
		customer.ContactId,
		customer.TextNote,
		customer.GpsLocal,
		customer.ResultCode,
		customer.InstDate,
		customer.Flgsts,
		customer.BoxNo,
	)

	return err
}

func (r *tblCustomersRepository) UpdateFlgsts(profileId, contactId, statusUpdate string) error {
	query := `
		UPDATE GLMOBILE_DB.dbo.tbl_customers
			SET flgsts=@p1
		WHERE PROFILE_ID=@p2
		AND CONTACT_ID=@p3
		;
	`
	_, errUpt := r.db.Exec(query,
		&statusUpdate,
		&profileId,
		&contactId,
	)
	return errUpt
}
