package handler

import (
	"errors"
	"net/http"

	"github.com/gin-gonic/gin"
	domainerrors "github.com/ryota1119/time_resport/internal/domain/errors"
	"github.com/ryota1119/time_resport/internal/interface/presenter"
	"github.com/ryota1119/time_resport/internal/usecase"
)

// ProjectHandler はprojectHandlerのインターフェース
type ProjectHandler interface {
	Create(c *gin.Context)
	List(c *gin.Context)
	Get(c *gin.Context)
	Update(c *gin.Context)
	Delete(c *gin.Context)
}

// projectHandler の実装
type projectHandler struct {
	projectUsecase usecase.ProjectUsecase
}

var _ ProjectHandler = (*projectHandler)(nil)

// NewProjectHandler はprojectHandlerの初期化を行う
func NewProjectHandler(
	projectUsecase usecase.ProjectUsecase,
) ProjectHandler {
	return &projectHandler{
		projectUsecase: projectUsecase,
	}
}

// CreateProjectsRequest はプロジェクト作成時のリクエストデータ
type CreateProjectsRequest struct {
	CustomerID uint    `json:"customer_id" binding:"required"`
	Name       string  `json:"name" binding:"required"`
	UnitPrice  *int64  `json:"unitPrice"`
	StartDate  *string `json:"startDate"`
	EndDate    *string `json:"endDate"`
}

// Create はプロジェクトを新規作成する
//
//	@Summary		Create Project
//	@Description	プロジェクトを新規作成する
//	@Tags			project
//	@Security		BearerAuth
//	@Param			request	body	CreateProjectsRequest	true	"project create payload"
//	@Accept			json
//	@Produce		json
//	@Success		201	{object}	presenter.ProjectCreateResponse
//	@Failure		400	{object}	nil	"BadRequest"
//	@Failure		401	{object}	nil	"Unauthorized"
//	@Failure		403	{object}	nil	"Forbidden"
//	@Router			/projects [POST]
func (h *projectHandler) Create(c *gin.Context) {
	ctx := c.Request.Context()

	var req CreateProjectsRequest
	err := c.ShouldBindJSON(&req)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	input := usecase.CreateProjectUsecaseInput{
		CustomerID: req.CustomerID,
		Name:       req.Name,
		UnitPrice:  req.UnitPrice,
		StartDate:  req.StartDate,
		EndDate:    req.EndDate,
	}
	project, err := h.projectUsecase.Create(ctx, input)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	res := presenter.NewProjectCreateResponse(project)
	c.JSON(http.StatusCreated, res)
}

// List はプロジェクト一覧を返却する
//
//	@Summary		List Projects
//	@Description	プロジェクト一覧を返却する
//	@Tags			project
//	@Security		BearerAuth
//	@Accept			json
//	@Produce		json
//	@Success		200	{array}		presenter.ProjectListResponse
//	@Failure		400	{object}	nil	"BadRequest"
//	@Failure		401	{object}	nil	"Unauthorized"
//	@Router			/projects [GET]
func (h *projectHandler) List(c *gin.Context) {
	ctx := c.Request.Context()

	projects, err := h.projectUsecase.List(ctx)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	resp := presenter.NewProjectListResponse(projects)
	c.JSON(http.StatusOK, resp)
}

// ProjectGetURIRequest はプロジェクト情報取得のURIリクエスト構造体
type ProjectGetURIRequest struct {
	ProjectID uint `uri:"projectID" binding:"required" example:"1"`
}

// Get はプロジェクト情報を返却する
//
//	@Summary		Get Project
//	@Description	プロジェクト情報を返却する
//	@Tags			project
//	@Security		BearerAuth
//	@Param			uri	parameter	path	ProjectGetURIRequest	true	"プロジェクトID"
//	@Accept			json
//	@Produce		json
//	@Success		200	{object}	presenter.ProjectGetResponse
//	@Failure		400	{object}	nil	"BadRequest"
//	@Failure		401	{object}	nil	"Unauthorized"
//	@Router			/projects/{projectID} [GET]
func (h *projectHandler) Get(c *gin.Context) {
	ctx := c.Request.Context()
	var req ProjectGetURIRequest
	if err := c.ShouldBindUri(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": ErrBadRequest.Error()})
		return
	}

	input := usecase.GetProjectUsecaseInput{
		ProjectID: req.ProjectID,
	}
	project, err := h.projectUsecase.Get(ctx, input)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": ErrBadRequest.Error()})
		return
	}

	resp := presenter.NewProjectGetResponse(project)
	c.JSON(http.StatusOK, resp)
}

