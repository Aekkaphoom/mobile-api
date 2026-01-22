package tblcustomers_usecase

import (
	"strings"

	"grouplease.co.th/mobile_api/internal/domain"
	"grouplease.co.th/mobile_api/internal/domain/tblcustomers"
)

type tblCustomersUsecase struct {
	repo tblcustomers.TblCustomersRepository
}

func NewCustomersUsecase(repo tblcustomers.TblCustomersRepository) tblcustomers.TblCustomersUsecase {
	return &tblCustomersUsecase{
		repo: repo,
	}
}

func (u *tblCustomersUsecase) AddFollowupCustomer(customer tblcustomers.CustomerRequest) error {
	if strings.TrimSpace(customer.ProfileId) == "" || strings.TrimSpace(customer.ContactId) == "" {
		return domain.ErrInvalidFollowUpRequest
	}

	return u.repo.Add(customer)
}

func (u *tblCustomersUsecase) UpdateFlgstsFollowup(profileId, contactId, statusUpdate string) error {
	if profileId == "" || contactId == "" || statusUpdate == "" {
		return domain.ErrInvalidInput
	}

	return u.repo.UpdateFlgsts(profileId, contactId, statusUpdate)
}
