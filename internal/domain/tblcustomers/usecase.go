package tblcustomers

type TblCustomersUsecase interface {
	AddFollowupCustomer(customer CustomerRequest) error
	UpdateFlgstsFollowup(profileId, contactId, statusUpdate string) error
}
