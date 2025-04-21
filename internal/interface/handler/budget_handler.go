package handler

import (
	"errors"
	"net/http"

	"github.com/ryota1119/time_resport_webapi/internal/domain/entities"

	"github.com/gin-gonic/gin"
	"github.com/ryota1119/time_resport_webapi/internal/interface/presenter"
	"github.com/ryota1119/time_resport_webapi/internal/usecase"
)

// BudgetHandler はbudgetHandlerのインターフェース
type BudgetHandler interface {
	Create(c *gin.Context)
	List(c *gin.Context)
	Get(c *gin.Context)
	Update(c *gin.Context)
	Delete(c *gin.Context)
}

// budgetHandler の実装
type budgetHandler struct {
	budgetCreateUsecase usecase.BudgetCreateUsecase
	budgetListUsecase   usecase.BudgetListUsecase
	budgetGetUsecase    usecase.BudgetGetUsecase
	budgetUpdateUsecase usecase.BudgetUpdateUsecase
	budgetDeleteUsecase usecase.BudgetDeleteUsecase
}

var _ BudgetHandler = (*budgetHandler)(nil)

// NewBudgetHandler はbudgetHandlerの初期化を行う
func NewBudgetHandler(
	budgetCreateUsecase usecase.BudgetCreateUsecase,
	budgetListUsecase usecase.BudgetListUsecase,
	budgetGetUsecase usecase.BudgetGetUsecase,
	budgetUpdateUsecase usecase.BudgetUpdateUsecase,
	budgetDeleteUsecase usecase.BudgetDeleteUsecase,
) BudgetHandler {
	return &budgetHandler{
		budgetCreateUsecase: budgetCreateUsecase,
		budgetListUsecase:   budgetListUsecase,
		budgetGetUsecase:    budgetGetUsecase,
		budgetUpdateUsecase: budgetUpdateUsecase,
		budgetDeleteUsecase: budgetDeleteUsecase,
	}
}

// CreateBudgetsRequest は予算作成時のリクエストデータ
type CreateBudgetsRequest struct {
	ProjectID    uint    `json:"projectID" binding:"required" example:"1"`
	BudgetAmount int64   `json:"budgetAmount" binding:"required" example:"300000"`
	BudgetMemo   *string `json:"budgetMemo" example:"budget memo"`
	StartDate    string  `json:"startDate" example:"2020-01"`
	EndDate      string  `json:"endDate" example:"2025-01"`
}

// Create は予算を新規作成する
//
//	@Summary		Create Budget
//	@Description	予算を新規作成する
//	@Tags			budget
//	@Security		BearerAuth
//	@Param			request	body	CreateBudgetsRequest	true	"budget create payload"
//	@Accept			json
//	@Produce		json
//	@Success		201	{object}	presenter.BudgetCreateResponse
//	@Failure		400	{object}	nil	"BadRequest"
//	@Failure		401	{object}	nil	"Unauthorized"
//	@Failure		403	{object}	nil	"Forbidden"
//	@Router			/budgets [POST]
func (h *budgetHandler) Create(c *gin.Context) {
	ctx := c.Request.Context()

	var req CreateBudgetsRequest
	err := c.ShouldBindJSON(&req)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	input := usecase.BudgetUsecaseCreateInput{
		ProjectID:    req.ProjectID,
		BudgetAmount: req.BudgetAmount,
		BudgetMemo:   req.BudgetMemo,
		StartDate:    req.StartDate,
		EndDate:      req.EndDate,
	}
	budget, err := h.budgetCreateUsecase.Create(ctx, input)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	res := presenter.NewBudgetCreateResponse(budget)
	c.JSON(http.StatusCreated, res)
}

// List は予算一覧を返却する
//
//	@Summary		List Budgets
//	@Description	予算一覧を返却する
//	@Tags			budget
//	@Security		BearerAuth
//	@Accept			json
//	@Produce		json
//	@Success		200	{array}		presenter.BudgetListResponse
//	@Failure		400	{object}	nil	"BadRequest"
//	@Failure		401	{object}	nil	"Unauthorized"
//	@Router			/budgets [GET]
func (h *budgetHandler) List(c *gin.Context) {
	ctx := c.Request.Context()

	budgets, err := h.budgetListUsecase.List(ctx)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	resp := presenter.NewBudgetListResponse(budgets)
	c.JSON(http.StatusOK, resp)
}

// BudgetGetURIRequest は予算情報取得のURIリクエスト構造体
type BudgetGetURIRequest struct {
	BudgetID uint `uri:"budgetID" binding:"required" example:"1"`
}

