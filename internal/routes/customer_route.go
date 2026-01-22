package routes

import (
	"database/sql"

	"github.com/gofiber/fiber/v2"
	mobilehttp "grouplease.co.th/mobile_api/internal/handler/http"
	"grouplease.co.th/mobile_api/internal/repository/sqlserver"
	tblcustomers_usecase "grouplease.co.th/mobile_api/internal/usecase/tblcustomers"
	tblcustomersphoto_usecase "grouplease.co.th/mobile_api/internal/usecase/tblcustomersphoto"
)

func FiberRouteCustomers(f fiber.Router, db *sql.DB) {
	repo := sqlserver.NewTblCustomersRepository(db)
	uc := tblcustomers_usecase.NewCustomersUsecase(repo)

	h := mobilehttp.NewCustomersHandler(uc)

	f.Post("customer/followup", h.FollowUpCustomer)
}

func FiberRouteCustomersPhoto(f fiber.Router, db *sql.DB, path string) {
	repo := sqlserver.NewTblCustomersPhotoRepo(db)
	uc := tblcustomersphoto_usecase.NewTblcustomersphotoUsecase(repo, path)

	h := mobilehttp.NewTblcustomersphotoHandler(uc)

	f.Post("customer/photo/save/path", h.SavePhotoToLocalApi)
	f.Post("customer/photo/save/db", h.SaveDataPhotoApi)
}
