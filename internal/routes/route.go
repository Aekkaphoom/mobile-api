package routes

import (
	"database/sql"

	"github.com/gofiber/fiber/v2"
	mobilehttp "grouplease.co.th/mobile_api/internal/handler/http"
	dbRepo "grouplease.co.th/mobile_api/internal/repository/sqlserver"
	"grouplease.co.th/mobile_api/internal/usecase/authtoken"
	tblusers_usecase "grouplease.co.th/mobile_api/internal/usecase/tblusers"
	tbluserslogin_usecase "grouplease.co.th/mobile_api/internal/usecase/tbluserslogin"
	middlewareAuthtoken "grouplease.co.th/mobile_api/middleware/authtoken"
)

func FiberRouteUser(f fiber.Router, db *sql.DB) {
	repo := dbRepo.NewTblUsersRepository(db)
	usecase := tblusers_usecase.NewTblUsersUsecase(repo)

	h := mobilehttp.NewTblUsersHandler(usecase)

	f.Post("user/add", h.AddUserApi)
	f.Post("user/update-sid", h.UpdateSidByEidApi)
	// f.Get("/user/:eid", h.LoginUserApi)
	f.Get("user/sid/:eid", h.GetSidByEidApi)

}

func FiberRouteUserLogin(f fiber.Router, db *sql.DB) {
	userRepo := dbRepo.NewTblUsersRepository(db)
	loginRepo := dbRepo.NewTblUsersLoginRepository(db)

	userUc := tblusers_usecase.NewTblUsersUsecase(userRepo)
	loginUc := tbluserslogin_usecase.NewUsersLoginUsecase(loginRepo)

	h := mobilehttp.NewUsersLoginHandler(userUc, loginUc)

	f.Post("user/login/:role", h.LoginApi)
	f.Post("user/logout/:eid", h.LogoutUserApi)

}

func FiberRouteAuthToken(f fiber.Router, db *sql.DB) {
	authRepo := dbRepo.NewAuthtokenRepository(db)
	userRepo := dbRepo.NewTblUsersRepository(db)
	loginRepo := dbRepo.NewTblUsersLoginRepository(db)

	authUc := authtoken.NewAuthtokenUsecase(authRepo, userRepo, loginRepo)

	userUc := tblusers_usecase.NewTblUsersUsecase(userRepo)
	loginUc := tbluserslogin_usecase.NewUsersLoginUsecase(loginRepo)

	h := mobilehttp.NewAuthtokenHandler(authUc)
	loginHandler := mobilehttp.NewUsersLoginHandler(userUc, loginUc)

	middleAuth := middlewareAuthtoken.MiddleAuthtokenHandler(authUc)

	f.Get("user/login/:eid/:authtoken", middleAuth.TokenLoginCheckExpire(), h.GetAuthExpire)
	f.Get("user/login/profile/:eid/:authtoken", loginHandler.GetUserLoginOnline)
	f.Post("user/logout/:eid/:authtoken", loginHandler.LogoutUserApi)

	f.Get("auth/check/:eid", h.GetAuthExpire)

}
