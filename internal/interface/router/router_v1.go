package router

import (
	"github.com/gin-gonic/gin"
	_ "github.com/ryota1119/time_resport/cmd/api/docs/swagger"
	"github.com/ryota1119/time_resport/internal/interface/handler"
	"github.com/ryota1119/time_resport/internal/interface/middleware"
	swaggerfiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

// Router はRouterを実装
type Router struct {
	authMiddleware      middleware.AuthMiddleware
	authHandler         handler.AuthHandler
	organizationHandler handler.OrganizationHandler
	userHandler         handler.UserHandler
	customerHandler     handler.CustomerHandler
	projectHandler      handler.ProjectHandler
	budgetHandler       handler.BudgetHandler
	timerHandler        handler.TimerHandler
}

// NewRouter はRouterを初期化
func NewRouter(
	authMiddleware middleware.AuthMiddleware,
	authHandler handler.AuthHandler,
	organizationHandler handler.OrganizationHandler,
	userHandler handler.UserHandler,
	customerHandler handler.CustomerHandler,
	projectHandler handler.ProjectHandler,
	budgetHandler handler.BudgetHandler,
	timerHandler handler.TimerHandler,
) *Router {
	return &Router{
		authMiddleware:      authMiddleware,
		authHandler:         authHandler,
		organizationHandler: organizationHandler,
		userHandler:         userHandler,
		customerHandler:     customerHandler,
		projectHandler:      projectHandler,
		budgetHandler:       budgetHandler,
		timerHandler:        timerHandler,
	}
}

// SetupRouter はRouterをセットアップする
func (r *Router) SetupRouter(engin *gin.Engine) {
	// **Swagger**
	engin.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerfiles.Handler))

	rootGroup := engin.Group("/api/v1/")
	r.setupOrganizationRoutes(rootGroup)
	r.setupAuthRoutes(rootGroup)
	r.setupUserRoutes(rootGroup)
	r.setupCustomerRoutes(rootGroup)
	r.setupProjectRoutes(rootGroup)
	r.setupBudgetRoutes(rootGroup)
	r.setupTimerRoutes(rootGroup)
}

// setupAuthRoutes 認証関連のルート設定
func (r *Router) setupAuthRoutes(group *gin.RouterGroup) {
	authGroup := group.Group("/auth")
	{
		authGroup.POST("/login", r.authHandler.Login)
		authGroup.POST("/refresh", r.authHandler.RefreshToken)
		authGroup.DELETE("/logout", r.authMiddleware.AuthMiddleware(), r.authHandler.Logout)
	}
}

// setupOrganizationRoutes 組織関連のルート設定
func (r *Router) setupOrganizationRoutes(group *gin.RouterGroup) {
	organizationGroup := group.Group("/organizations")
	{
		organizationGroup.POST("/register", r.organizationHandler.Register)
	}
}

// setupUserRoutes ユーザー関連のルート設定
func (r *Router) setupUserRoutes(group *gin.RouterGroup) {
	userGroup := group.Group("/users").Use(r.authMiddleware.AuthMiddleware())
	{
		userGroup.GET("", r.userHandler.List)
		userGroup.GET("/:userID", r.userHandler.Get)
		userGroup.PUT("/:userID", r.userHandler.Update)
	}

	// 管理者権限のみ実行可能
	adminGroup := userGroup.Use(r.authMiddleware.RequireAdmin())
	{
		adminGroup.POST("", r.userHandler.Create)
		adminGroup.DELETE("/:userID", r.userHandler.Delete)
	}
}

// setupCustomerRoutes 顧客関連のルート設定
func (r *Router) setupCustomerRoutes(group *gin.RouterGroup) {
	customerGroup := group.Group("/customers").Use(r.authMiddleware.AuthMiddleware())
	{
		customerGroup.GET("", r.customerHandler.List)
		customerGroup.GET("/:customerID", r.customerHandler.Get)
	}

	// 管理者権限のみ実行可能
	adminGroup := customerGroup.Use(r.authMiddleware.RequireAdmin())
	{
		adminGroup.POST("", r.customerHandler.Create)
		adminGroup.PUT("/:customerID", r.customerHandler.Update)
		adminGroup.DELETE("/:customerID", r.customerHandler.Delete)
	}
}

// setupProjectRoutes プロジェクト関連のルート設定
func (r *Router) setupProjectRoutes(group *gin.RouterGroup) {
	projectGroup := group.Group("/projects").Use(r.authMiddleware.AuthMiddleware())
	{
		projectGroup.GET("", r.projectHandler.List)
		projectGroup.GET("/:projectID", r.projectHandler.Get)
	}

	// 管理者権限のみ実行可能
	adminGroup := projectGroup.Use(r.authMiddleware.RequireAdmin())
	{
		adminGroup.POST("", r.projectHandler.Create)
		adminGroup.PUT("/:customerID", r.projectHandler.Update)
		adminGroup.DELETE("/:customerID", r.projectHandler.Delete)
	}
}

// setupBudgetRoutes 予算関連のルート設定
func (r *Router) setupBudgetRoutes(group *gin.RouterGroup) {
	budgetGroup := group.Group("/budgets").Use(r.authMiddleware.AuthMiddleware())
	{
		budgetGroup.GET("", r.budgetHandler.List)
		budgetGroup.GET("/:budgetID", r.budgetHandler.Get)
	}

	// 管理者権限のみ実行可能
	adminGroup := budgetGroup.Use(r.authMiddleware.RequireAdmin())
	{
		adminGroup.POST("", r.budgetHandler.Create)
		adminGroup.PUT("/:budgetID", r.budgetHandler.Update)
		adminGroup.DELETE("/:budgetID", r.budgetHandler.Delete)
	}
}

// setupTimerRoutes タイマー関連のルート設定
func (r *Router) setupTimerRoutes(group *gin.RouterGroup) {
	timerGroup := group.Group("/timers").Use(r.authMiddleware.AuthMiddleware())
	{
		timerGroup.POST("start", r.timerHandler.Start)
	}
}
