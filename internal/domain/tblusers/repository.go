package tblusers

import (
	"context"
)

type TblUsersRepository interface {
	Create(ctx context.Context, user *TblUsers) error
	UpdateSidByEid(req UpdateSidRequestResponse) error
	FindByEid(eic string) (*TblUsers, error)
	FindSidByEid(ctx context.Context, eid string) (string, error)
	FindByEIDAndPass(eid string, pass string) (*TblUsers, error)
	GetSID(eid, sid string) (string, error)
}
