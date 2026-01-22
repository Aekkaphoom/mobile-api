package tblcustomersphoto

type TblCustomersPhoto struct {
	ProfileId string
	Eid       string
	ContactId string
	PhotoName string
	GpsLocal  string
	InstDate  string
}

type CustomersPhotoRequest struct {
	ProfileId string `json:"profile_id"`
	Eid       string `json:"eid"`
	ContactId string `json:"contact_id"`
	PhotoName string `json:"photo_name"`
	GpsLocal  string `json:"gps_local"`
	InstDate  string `json:"inst_date"`
}

type PhotoRequest struct {
	// ProfileId string `json:"profile_id"`
	PhotoName string `json:"photo_name"`
	Base64    string `json:"base64"`
}
