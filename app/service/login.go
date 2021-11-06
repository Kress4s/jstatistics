package service

// var (
// 	loginServiceInstance UserService
// 	loginOnce            sync.Once
// )

// type loginServiceImpl struct {
// 	db   *gorm.DB
// 	repo repositories.UserRepo
// }

// type LoginService interface {
// 	Create(params *vo.UserReq) exception.Exception
// }

// func GetLoginService() UserService {
// 	loginOnce.Do(func() {
// 		loginServiceInstance = &userServiceImpl{
// 			db:   database.GetDriver(),
// 			repo: repositories.GetUserRepo(),
// 		}
// 	})
// 	return loginServiceInstance
// }
