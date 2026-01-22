package tblcustomersphoto

type UsersUsecase interface {
	AddCustomersPhoto(customersPhoto *CustomersPhotoRequest) error
	SavePhoto(photo *PhotoRequest) error
}
