package tblusers_usecase

import (
	"context"
	"strings"

	"grouplease.co.th/mobile_api/internal/domain"
	"grouplease.co.th/mobile_api/internal/domain/tblusers"
)

type tblUsersUsecase struct {
	tblUsersRepo tblusers.TblUsersRepository
}

func NewTblUsersUsecase(tblUsersRepo tblusers.TblUsersRepository) tblusers.TblUsersUseCase {
	return &tblUsersUsecase{
		tblUsersRepo: tblUsersRepo,
	}
}

// === Implement TblUsersUsecase interface ===

func (u *tblUsersUsecase) Create(ctx context.Context, user *tblusers.TblUsers) error {
	if user == nil {
		return domain.ErrInvalidInput
	}
	return u.tblUsersRepo.Create(ctx, user)
}

func (u *tblUsersUsecase) UpdateSidByEid(ctx context.Context, req tblusers.UpdateSidRequestResponse) (*tblusers.UpdateSidRequestResponse, error) {
	if strings.TrimSpace(req.Eid) == "" || strings.TrimSpace(req.Sid) == "" {
		return nil, domain.ErrRequestMissing
	}

	sid, err := u.FindSidByEid(ctx, req.Eid)
	if err != nil {
		return nil, err
	}

	var newSid string
	chkDuplicate := false

	if sid == "" {
		newSid = req.Sid
	} else {
		vSid := strings.Split(sid, ",")
		for i := range vSid {
			if vSid[i] == req.Sid {
				chkDuplicate = true
				break
			}
		}
		if !chkDuplicate {
			newSid = sid + "," + req.Sid
		}
	}

	if !chkDuplicate {
		req.Sid = newSid
		err := u.tblUsersRepo.UpdateSidByEid(req)
		if err != nil {
			return nil, err
		}
	} else {
		return nil, domain.ErrSIDExists
	}

	resp, err := u.tblUsersRepo.FindByEid(req.Eid)
	if err != nil {
		return nil, err
	}
	user := &tblusers.UpdateSidRequestResponse{
		Eid: resp.EID,
		Sid: resp.SID,
	}

	return user, nil
}

func (u *tblUsersUsecase) FindByEid(eic string) (*tblusers.TblUsers, error) {
	return u.tblUsersRepo.FindByEid(eic)
}

func (u *tblUsersUsecase) FindSidByEid(ctx context.Context, eid string) (string, error) {
	return u.tblUsersRepo.FindSidByEid(ctx, eid)
}

func (u *tblUsersUsecase) FindByEIDAndPass(eid string, pass string) (*tblusers.TblUsers, error) {
	if eid == "" || pass == "" {
		return nil, domain.ErrInvalidInput
	}
	return u.tblUsersRepo.FindByEIDAndPass(eid, pass)
}

func (u *tblUsersUsecase) CheckSidDuplicate(eid, sid string) (string, error) {
	result, err := u.tblUsersRepo.GetSID(eid, sid)
	if err != nil {
		return "false", err
	}

	return result, nil
}
