package handler

import (
	"net/http"

	"github.com/ryota1119/time_resport_webapi/internal/interface/presenter"

	"github.com/gin-gonic/gin"
	"github.com/ryota1119/time_resport_webapi/internal/usecase"
)

// OrganizationHandler はorganizationHandlerのインターフェース
type OrganizationHandler interface {
	Register(ctx *gin.Context)
	GetOrganization(ctx *gin.Context)
}

// organizationHandler の実装
type organizationHandler struct {
	organizationRegisterUsecase  usecase.OrganizationRegisterUsecase
	organizationGetByCodeUsecase usecase.OrganizationGetByCodeUsecase
}

var _ OrganizationHandler = (*organizationHandler)(nil)

// NewOrganizationHandler はorganizationHandlerの初期化を行う
func NewOrganizationHandler(
	organizationRegisterUsecase usecase.OrganizationRegisterUsecase,
	organizationGetByCodeUsecase usecase.OrganizationGetByCodeUsecase,
) OrganizationHandler {
	return &organizationHandler{
		organizationRegisterUsecase:  organizationRegisterUsecase,
		organizationGetByCodeUsecase: organizationGetByCodeUsecase,
	}
}

type RegisterOrganizationBodyRequest struct {
	OrganizationName string `json:"organization_name" binding:"required" example:"My Organization"`
	OrganizationCode string `json:"organization_code" binding:"required" example:"my_organization_code"`
	UserName         string `json:"user_name" binding:"required" example:"山田太郎"`
	UserEmail        string `json:"user_email" binding:"required" example:"example@example.com"`
	Password         string `json:"password" binding:"required" example:"password"`
}

// Register は組織と管理者ユーザーを新規作成する
//
//	@Summary		Register
//	@Description	組織と管理者ユーザーを新規作成する
//	@Tags			organization
//	@Param			payload	body	RegisterOrganizationBodyRequest	true	"組織と管理者ユーザーを新規作成するボディリクエスト"
//	@Accept			json
//	@Produce		json
//	@Success		201	{object}	presenter.OrganizationRegisterResponse
//	@Failure		400	{object}	nil	"BadRequest"
//	@Router			/organization/register [POST]
func (h *organizationHandler) Register(c *gin.Context) {
	ctx := c.Request.Context()

	var req RegisterOrganizationBodyRequest
	err := c.ShouldBindJSON(&req)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	input := usecase.OrganizationUsecaseRegisterInput{
		OrganizationName: req.OrganizationName,
		OrganizationCode: req.OrganizationCode,
		UserName:         req.UserName,
		UserEmail:        req.UserEmail,
		Password:         req.Password,
	}
	org, err := h.organizationRegisterUsecase.Register(ctx, input)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	res := presenter.NewOrganizationRegisterResponse(org)
	c.JSON(http.StatusCreated, res)
}

// GetOrganization は組織情報を取得する
//
//	@Summary		GetOrganization
//	@Description	組織情報を取得する
//	@Tags			organization
//	@Accept			json
//	@Produce		json
//	@Success		200	{object}	presenter.OrganizationRegisterResponse
//	@Failure		404	{object}	nil	"NotFound"
//	@Failure		400	{object}	nil	"BadRequest"
//	@Router			/organization [GET]
func (h *organizationHandler) GetOrganization(c *gin.Context) {
	//ctx := c.Request.Context()
	//
	//org, err := h.organizationGetUsecase.GetByCode(ctx)
	//if err != nil {
	//	if errors.Is(err, domainerrors.ErrOrganizationNotFound) {
	//		c.JSON(http.StatusNotFound, gin.H{"error": ErrNotFound.Error()})
	//		return
	//	}
	//	c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	//	return
	//}
	//
	//res := presenter.NewOrganizationRegisterResponse(org)
	//c.JSON(http.StatusOK, res)
	c.JSON(http.StatusOK, nil)
}
