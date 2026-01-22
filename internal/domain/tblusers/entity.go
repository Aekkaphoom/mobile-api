package tblusers

type TblUsers struct {
	EID         string
	NAM         string
	PWORD1      string
	DEPT        string
	SID         string
	GROUPS      int
	SLEVEL      string
	USER_STATUS string
	SID2        string
}

type UpdateSidFormRequest struct {
	EID string `json:"eid" form:"eid" validate:"required"`
	SID string `json:"sid" form:"sid" validate:"required"`
}

type UpdateSidRequestResponse struct {
	Eid string `json:"eid"`
	Sid string `json:"sid"`
}
