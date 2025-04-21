package usecase

//
//import (
//	"context"
//	"database/sql"
//	"errors"
//	"testing"
//	"time"
//
//	"github.com/DATA-DOG/go-sqlmock"
//	"github.com/ryota1119/time_resport_webapi/internal/domain/entities"
//	domainerrors "github.com/ryota1119/time_resport_webapi/internal/domain/errors"
//	mockrepo "github.com/ryota1119/time_resport_webapi/internal/domain/repository/mock"
//	mocksvc "github.com/ryota1119/time_resport_webapi/internal/domain/service/mock"
//	"github.com/ryota1119/time_resport_webapi/internal/helper/auth_context"
//)
//
//func TestAuthLoginUsecase_Login_Success(t *testing.T) {
//	const (
//		organizationID = 1
//		userID         = 1
//		userName       = "John Doe"
//		email          = "test@test.com"
//		password       = "password"
//		role           = "admin"
//		accessToken    = "access-token"
//		refreshToken   = "refresh-token"
//		accessJti      = "access-jti"
//		refreshJti     = "refresh-jti"
//	)
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
//	ctx := context.Background()
//
//	orgID := entities.OrganizationID(organizationID)
//	ctx = auth_context.SetContextOrganizationID(ctx, orgID)
//
//	user, err := entities.NewUser(userName, email, password, role)
//	if err != nil {
//		t.Fatal(err)
//	}
//	user.ID = userID
//
//	aToken := accessToken
//	rToken := refreshToken
//	aJti := entities.Jti(accessJti)
//	rJti := entities.Jti(refreshJti)
//
//	usecase := NewAuthLoginUsecase(
//		db,
//		&mocksvc.JwtTokenServiceMock{
//			GenerateJwtTokenFunc: func(user *entities.User, organizationID *entities.OrganizationID, expires time.Duration) (*string, *entities.Jti, error) {
//				if expires == 15*time.Minute {
//					return &aToken, &aJti, nil
//				}
//				return &rToken, &rJti, nil
//			},
//		},
//		&mockrepo.AuthRepositoryMock{
//			SaveAccessTokenFunc: func(ctx context.Context, userID entities.UserID, jti *entities.Jti, duration time.Duration) error {
//				return nil
//			},
//			SaveRefreshTokenFunc: func(ctx context.Context, userID entities.UserID, jti *entities.Jti, duration time.Duration) error {
//				return nil
//			},
//		},
//		&mockrepo.UserRepositoryMock{
//			FindByEmailFunc: func(ctx context.Context, tx *sql.Tx, email *entities.UserEmail) (*entities.User, error) {
//				return user, nil
//			},
//		},
//	)
//
//	input := AuthUsecaseLoginInput{
//		Email:    email,
//		Password: password,
//	}
//
//	// authLoginUsecase.Login 実行
//	token, err := usecase.Login(ctx, input)
//	if err != nil {
//		t.Fatal(err)
//	}
//	// アサーション
//	if token.AccessToken != accessToken {
//		t.Errorf("access token expected: %s, got: %s", accessToken, token.AccessToken)
//	}
//	if token.RefreshToken != refreshToken {
//		t.Errorf("refresh token expected: %s, got: %s", refreshToken, token.RefreshToken)
//	}
//
//	if err := sqlMock.ExpectationsWereMet(); err != nil {
//		t.Errorf("there were unfulfilled expectations: %s", err)
//	}
//}
//
//func TestAuthLoginUsecase_Login_UserNotFound(t *testing.T) {
//	const (
//		organizationID = 1
//		email          = "test@test.com"
//		password       = "password"
//	)
//
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
//	sqlMock.ExpectRollback()
//
//	usecase := NewAuthLoginUsecase(
//		db,
//		nil,
//		nil,
//		&mockrepo.UserRepositoryMock{
//			FindByEmailFunc: func(ctx context.Context, tx *sql.Tx, email *entities.UserEmail) (*entities.User, error) {
//				return nil, domainerrors.ErrUserNotFound
//			},
//		},
//	)
//
//	input := AuthUsecaseLoginInput{
//		Email:    email,
//		Password: password,
//	}
//
//	// authLoginUsecase.Login 実行
//	_, err = usecase.Login(ctx, input)
//	if err == nil {
//		t.Fatal("expected an error, got none")
//	}
//
//	if !errors.Is(err, domainerrors.ErrUserNotFound) {
//		t.Errorf("expected: %s, got: %s", domainerrors.ErrUserNotFound, err)
//	}
//
//	if err := sqlMock.ExpectationsWereMet(); err != nil {
//		t.Errorf("there were unfulfilled expectations: %s", err)
//	}
//}