// Get は予算情報を返却する
//
//	@Summary		Get Budget
//	@Description	予算情報を返却する
//	@Tags			budget
//	@Security		BearerAuth
//	@Param			uri	path	BudgetGetURIRequest	true	"予算ID"
//	@Accept			json
//	@Produce		json
//	@Success		200	{object}	presenter.BudgetGetResponse
//	@Failure		400	{object}	nil	"BadRequest"
//	@Failure		401	{object}	nil	"Unauthorized"
//	@Router			/budgets/{budgetID} [GET]
func (h *budgetHandler) Get(c *gin.Context) {
	ctx := c.Request.Context()
	var req BudgetGetURIRequest
	if err := c.ShouldBindUri(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	input := usecase.BudgetUsecaseGetInput{
		BudgetID: req.BudgetID,
	}
	budget, err := h.budgetGetUsecase.Get(ctx, input)
	if err != nil {
		if errors.Is(err, entities.ErrBudgetNotFound) {
			c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
			return
		}
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	resp := presenter.NewBudgetGetResponse(budget)
	c.JSON(http.StatusOK, resp)
}

// BudgetUpdateURIRequest は予算更新のURIリクエスト構造体
type BudgetUpdateURIRequest struct {
	BudgetID uint `uri:"budgetID" binding:"required" example:"1"`
}

// BudgetUpdateBodyRequest は予算更新のボディリクエスト構造体
type BudgetUpdateBodyRequest struct {
	ProjectID    uint    `json:"projectID" binding:"required" example:"1"`
	BudgetAmount int64   `json:"amount" binding:"required" example:"300000"`
	BudgetMemo   *string `json:"memo" example:"budget memo"`
	StartMonth   string  `json:"startMonth" binding:"required" example:"2020-01-01"`
	EndMonth     string  `json:"endMonth" binding:"required" example:"2025-01-02"`
}

// Update は予算情報を更新する
//
//	@Summary		Update Budget
//	@Description	予算情報を更新する
//	@Tags			budget
//	@Security		BearerAuth
//	@Param			uri		path	BudgetUpdateURIRequest	true	"予算ID"
//	@Param			payload	body	BudgetUpdateBodyRequest	true	"update budgets payload"
//	@Accept			json
//	@Produce		json
//	@Success		201	{object}	presenter.BudgetUpdateResponse
//	@Success		204	{object}	nil	"NoContent"
//	@Failure		400	{object}	nil	"BadRequest"
//	@Failure		401	{object}	nil	"Unauthorized"
//	@Failure		403	{object}	nil	"Forbidden"
//	@Router			/budgets/{budgetID} [PUT]
func (h *budgetHandler) Update(c *gin.Context) {
	ctx := c.Request.Context()

	var uriReq BudgetUpdateURIRequest
	if err := c.ShouldBindUri(&uriReq); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	var bodyReq BudgetUpdateBodyRequest
	if err := c.ShouldBindJSON(&bodyReq); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// 更新処理
	input := usecase.BudgetUsecaseUpdateInput{
		BudgetID:     uriReq.BudgetID,
		ProjectID:    bodyReq.ProjectID,
		BudgetAmount: bodyReq.BudgetAmount,
		BudgetMemo:   bodyReq.BudgetMemo,
		StartDate:    bodyReq.StartMonth,
		EndDate:      bodyReq.EndMonth,
	}
	budget, err := h.budgetUpdateUsecase.Update(ctx, input)
	if err != nil {
		if errors.Is(err, entities.ErrNoContentUpdated) {
			c.JSON(http.StatusNoContent, nil)
		}
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	resp := presenter.NewBudgetUpdateResponse(budget)
	c.JSON(http.StatusCreated, resp)
}

// BudgetDeleteURIRequest は予算削除のURIリクエスト構造体
type BudgetDeleteURIRequest struct {
	BudgetID uint `uri:"budgetID" binding:"required" example:"1"`
}

// Delete は予算情報を削除する
//
//	@Summary		Delete Budget
//	@Description	予算情報を削除する
//	@Tags			budget
//	@Security		BearerAuth
//	@Param			uri	path	BudgetDeleteURIRequest	true	"budget ID"
//	@Accept			json
//	@Produce		json
//	@Success		204	{object}	nil
//	@Failure		400	{object}	nil	"BadRequest"
//	@Failure		401	{object}	nil	"Unauthorized"
//	@Failure		403	{object}	nil	"Forbidden"
//	@Router			/budgets/{budgetID} [DELETE]
func (h *budgetHandler) Delete(c *gin.Context) {
	ctx := c.Request.Context()

	var req BudgetDeleteURIRequest
	if err := c.ShouldBindUri(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	input := usecase.DeleteBudgetUsecaseInput{
		BudgetID: req.BudgetID,
	}
	err := h.budgetDeleteUsecase.Delete(ctx, input)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	}

	c.JSON(http.StatusNoContent, nil)
}
