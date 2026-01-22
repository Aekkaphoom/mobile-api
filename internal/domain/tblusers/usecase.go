package tblusers

import "context"

type TblUsersUseCase interface {
	Create(ctx context.Context, user *TblUsers) error
	UpdateSidByEid(ctx context.Context, req UpdateSidRequestResponse) (*UpdateSidRequestResponse, error)
	FindByEid(eic string) (*TblUsers, error)
	FindSidByEid(ctx context.Context, eid string) (string, error)
	FindByEIDAndPass(eid string, pass string) (*TblUsers, error)
	CheckSidDuplicate(eid, sid string) (string, error)
}
