package main

import (
	"github.com/ryota1119/time_resport/internal/infrastructure/logger"
	"log"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"github.com/ryota1119/time_resport/internal/infrastructure/auth"
	"github.com/ryota1119/time_resport/internal/infrastructure/database"
	"github.com/ryota1119/time_resport/internal/infrastructure/jwt_token"
	"github.com/ryota1119/time_resport/internal/infrastructure/redis"
	"github.com/ryota1119/time_resport/internal/infrastructure/repository"
	"github.com/ryota1119/time_resport/internal/interface/handler"
	"github.com/ryota1119/time_resport/internal/interface/middleware"
	"github.com/ryota1119/time_resport/internal/interface/router"
	"github.com/ryota1119/time_resport/internal/usecase"
)

// @title						Time Report WebAPI
// @version					1.0
// @description				Time Report WebAPIのSwaggerドキュメント
// @termsOfService				http://example.com/terms/
// @host						localhost:8080
// @BasePath					/api/v1
// @securityDefinitions.apikey	BearerAuth
// @in							header
// @name						Authorization
func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	r := gin.Default()
	// Default middleware適用
	for _, mw := range middleware.Default() {
		r.Use(mw)
	}

	// Logger
	logger.Init(os.Getenv("APP_ENV"))

	// データベース初期化
	if err = database.NewDB(); err != nil {
		logger.Logger.Error(err.Error())
		os.Exit(1)
	}
	// Redis初期化
	if err = redis.NewRedis(); err != nil {
		logger.Logger.Error(err.Error())
		os.Exit(1)
	}

	// jwtAuth Serviceの初期化
	secretKey := []byte(os.Getenv("JWT_SECRET_KEY"))
	if len(secretKey) == 0 {
		logger.Logger.Error("JWT_SECRET_KEY is not set in environment variables")
		os.Exit(1)
	}
	jwtToken := jwt_token.NewJwtToken(secretKey)

	// 依存性の注入
	// MySQL
	db := database.GetDB()
	// Redis
	redisClient := redis.GetRedisClient()

	// Repositoryのセットアップ
	authRepo := repository.NewAuthRepository(redisClient)
	organizationRepo := repository.NewOrganizationRepository()
	userRepo := repository.NewUserRepository()
	customerRepo := repository.NewCustomerRepository()
	projectRepo := repository.NewProjectRepository()
	budgetRepo := repository.NewBudgetRepository()
	timerRepo := repository.NewTimerRepository()

	// Usecase層のセットアップ
	authLoginUsecase := usecase.NewAuthLoginUsecase(db, jwtToken, authRepo, organizationRepo, userRepo)
	authRefreshTokenUsecase := usecase.NewAuthRefreshTokenUsecase(db, jwtToken, authRepo, organizationRepo, userRepo)
	authLogoutUsecase := usecase.NewAuthLogoutUsecase(authRepo)
	organizationRegisterUsecase := usecase.NewOrganizationRegisterUsecase(db, organizationRepo, userRepo)
	organizationGetByCodeUsecase := usecase.NewOrganizationGetByCodeUsecase(db, organizationRepo)
	userUsecase := usecase.NewUserUsecase(db, userRepo)
	customerUsecase := usecase.NewCustomerUsecase(db, customerRepo)
	projectUsecase := usecase.NewProjectUsecase(db, projectRepo)
	budgetCreateUsecase := usecase.NewBudgetCreateUsecase(db, budgetRepo)
	budgetListUsecase := usecase.NewBudgetListUsecase(db, budgetRepo)
	budgetGetUsecase := usecase.NewBudgetGetUsecase(db, budgetRepo)
	budgetUpdateUsecase := usecase.NewBudgetUpdateUsecase(db, budgetRepo)
	budgetDeleteUsecase := usecase.NewBudgetDeleteUsecase(db, budgetRepo)
	timerStartUsecase := usecase.NewTimerStartUsecase(db, timerRepo)
	timerStopUsecase := usecase.NewTimerStopUsecase(db, timerRepo)

	// Handler層のセットアップ
	authHandler := handler.NewAuthHandler(authLoginUsecase, authRefreshTokenUsecase, authLogoutUsecase)
	organizationHandler := handler.NewOrganizationHandler(organizationRegisterUsecase, organizationGetByCodeUsecase)
	userHandler := handler.NewUserHandler(userUsecase)
	customerHandler := handler.NewCustomerHandler(customerUsecase)
	projectHandler := handler.NewProjectHandler(projectUsecase)
	budgetHandler := handler.NewBudgetHandler(budgetCreateUsecase, budgetListUsecase, budgetGetUsecase, budgetUpdateUsecase, budgetDeleteUsecase)
	timerHandler := handler.NewTimerHandler(timerStartUsecase, timerStopUsecase)

	// infrastructure/auth Service層のセットアップ
	authService := auth.NewAuthService(db, redisClient, jwtToken, authRepo, organizationRepo, userRepo)
	// Middleware層のセットアップ
	// 認証ミドルウェア
	authMiddleware := middleware.NewAuthMiddleware(authService)

	// Routerのセットアップ
	// bootstrap初期化
	newRouter := router.NewRouter(
		authMiddleware,
		authHandler,
		organizationHandler,
		userHandler,
		customerHandler,
		projectHandler,
		budgetHandler,
		timerHandler,
	)
	newRouter.SetupRouter(r)

	// ポート番号8080番で起動
	err = r.Run(":8080")
	if err != nil {
		logger.Logger.Error(err.Error())
		os.Exit(1)
	}
}
