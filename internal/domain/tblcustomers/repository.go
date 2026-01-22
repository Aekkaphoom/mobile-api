package tblcustomers

type TblCustomersRepository interface {
	Add(customer CustomerRequest) error
	UpdateFlgsts(profileId, contactId, statusUpdate string) error
}
