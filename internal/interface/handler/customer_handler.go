package handler

import (
	"errors"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/ryota1119/time_resport_webapi/internal/interface/presenter"
	"github.com/ryota1119/time_resport_webapi/internal/usecase"
)

// CustomerHandler はcustomerHandlerのインターフェース
type CustomerHandler interface {
	Create(c *gin.Context)
	List(c *gin.Context)
	Get(c *gin.Context)
	Update(c *gin.Context)
	Delete(c *gin.Context)
}

// customerHandler の実装
type customerHandler struct {
	customerCreateUsecase     usecase.CustomerCreateUsecase
	customerListUsecase       usecase.CustomerListUsecase
	customerGetUsecase        usecase.CustomerGetUsecase
	customerUpdateUsecase     usecase.CustomerUpdateUsecase
	customerSoftDeleteUsecase usecase.CustomerSoftDeleteUsecase
}

var _ CustomerHandler = (*customerHandler)(nil)

// NewCustomerHandler はcustomerHandlerの初期化を行う
func NewCustomerHandler(
	customerCreateUsecase usecase.CustomerCreateUsecase,
	customerListUsecase usecase.CustomerListUsecase,
	customerGetUsecase usecase.CustomerGetUsecase,
	customerUpdateUsecase usecase.CustomerUpdateUsecase,
	customerSoftDeleteUsecase usecase.CustomerSoftDeleteUsecase,
) CustomerHandler {
	return &customerHandler{
		customerCreateUsecase:     customerCreateUsecase,
		customerListUsecase:       customerListUsecase,
		customerGetUsecase:        customerGetUsecase,
		customerUpdateUsecase:     customerUpdateUsecase,
		customerSoftDeleteUsecase: customerSoftDeleteUsecase,
	}
}

// CreateCustomersBodyRequest は顧客作成のボディリクエスト構造体
type CreateCustomersBodyRequest struct {
	Name      string  `json:"name" binding:"required"`
	UnitPrice *int64  `json:"unitPrice"`
	StartDate *string `json:"startDate"`
	EndDate   *string `json:"endDate"`
}

// Create は顧客を新規作成する
//
//	@Summary		Create Customer
//	@Description	顧客を新規作成する
//	@Tags			customer
//	@Security		BearerAuth
//	@Param			payload	body	CreateCustomersBodyRequest	true	"組織新規作成APIのペイロード"
//	@Accept			json
//	@Produce		json
//	@Success		201	{object}	presenter.CustomerCreateResponse
//	@Failure		400	{object}	nil	"BadRequest"
//	@Failure		401	{object}	nil	"Unauthorized"
//	@Failure		403	{object}	nil	"Forbidden"
//	@Router			/customers [POST]
func (h *customerHandler) Create(c *gin.Context) {
	ctx := c.Request.Context()

	var req CreateCustomersBodyRequest
	err := c.ShouldBindJSON(&req)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	input := usecase.CustomerCreateUsecaseInput{
		Name:      req.Name,
		UnitPrice: req.UnitPrice,
		StartDate: req.StartDate,
		EndDate:   req.EndDate,
	}
	customer, err := h.customerCreateUsecase.Create(ctx, input)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	res := presenter.NewCustomerCreateResponse(customer)
	c.JSON(http.StatusCreated, res)
}

// List は顧客一覧を返却する
//
//	@Summary		List Customers
//	@Description	顧客一覧を返却する
//	@Tags			customer
//	@Security		BearerAuth
//	@Accept			json
//	@Produce		json
//	@Success		200	{array}		presenter.CustomerListResponse
//	@Failure		400	{object}	nil	"BadRequest"
//	@Failure		401	{object}	nil	"Unauthorized"
//	@Router			/customers [GET]
func (h *customerHandler) List(c *gin.Context) {
	ctx := c.Request.Context()

	customers, err := h.customerListUsecase.List(ctx)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	resp := presenter.NewCustomerListResponse(customers)
	c.JSON(http.StatusOK, resp)
}

// CustomerGetURIRequest は顧客情報取得のリクエスト構造体
type CustomerGetURIRequest struct {
	CustomerID uint `uri:"customerID" binding:"required" example:"1"`
}

