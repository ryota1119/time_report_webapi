package handler

//
//import (
//	"bytes"
//	"context"
//	"encoding/json"
//	"github.com/gin-gonic/gin"
//	"github.com/ryota1119/time_resport/internal/domain/entities"
//	"github.com/ryota1119/time_resport/internal/usecase"
//	"github.com/ryota1119/time_resport/internal/usecase/mock"
//	"net/http"
//	"net/http/httptest"
//	"testing"
//	"time"
//)
//
//func TestAuthHandler_Login_Success(t *testing.T) {
//	gin.SetMode(gin.TestMode)
//
//	token := &entities.NewAuthToken{
//		AccessToken:  "access-token",
//		RefreshToken: "refresh-token",
//		ExpiresAt:    time.Now().Add(15 * time.Minute),
//	}
//
//	authLoginUsecase := &mock.AuthLoginUsecaseMock{
//		LoginFunc: func(ctx context.Context, input usecase.AuthLoginUsecaseLoginInput) (*entities.AuthToken, error) {
//			return token, nil
//		},
//	}
//
//	authHandler := NewAuthHandler(nil, authLoginUsecase)
//
//	req := AuthLoginBodyRequest{
//		Email:    "test@example.com",
//		Password: "password",
//	}
//	body, _ := json.Marshal(req)
//
//	w := httptest.NewRecorder()
//	c, _ := gin.CreateTestContext(w)
//	c.Request, _ = http.NewRequest(http.MethodPost, "/my_organization_code/login", bytes.NewReader(body))
//	c.Request.Header.Set("Authorization", "Bearer access-token")
//
//	authHandler.Login(c)
//
//	// アサート
//	if w.Code != http.StatusOK {
//		t.Fatalf("expected 200, got %d", w.Code)
//	}
//	if !bytes.Contains(w.Body.Bytes(), []byte("access-token")) {
//		t.Errorf("response body does not contain access token")
//	}
//}
