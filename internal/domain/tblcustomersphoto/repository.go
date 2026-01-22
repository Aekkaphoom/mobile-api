package tblcustomersphoto

type UsersRepository interface {
	Add(customerPhoto *CustomersPhotoRequest) error
	SavePhotoToLocal(photoreq *PhotoRequest, path string) error
}
