package handler

import (
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/ryota1119/time_resport/internal/interface/presenter"
	"github.com/ryota1119/time_resport/internal/usecase"
)

// AuthHandler はauthHandlerのインターフェース
type AuthHandler interface {
	Login(c *gin.Context)
	RefreshToken(c *gin.Context)
	Logout(c *gin.Context)
}

// timerHandler の実装
type authHandler struct {
	authLoginUsecase        usecase.AuthLoginUsecase
	authRefreshTokenUsecase usecase.AuthRefreshTokenUsecase
	authLogoutUsecase       usecase.AuthLogoutUsecase
}

var _ AuthHandler = (*authHandler)(nil)

// NewAuthHandler はauthHandlerの初期化を行う
func NewAuthHandler(
	authLoginUsecase usecase.AuthLoginUsecase,
	authRefreshTokenUsecase usecase.AuthRefreshTokenUsecase,
	authLogoutUsecase usecase.AuthLogoutUsecase,
) AuthHandler {
	return &authHandler{
		authLoginUsecase:        authLoginUsecase,
		authRefreshTokenUsecase: authRefreshTokenUsecase,
		authLogoutUsecase:       authLogoutUsecase,
	}
}

// AuthLoginBodyRequest はログインのボディリクエスト構造体
type AuthLoginBodyRequest struct {
	OrganizationCode string `json:"organization_code" binding:"required" example:"my_organization_code"`
	Email            string `json:"email" binding:"required" example:"example@example.com"`
	Password         string `json:"password" binding:"required" example:"password"`
}

// Login はログインを実行する
//
//	@Summary		Login
//	@Description	ユーザーログインを行う
//	@Tags			auth
//	@Param			payload	body	AuthLoginBodyRequest	true	"payload"
//	@Accept			json
//	@Produce		json
//	@Success		200	{object}	presenter.AuthLoginResponse
//	@Failure		400	{object}	nil	"BadRequest"
//	@Router			/auth/login [post]
func (h *authHandler) Login(c *gin.Context) {
	ctx := c.Request.Context()

	var req AuthLoginBodyRequest
	err := c.ShouldBindJSON(&req)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	input := usecase.AuthUsecaseLoginInput{
		OrganizationCode: req.OrganizationCode,
		Email:            req.Email,
		Password:         req.Password,
	}
	token, err := h.authLoginUsecase.Login(ctx, input)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.SetSameSite(http.SameSiteNoneMode)
	c.SetCookie("access_token", token.AccessToken.String(), 3600, "/", os.Getenv("COOKIE_DOMAIN"), true, true)
	c.SetCookie("refresh_token", token.RefreshToken.String(), 30*24*3600, "/", os.Getenv("COOKIE_DOMAIN"), true, true)

	resp := presenter.NewAuthLoginResponse(token)
	c.JSON(http.StatusOK, resp)
}

type RefreshTokenAuthRequest struct {
	RefreshToken string `json:"refresh_token" binding:"required" example:"refresh_token"`
}

// RefreshToken は RefreshToken からAccessTokenを再生成する
//
//	@Summary		RefreshToken
//	@Description	リフレッシュトークンを利用してアクセストークンを再生成を行う
//	@Tags			auth
//	@Param			payload	body	RefreshTokenAuthRequest	true	"payload"
//	@Success		200
//	@Accept			json
//	@Produce		json
//	@Success		201	{object}	presenter.AuthRefreshTokenResponse
//	@Failure		400	{object}	nil	"BadRequest"
//	@Router			/auth/refresh [post]
func (h *authHandler) RefreshToken(c *gin.Context) {
	ctx := c.Request.Context()

	var req RefreshTokenAuthRequest
	err := c.ShouldBindJSON(&req)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	input := usecase.AuthUsecaseRefreshTokenInput{
		RefreshToken: req.RefreshToken,
	}
	token, err := h.authRefreshTokenUsecase.RefreshToken(ctx, input)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	resp := presenter.NewAuthRefreshTokenResponse(token)
	c.JSON(http.StatusOK, resp)
}

// Logout はログアウトを行う
//
//	@Summary		Logout
//	@Description	ログアウトを行う
//	@Tags			auth
//	@Security		BearerAuth
//	@Success		200
//	@Accept			json
//	@Produce		json
//	@Success		200	{object}	string	"logged out"
//	@Failure		400	{object}	nil		"BadRequest"
//	@Router			/auth/logout [DELETE]
func (h *authHandler) Logout(c *gin.Context) {
	ctx := c.Request.Context()
	if err := h.authLogoutUsecase.Logout(ctx); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "logged out"})
}
