package sqlserver

import (
	"database/sql"
	"encoding/base64"
	"fmt"
	"os"
	"path/filepath"
	"strings"

	"grouplease.co.th/mobile_api/internal/domain/tblcustomersphoto"
)

type tblcustomersphotoRepository struct {
	db *sql.DB
}

func NewTblCustomersPhotoRepo(db *sql.DB) tblcustomersphoto.UsersRepository {
	return &tblcustomersphotoRepository{
		db: db,
	}
}

func saveBase64Image(b64Data, fileName, path string) error {
	// 1. Handle "Data URI" Scheme
	// If the string starts with a header like "data:image/png;base64,", strip it.
	// We look for the comma separator.
	if idx := strings.Index(b64Data, ","); idx != -1 {
		b64Data = b64Data[idx+1:]
	}

	// 2. Decode the Base64 string into bytes
	// Use StdEncoding unless you are using URL-safe base64 (then use URLEncoding)
	dec, err := base64.StdEncoding.DecodeString(b64Data)
	if err != nil {
		return fmt.Errorf("failed to decode base64: %w", err)
	}

	// // Create the directories
	// dir := filepath.Dir(path)
	// if err := os.MkdirAll(dir, 0755); err != nil {
	// 	return fmt.Errorf("failed to create directories: %w", err)
	// }

	// == Join path with filenam
	fullPath := filepath.Join(path, fileName)

	// 3. Write the bytes to a file
	// 0644 is a standard permission (Read/Write for owner, Read for others)
	err = os.WriteFile(fullPath, dec, 0644)
	if err != nil {
		return fmt.Errorf("failed to write file: %w", err)
	}

	return nil
}

// === Implement

func (r *tblcustomersphotoRepository) Add(customerPhoto *tblcustomersphoto.CustomersPhotoRequest) error {
	query := `
		INSERT INTO tbl_customers_photo
			(PROFILE_ID, EID, CONTACT_ID, PHOTO_NAME, GPS_LOCAL, INST_DATE)
		VALUES
			(@p1, @p2, @p3, @p4, @p5, @p6);
	`
	_, err := r.db.Exec(
		query,
		&customerPhoto.ProfileId,
		&customerPhoto.Eid,
		&customerPhoto.ContactId,
		&customerPhoto.PhotoName,
		&customerPhoto.GpsLocal,
		&customerPhoto.InstDate,
	)
	return err
}

func (r *tblcustomersphotoRepository) SavePhotoToLocal(photoreq *tblcustomersphoto.PhotoRequest, path string) error {

	err := saveBase64Image(photoreq.Base64, photoreq.PhotoName, path)
	if err != nil {
		return err
	}

	return nil
}
