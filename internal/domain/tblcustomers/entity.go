package tblcustomers

type TblCustomers struct {
	PROFILE_ID    string
	EID           string
	CUSTOMER_NAME string
	CONTACT_ID    string
	TEXT_NOTE     string
	FILE_NAME     string
	GPS_LOCAL     string
	RESULT_CODE   string
	INST_DATE     string
	Flgsts        string
	PIC_BASE      string
	EDIT_BY       string
	box_no        string
}

type CustomerRequest struct {
	ProfileId    string `json:"profile_id"`
	Eid          string `json:"eid"`
	CustomerName string `json:"customer_name"`
	ContactId    string `json:"contact_id"`
	TextNote     string `json:"text_note"`
	// FileName     string `json:"file_name"`
	GpsLocal   string `json:"gps_local"`
	ResultCode string `json:"result_code"`
	InstDate   string `json:"inst_date"`
	Flgsts     string `json:"flgsts"`
	// PicBase      string `json:"pic_base"`
	// EditBy       string `json:"edit_by"`
	BoxNo string `json:"box_no"`
}
