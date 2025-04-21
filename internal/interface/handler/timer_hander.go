package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/ryota1119/time_resport_webapi/internal/interface/presenter"
	"github.com/ryota1119/time_resport_webapi/internal/usecase"
)

var _ TimerHandler = (*timerHandler)(nil)

// TimerHandler は timerHandler のインターフェース
type TimerHandler interface {
	Start(c *gin.Context)
	Stop(c *gin.Context)
}

// timerHandler の実装
type timerHandler struct {
	timerStartUsecase usecase.TimerStartUsecase
	timerStopUsecase  usecase.TimerStopUsecase
}

// NewTimerHandler は timerHandler の初期化を行う
func NewTimerHandler(
	timerStartUsecase usecase.TimerStartUsecase,
	timerStopUsecase usecase.TimerStopUsecase,
) TimerHandler {
	return &timerHandler{
		timerStartUsecase: timerStartUsecase,
		timerStopUsecase:  timerStopUsecase,
	}
}

// StartTimerRequest はタイマー作成時のリクエストパラメータ
type StartTimerRequest struct {
	ProjectID uint    `json:"projectID" binding:"required" example:"1"`
	Title     string  `json:"title" binding:"required" example:"Start"`
	Memo      *string `json:"memo" example:"timer memo"`
}

// Start はタイマーを開始する
//
//	@Summary		Start Timer
//	@Description	タイマーを開始する
//	@Tags			Timer
//	@Security		BearerAuth
//	@Param			request	body	StartTimerRequest	true	"start timer payload"
//	@Accept			json
//	@Produce		json
//	@Success		201	{object}	presenter.StartTimerResponse
//	@Failure		400	{object}	nil	"BadRequest"
//	@Failure		401	{object}	nil	"Unauthorized"
//	@Router			/timers/start [POST]
func (h *timerHandler) Start(c *gin.Context) {
	ctx := c.Request.Context()

	var req StartTimerRequest
	err := c.ShouldBindJSON(&req)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	}

	input := usecase.TimerStartUsecaseInput{
		ProjectID: req.ProjectID,
		Title:     req.Title,
		Memo:      req.Memo,
	}
	timer, err := h.timerStartUsecase.Start(ctx, input)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
	}

	res := presenter.NewStartTimerResponse(timer)
	c.JSON(http.StatusCreated, res)
}

// StopTimerURIRequest はタイマー終了時のURIパラメータ
type StopTimerURIRequest struct {
	ID uint `uri:"timerID" binding:"required" example:"1"`
}

// Stop はタイマーを終了する
//
//	@Summary		Stop Timer
//	@Description	タイマーを終了する
//	@Tags			Timer
//	@Security		BearerAuth
//	@Param			request	path	StopTimerURIRequest	true	"stop timer uri parameter"
//	@Accept			json
//	@Produce		json
//	@Success		200	{object}	presenter.StartTimerResponse
//	@Failure		400	{object}	nil	"BadRequest"
//	@Failure		401	{object}	nil	"Unauthorized"
//	@Router			/timers/{timerID}/stop [POST]
func (h *timerHandler) Stop(c *gin.Context) {
	ctx := c.Request.Context()

	var req StopTimerURIRequest
	err := c.ShouldBindUri(&req)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	}

	input := usecase.TimerStopUsecaseInput{
		TimerID: req.ID,
	}
	timer, err := h.timerStopUsecase.Stop(ctx, input)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
	}

	res := presenter.NewStopTimerResponse(timer)
	c.JSON(http.StatusOK, res)
}