// Get は顧客情報を返却する
//
//	@Summary		Get Customer
//	@Description	顧客情報を返却する
//	@Tags			customer
//	@Security		BearerAuth
//	@Param			uri	parameter	path	CustomerGetURIRequest	true	"URI parameter"
//	@Accept			json
//	@Produce		json
//	@Success		200	{object}	presenter.CustomerGetResponse
//	@Failure		400	{object}	nil	"BadRequest"
//	@Failure		401	{object}	nil	"Unauthorized"
//	@Router			/customers/{customerID} [GET]
func (h *customerHandler) Get(c *gin.Context) {
	ctx := c.Request.Context()
	var req CustomerGetURIRequest
	if err := c.ShouldBindUri(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	input := usecase.CustomerGetUsecaseInput{
		CustomerID: req.CustomerID,
	}
	customer, err := h.customerGetUsecase.Get(ctx, input)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	resp := presenter.NewCustomerGetResponse(customer)
	c.JSON(http.StatusOK, resp)
}

// CustomerUpdateURIRequest は顧客作成のURIリクエスト構造体
type CustomerUpdateURIRequest struct {
	CustomerID uint `uri:"customerID" binding:"required" example:"1"`
}

// CustomerUpdateBodyRequest は顧客作成のボディリクエスト構造体
type CustomerUpdateBodyRequest struct {
	Name      string  `json:"name" binding:"required"`
	UnitPrice *int64  `json:"unitPrice"`
	StartDate *string `json:"startDate"`
	EndDate   *string `json:"endDate"`
}

// Update は顧客情報を更新する
//
//	@Summary		Update Customer
//	@Description	顧客情報を更新する
//	@Tags			customer
//	@Security		BearerAuth
//	@Param			request	path	CustomerUpdateURIRequest	true	"uri request"
//	@Param			payload	body	CustomerUpdateBodyRequest	true	"payload"
//	@Accept			json
//	@Produce		json
//	@Success		201	{object}	presenter.CustomerUpdateResponse
//	@Success		204	{object}	nil	"NoContent"
//	@Failure		400	{object}	nil	"BadRequest"
//	@Failure		401	{object}	nil	"Unauthorized"
//	@Failure		403	{object}	nil	"Forbidden"
//	@Router			/customers/{customerID} [PUT]
func (h *customerHandler) Update(c *gin.Context) {
	ctx := c.Request.Context()

	var uriReq CustomerGetURIRequest
	if err := c.ShouldBindUri(&uriReq); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	var bodyReq CustomerUpdateBodyRequest
	if err := c.ShouldBindJSON(&bodyReq); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// 更新処理
	input := usecase.CustomerUpdateUsecaseInput{
		CustomerID: uriReq.CustomerID,
		Name:       bodyReq.Name,
		UnitPrice:  bodyReq.UnitPrice,
		StartDate:  bodyReq.StartDate,
		EndDate:    bodyReq.EndDate,
	}
	customer, err := h.customerUpdateUsecase.Update(ctx, input)
	if err != nil {
		if errors.Is(err, errors.New("no update content")) {
			c.JSON(http.StatusNoContent, nil)
		}
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	resp := presenter.NewCustomerUpdateResponse(customer)
	c.JSON(http.StatusCreated, resp)
}

// CustomerDeleteURIRequest はユーザー削除のURIリクエスト構造体
type CustomerDeleteURIRequest struct {
	CustomerID uint `uri:"customerID" binding:"required" example:"1"`
}

// Delete は顧客情報を削除する
//
//	@Summary		Delete Customer
//	@Description	顧客情報を削除する
//	@Tags			customer
//	@Security		BearerAuth
//	@Param			uri	request	path	CustomerDeleteURIRequest	true	"customer ID"
//	@Accept			json
//	@Produce		json
//	@Success		200	{object}	nil
//	@Failure		400	{object}	nil	"BadRequest"
//	@Failure		401	{object}	nil	"Unauthorized"
//	@Failure		403	{object}	nil	"Forbidden"
//	@Router			/customers/{customerID} [DELETE]
func (h *customerHandler) Delete(c *gin.Context) {
	ctx := c.Request.Context()
	var req CustomerDeleteURIRequest
	if err := c.ShouldBindUri(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	input := usecase.CustomerSoftDeleteUsecaseInput{
		CustomerID: req.CustomerID,
	}
	err := h.customerSoftDeleteUsecase.SoftDelete(ctx, input)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	}

	c.JSON(http.StatusNoContent, nil)
}
