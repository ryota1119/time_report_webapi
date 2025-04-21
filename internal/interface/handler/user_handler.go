package handler

import (
	"errors"
	"github.com/ryota1119/time_resport_webapi/internal/helper/auth_context"
	"net/http"

	"github.com/ryota1119/time_resport_webapi/internal/domain/entities"

	"github.com/gin-gonic/gin"
	"github.com/ryota1119/time_resport_webapi/internal/interface/presenter"
	"github.com/ryota1119/time_resport_webapi/internal/usecase"
)

// UserHandler はuserhandlerのインターフェース
type UserHandler interface {
	Create(c *gin.Context)
	List(c *gin.Context)
	Get(c *gin.Context)
	Update(c *gin.Context)
	Delete(c *gin.Context)
	Me(c *gin.Context)
}

// userHandler の実装
type userHandler struct {
	userCreateUsecase     usecase.UserCreateUsecase
	userListUsecase       usecase.UserListUsecase
	userGetUsecase        usecase.UserGetUsecase
	userUpdateUsecase     usecase.UserUpdateUsecase
	userSoftDeleteUsecase usecase.UserSoftDeleteUsecase
}

var _ UserHandler = (*userHandler)(nil)

// NewUserHandler はuserHandlerの初期化を行う
func NewUserHandler(
	userCreateUsecase usecase.UserCreateUsecase,
	userListUsecase usecase.UserListUsecase,
	userGetUsecase usecase.UserGetUsecase,
	userUpdateUsecase usecase.UserUpdateUsecase,
	userSoftDeleteUsecase usecase.UserSoftDeleteUsecase,
) UserHandler {
	return &userHandler{
		userCreateUsecase:     userCreateUsecase,
		userListUsecase:       userListUsecase,
		userGetUsecase:        userGetUsecase,
		userUpdateUsecase:     userUpdateUsecase,
		userSoftDeleteUsecase: userSoftDeleteUsecase,
	}
}

// UserCreateBodyRequest はユーザー作成のボディリクエスト構造体
type UserCreateBodyRequest struct {
	Name     string `json:"name" binding:"required" example:"山田太郎"`
	Email    string `json:"email" binding:"required" example:"example@example.com"`
	Password string `json:"password" binding:"required" example:"password"`
	Role     string `json:"role" binding:"required,oneof=admin member" example:"admin"`
}

// Create はユーザーを新規作成する
//
//	@Summary		Create User
//	@Description	ユーザーを新規作成する
//	@Tags			user
//	@Security		BearerAuth
//	@Param			payload	body	UserCreateBodyRequest	true	"ユーザー新規作成APIのペイロード"
//	@Accept			json
//	@Produce		json
//	@Success		201	{object}	presenter.UserCreateResponse
//	@Failure		400	{object}	nil	"BadRequest"
//	@Failure		401	{object}	nil	"Unauthorized"
//	@Failure		403	{object}	nil	"Forbidden"
//	@Router			/users [POST]
func (h *userHandler) Create(c *gin.Context) {
	ctx := c.Request.Context()

	var req UserCreateBodyRequest
	err := c.ShouldBindJSON(&req)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	input := usecase.UserCreateUsecaseInput{
		Name:     req.Name,
		Email:    req.Email,
		Password: req.Password,
		Role:     req.Role,
	}
	user, err := h.userCreateUsecase.Create(ctx, input)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	resp := presenter.NewUserCreateResponse(user)
	c.JSON(http.StatusCreated, resp)
}

// List はユーザー一覧を返却する
//
//	@Summary		List Users
//	@Description	ユーザー一覧を返却する
//	@Tags			user
//	@Security		BearerAuth
//	@Accept			json
//	@Produce		json
//	@Success		200	{array}		presenter.UserListResponse
//	@Failure		400	{object}	nil	"BadRequest"
//	@Failure		401	{object}	nil	"Unauthorized"
//	@Router			/users [GET]
func (h *userHandler) List(c *gin.Context) {
	ctx := c.Request.Context()
	users, err := h.userListUsecase.List(ctx)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	resp := presenter.NewUserListResponse(users)
	c.JSON(http.StatusOK, resp)
}

// UserGetURIRequest はユーザー取得のリクエスト構造体
type UserGetURIRequest struct {
	UserID int `uri:"userID" binding:"required" example:"1"`
}

