package auth

import (
	"context"
	"database/sql"

	"github.com/redis/go-redis/v9"
	"github.com/ryota1119/time_resport/internal/domain/entities"
	"github.com/ryota1119/time_resport/internal/domain/repository"
	"github.com/ryota1119/time_resport/internal/domain/service"
)

// Service AuthServiceのインターフェースを定義
type Service interface {
	// AuthenticateFromToken はトークンを受け取って、ユーザー情報を返す
	AuthenticateFromToken(ctx context.Context, accessToken string) (*entities.User, *entities.Organization, error)
}

// authService
type authService struct {
	db               *sql.DB
	redisClient      *redis.Client
	jwtTokenService  service.JwtTokenService
	authRepo         repository.AuthRepository
	organizationRepo repository.OrganizationRepository
	userRepo         repository.UserRepository
}

var _ Service = (*authService)(nil)

func NewAuthService(
	db *sql.DB,
	redisClient *redis.Client,
	jwtTokenService service.JwtTokenService,
	authRepo repository.AuthRepository,
	organizationRepo repository.OrganizationRepository,
	userRepo repository.UserRepository,
) Service {
	return &authService{
		db:               db,
		redisClient:      redisClient,
		jwtTokenService:  jwtTokenService,
		authRepo:         authRepo,
		organizationRepo: organizationRepo,
		userRepo:         userRepo,
	}
}
