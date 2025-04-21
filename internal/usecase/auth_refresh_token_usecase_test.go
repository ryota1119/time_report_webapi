package usecase

//
//import (
//	"context"
//	"database/sql"
//	"strconv"
//	"testing"
//	"time"
//
//	"github.com/DATA-DOG/go-sqlmock"
//	"github.com/golang-jwt/jwt/v5"
//	"github.com/ryota1119/time_resport_webapi/internal/domain/entities"
//	mockrepo "github.com/ryota1119/time_resport_webapi/internal/domain/repository/mock"
//	mocksvc "github.com/ryota1119/time_resport_webapi/internal/domain/service/mock"
//	"github.com/ryota1119/time_resport_webapi/internal/helper/auth_context"
//)
//
//func TestAuthLoginUsecase_RefreshToken_Success(t *testing.T) {
//	const (
//		organizationID = 1
//		userID         = 1
//		userName       = "john doe"
//		email          = "test@test.com"
//		password       = "password"
//		role           = "admin"
//		accessToken    = "access-token"
//		refreshToken   = "refresh-token"
//		accessJti      = "access-jti"
//		refreshJti     = "refresh-jti"
//	)
//	orgID := entities.OrganizationID(organizationID)
//	ctx := auth_context.SetContextOrganizationID(context.Background(), orgID)
//
//	db, sqlMock, err := sqlmock.New()
//	if err != nil {
//		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
//	}
//	defer db.Close()
//	// トランザクションの期待
//	sqlMock.ExpectBegin()
//	sqlMock.ExpectCommit()
//
//	user, err := entities.NewUser(userName, email, password, role)
//	if err != nil {
//		t.Fatal(err)
//	}
//	user.ID = userID
//
//	usecase := NewAuthRefreshTokenUsecase(
//		db,
//		&mocksvc.JwtTokenServiceMock{
//			GenerateJwtTokenFunc: func(user *entities.User, organizationID *entities.OrganizationID, expires time.Duration) (*string, *entities.Jti, error) {
//				aToken := accessToken
//				rToken := refreshToken
//				aJti := entities.Jti(accessJti)
//				rJti := entities.Jti(refreshJti)
//				if expires == 15*time.Minute {
//					return &aToken, &aJti, nil
//				}
//				return &rToken, &rJti, nil
//			},
//			ValidateJwtTokenFunc: func(token string) (*entities.CustomClaims, error) {
//				return &entities.CustomClaims{
//					Role:           role,
//					OrganizationID: strconv.Itoa(organizationID),
//					RegisteredClaims: jwt.RegisteredClaims{
//						Subject: strconv.Itoa(userID),
//						ID:      refreshToken,
//					},
//				}, nil
//			},
//		},
//		&mockrepo.AuthRepositoryMock{
//			SaveAccessTokenFunc: func(ctx context.Context, userID entities.UserID, jti *entities.Jti, duration time.Duration) error {
//				return nil
//			},
//			SaveRefreshTokenFunc: func(ctx context.Context, userID entities.UserID, jti *entities.Jti, duration time.Duration) error {
//				return nil
//			},
//			GetUserIDByRefreshTokenFunc: func(ctx context.Context, jti *entities.Jti) (*entities.UserID, error) {
//				return &user.ID, nil
//			},
//		},
//		&mockrepo.UserRepositoryMock{
//			FindFunc: func(ctx context.Context, tx *sql.Tx, userID *entities.UserID) (*entities.User, error) {
//				return user, nil
//			},
//		},
//	)
//
//	input := AuthUsecaseRefreshTokenInput{
//		RefreshToken: refreshToken,
//	}
//
//	// authLoginUsecase.RefreshToken 実行
//	token, err := usecase.RefreshToken(ctx, input)
//	if err != nil {
//		t.Fatal(err)
//	}
//	// アサーション
//	if token.AccessToken != accessToken {
//		t.Errorf("access token expected: %s, got: %s", accessToken, token.AccessToken)
//	}
//	if token.RefreshToken != refreshToken {
//		t.Errorf("refresh token expected: %s, got: %s", accessToken, token.RefreshToken)
//	}
//
//	if err := sqlMock.ExpectationsWereMet(); err != nil {
//		t.Errorf("there were unfulfilled expectations: %s", err)
//	}
//}
//
//func TestAuthLoginUsecase_RefreshToken_InvalidToken(t *testing.T) {
//	const (
//		organizationID = 1
//		refreshToken   = "refresh-token"
//	)
//
//	orgID := entities.OrganizationID(organizationID)
//	ctx := auth_context.SetContextOrganizationID(context.Background(), orgID)
//
//	usecase := NewAuthRefreshTokenUsecase(
//		nil,
//		&mocksvc.JwtTokenServiceMock{
//			ValidateJwtTokenFunc: func(token string) (*entities.CustomClaims, error) {
//				return nil, entities.ErrInvalidToken
//			},
//		},
//		nil,
//		nil,
//	)
//
//	input := AuthUsecaseRefreshTokenInput{
//		RefreshToken: refreshToken,
//	}
//
//	// authLoginUsecase.RefreshToken 実行
//	_, err := usecase.RefreshToken(ctx, input)
//	if err == nil {
//		t.Fatal("expected an error, got none")
//
//	}
//}
