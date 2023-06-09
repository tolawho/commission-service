package routes

import (
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"medici.vn/commission-serivce/config"
	"medici.vn/commission-serivce/controller"
	"medici.vn/commission-serivce/middleware"
	"medici.vn/commission-serivce/repository"
	"medici.vn/commission-serivce/services"
	"net/http"
)

var (
	db                              *gorm.DB                                   = config.SetupDatabaseConnection()
	userRepository                  repository.UserRepository                  = repository.NewUserRepository(db)
	pntDailyCommissionRepository    repository.PntDailyCommissionRepository    = repository.NewPntDailyCommissionRepository(db)
	pntContractRepository           repository.PntContractRepository           = repository.NewPntContractRepository(db)
	pntCommissionFormulaRepository  repository.PntCommissionFormulaRepository  = repository.NewPntCommissionFormulaRepository(db)
	pntPolicyRepository             repository.PntPolicyRepository             = repository.NewPntPolicyRepository(db)
	pntAgencyTreeRepository         repository.PntAgencyTreeRepository         = repository.NewPntAgencyTreeRepository(db)
	agencyRepository                repository.AgencyRepository                = repository.NewAgencyRepository(db)
	pntTransactionRepository        repository.PntTransactionRepository        = repository.NewPntTransactionRepository(db)
	pntTransactionHistoryRepository repository.PntTransactionHistoryRepository = repository.NewPntTransactionHistoryRepository(db)
	jwtService                      services.JWTService                        = services.NewJWTService()
	authService                     services.AuthService                       = services.NewAuthService(userRepository)
	pntDailyCommissionService       services.PntDailyCommissionService         = services.NewPntDailyCommissionService(
		pntDailyCommissionRepository,
		pntContractRepository,
		pntCommissionFormulaRepository,
		pntAgencyTreeRepository,
		pntPolicyRepository,
		agencyRepository,
		pntTransactionRepository,
		pntTransactionHistoryRepository,
	)
	authController    controller.AuthController    = controller.NewAuthController(authService, jwtService)
	nonLifeController controller.NonLifeController = controller.NewNonLifeController(pntDailyCommissionService)
)

func Router() *gin.Engine {
	return setupRouter()
}

func setupRouter() *gin.Engine {
	r := gin.Default()

	r.GET("/", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"data": "Welcome to Medici!"})
	})

	apiV1 := r.Group("/api/v1")
	{
		apiV1.GET("/test", middleware.AuthorizeJWT(jwtService))
		apiV1.POST("auth/token", authController.Login)

		apiV1NonLife := apiV1.Group("/non-life", middleware.AuthorizeJWT(jwtService))

		apiV1NonLife.POST("/commission/:contract_id", nonLifeController.Calculator)

		apiV1NonLife.POST("/commission/temporary/:contract_id", nonLifeController.Temporary)
	}

	return r
}
