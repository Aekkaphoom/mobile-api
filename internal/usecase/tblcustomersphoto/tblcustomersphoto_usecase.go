package tblcustomersphoto_usecase

import (
	"grouplease.co.th/mobile_api/internal/domain"
	"grouplease.co.th/mobile_api/internal/domain/tblcustomersphoto"
)

type tblcustomersphotoUsecase struct {
	repo      tblcustomersphoto.UsersRepository
	pathlocal string
}

func NewTblcustomersphotoUsecase(repo tblcustomersphoto.UsersRepository, path string) tblcustomersphoto.UsersUsecase {
	return &tblcustomersphotoUsecase{
		repo:      repo,
		pathlocal: path,
	}
}

// === Implement ===

func (r *tblcustomersphotoUsecase) AddCustomersPhoto(customersPhoto *tblcustomersphoto.CustomersPhotoRequest) error {
	if customersPhoto.ProfileId == "" || customersPhoto.ContactId == "" || customersPhoto.Eid == "" {
		return domain.ErrPhotoRequest
	}

	return r.repo.Add(customersPhoto)
}

func (r *tblcustomersphotoUsecase) SavePhoto(photo *tblcustomersphoto.PhotoRequest) error {
	// log.Printf("photo = %v", photo)
	return r.repo.SavePhotoToLocal(photo, r.pathlocal)
}