// Get はユーザー情報を取得する
//
//	@Summary		Get User
//	@Description	ユーザー情報を返却する
//	@Tags			user
//	@Security		BearerAuth
//	@Param			uri	parameter	path	UserGetURIRequest	true	"URI parameter"
//	@Accept			json
//	@Produce		json
//	@Success		200	{object}	presenter.UserGetResponse
//	@Failure		400	{object}	nil	"BadRequest"
//	@Failure		401	{object}	nil	"Unauthorized"
//	@Router			/users/{userID} [GET]
func (h *userHandler) Get(c *gin.Context) {
	ctx := c.Request.Context()
	var req UserGetURIRequest
	if err := c.ShouldBindUri(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	input := usecase.UserGetUsecaseInput{
		UserID: req.UserID,
	}
	user, err := h.userGetUsecase.Get(ctx, input)
	if err != nil {
		if errors.Is(err, entities.ErrUserNotFound) {
			c.JSON(http.StatusNotFound, gin.H{"error": entities.ErrUserNotFound.Error()})
			return
		}
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	resp := presenter.NewUserGetResponse(user)
	c.JSON(http.StatusOK, resp)
}

// UserUpdateURIRequest はユーザー作成のURIリクエスト構造体
type UserUpdateURIRequest struct {
	UserID int `uri:"userID" binding:"required" example:"1"`
}

// UserUpdateBodyRequest はユーザー作成のボディリクエスト構造体
type UserUpdateBodyRequest struct {
	Name  string `json:"name" binding:"required" example:"山田太郎"`
	Email string `json:"email" binding:"required" example:"example@example.com"`
	Role  string `json:"role" binding:"required,oneof=admin member" example:"admin"`
}

// Update はユーザー情報を更新する
//
//	@Summary		Update User
//	@Description	ユーザー情報を更新する
//	@Tags			user
//	@Security		BearerAuth
//	@Param			request	path	UserUpdateURIRequest	true	"uri parameter"
//	@Param			request	body	UserUpdateBodyRequest	true	"payload"
//	@Accept			json
//	@Produce		json
//	@Success		200	{object}	presenter.UserUpdateResponse
//	@Success		204	{object}	nil	"NoContent"
//	@Failure		400	{object}	nil	"BadRequest"
//	@Failure		401	{object}	nil	"Unauthorized"
//	@Failure		403	{object}	nil	"Forbidden"
//	@Router			/users/{userID} [PUT]
func (h *userHandler) Update(c *gin.Context) {
	ctx := c.Request.Context()
	var uriReq UserUpdateURIRequest
	if err := c.ShouldBindUri(&uriReq); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	var bodyReq UserUpdateBodyRequest
	if err := c.ShouldBindJSON(&bodyReq); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	input := usecase.UserUpdateUsecaseInput{
		UserID: uriReq.UserID,
		Name:   bodyReq.Name,
		Email:  bodyReq.Email,
		Role:   bodyReq.Role,
	}
	user, err := h.userUpdateUsecase.Update(ctx, input)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	resp := presenter.NewUserUpdateResponse(user)
	c.JSON(http.StatusOK, resp)
}

// UserDeleteURIRequest はユーザー削除のURIリクエスト構造体
type UserDeleteURIRequest struct {
	UserID int `uri:"userID" binding:"required" example:"1"`
}

// Delete はユーザー情報を削除する
//
//	@Summary		Delete User
//	@Description	ユーザー情報を削除する
//	@Tags			user
//	@Security		BearerAuth
//	@Param			request	path	UserDeleteURIRequest	true	"user ID"
//	@Accept			json
//	@Produce		json
//	@Success		200	{object}	nil
//	@Failure		400	{object}	nil	"BadRequest"
//	@Failure		401	{object}	nil	"Unauthorized"
//	@Failure		403	{object}	nil	"Forbidden"
//	@Router			/users/{userID} [DELETE]
func (h *userHandler) Delete(c *gin.Context) {
	ctx := c.Request.Context()

	var req UserDeleteURIRequest
	if err := c.ShouldBindUri(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	input := usecase.UserSoftDeleteUsecaseInput{
		UserID: req.UserID,
	}
	err := h.userSoftDeleteUsecase.SoftDelete(ctx, input)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, nil)
}

// Me はログインしているユーザー情報を返却する
//
//	@Summary		Get User Myself
//	@Description	ログインしているユーザー情報を返却する
//	@Tags			user
//	@Security		BearerAuth
//	@Accept			json
//	@Produce		json
//	@Success		200	{object}	presenter.UserGetMeResponse
//	@Failure		400	{object}	nil	"BadRequest"
//	@Failure		401	{object}	nil	"Unauthorized"
//	@Router			/users/me [GET]
func (h *userHandler) Me(c *gin.Context) {
	ctx := c.Request.Context()

	userID := auth_context.ContextUserID(ctx)
	input := usecase.UserGetUsecaseInput{
		UserID: userID.Int(),
	}
	user, err := h.userGetUsecase.Get(ctx, input)
	if err != nil {
		if errors.Is(err, entities.ErrUserNotFound) {
			c.JSON(http.StatusNotFound, gin.H{"error": entities.ErrUserNotFound.Error()})
			return
		}
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	resp := presenter.NewUserGetMeResponse(user)
	c.JSON(http.StatusOK, resp)
}