// ProjectUpdateURIRequest はプロジェクト更新のURIリクエスト構造体
type ProjectUpdateURIRequest struct {
	ProjectID uint `uri:"projectID" binding:"required" example:"1"`
}

// ProjectUpdateBodyRequest はプロジェクト更新のボディリクエスト構造体
type ProjectUpdateBodyRequest struct {
	CustomerID uint    `json:"customer_id" binding:"required"`
	Name       string  `json:"name"`
	UnitPrice  *int64  `json:"unitPrice"`
	StartDate  *string `json:"startDate"`
	EndDate    *string `json:"endDate"`
}

// Update はプロジェクト情報を更新する
//
//	@Summary		Update Project
//	@Description	プロジェクト情報を更新する
//	@Tags			project
//	@Security		BearerAuth
//	@Param			uri		path	ProjectUpdateURIRequest		true	"プロジェクトID"
//	@Param			payload	body	ProjectUpdateBodyRequest	true	"update projects payload"
//	@Accept			json
//	@Produce		json
//	@Success		201	{object}	presenter.ProjectUpdateResponse
//	@Success		204	{object}	nil	"NoContent"
//	@Failure		400	{object}	nil	"BadRequest"
//	@Failure		401	{object}	nil	"Unauthorized"
//	@Failure		403	{object}	nil	"Forbidden"
//	@Router			/projects/{projectID} [PUT]
func (h *projectHandler) Update(c *gin.Context) {
	ctx := c.Request.Context()

	var uriReq ProjectUpdateURIRequest
	if err := c.ShouldBindUri(&uriReq); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": ErrBadRequest.Error()})
		return
	}

	var bodyReq ProjectUpdateBodyRequest
	if err := c.ShouldBindJSON(&bodyReq); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": ErrBadRequest.Error()})
		return
	}

	// 更新処理
	input := usecase.UpdateProjectUsecaseInput{
		ProjectID:  uriReq.ProjectID,
		CustomerID: bodyReq.CustomerID,
		Name:       bodyReq.Name,
		UnitPrice:  bodyReq.UnitPrice,
		StartDate:  bodyReq.StartDate,
		EndDate:    bodyReq.EndDate,
	}
	project, err := h.projectUsecase.Update(ctx, input)
	if err != nil {
		if errors.Is(err, domainerrors.ErrNoContentUpdated) {
			c.JSON(http.StatusNoContent, nil)
		}
		c.JSON(http.StatusBadRequest, gin.H{"error": ErrBadRequest.Error()})
		return
	}

	resp := presenter.NewProjectUpdateResponse(project)
	c.JSON(http.StatusCreated, resp)
}

// ProjectDeleteURIRequest はプロジェクト削除のURIリクエスト構造体
type ProjectDeleteURIRequest struct {
	ProjectID uint `uri:"projectID" binding:"required" example:"1"`
}

// Delete はプロジェクト情報を削除する
//
//	@Summary		Delete Project
//	@Description	プロジェクト情報を削除する
//	@Tags			project
//	@Security		BearerAuth
//	@Param			uri	path	ProjectDeleteURIRequest	true	"project ID"
//	@Accept			json
//	@Produce		json
//	@Success		204	{object}	nil
//	@Failure		400	{object}	nil	"BadRequest"
//	@Failure		401	{object}	nil	"Unauthorized"
//	@Failure		403	{object}	nil	"Forbidden"
//	@Router			/projects/{projectID} [DELETE]
func (h *projectHandler) Delete(c *gin.Context) {
	ctx := c.Request.Context()

	var req ProjectDeleteURIRequest
	if err := c.ShouldBindUri(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": ErrBadRequest.Error()})
		return
	}

	input := usecase.DeleteProjectUsecaseInput{
		ProjectID: req.ProjectID,
	}
	err := h.projectUsecase.SoftDelete(ctx, input)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": ErrBadRequest.Error()})
	}

	c.JSON(http.StatusNoContent, nil)
}
