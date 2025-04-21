package usecase

import (
	"context"
	"database/sql"

	"github.com/ryota1119/time_resport_webapi/internal/domain/entities"
	"github.com/ryota1119/time_resport_webapi/internal/domain/repository"
)

// OrganizationRegisterUsecase organizationRegisterUsecase のインターフェースを定義
type OrganizationRegisterUsecase interface {
	// Register は組織情報を作成する
	Register(ctx context.Context, input OrganizationUsecaseRegisterInput) (*entities.Organization, error)
}

// organizationRegisterUsecase ユースケース
type organizationRegisterUsecase struct {
	db               *sql.DB
	organizationRepo repository.OrganizationRepository
	userRepo         repository.UserRepository
}

var _ OrganizationRegisterUsecase = (*organizationRegisterUsecase)(nil)

func NewOrganizationRegisterUsecase(
	db *sql.DB,
	organizationRepo repository.OrganizationRepository,
	userRepo repository.UserRepository,
) OrganizationRegisterUsecase {
	return &organizationRegisterUsecase{
		db:               db,
		organizationRepo: organizationRepo,
		userRepo:         userRepo,
	}
}

// OrganizationUsecaseRegisterInput は organizationUsecase.Register のインプット
type OrganizationUsecaseRegisterInput struct {
	OrganizationName string
	OrganizationCode string
	UserName         string
	UserEmail        string
	Password         string
}

// Register は組織情報を作成する
func (a *organizationRegisterUsecase) Register(ctx context.Context, input OrganizationUsecaseRegisterInput) (*entities.Organization, error) {
	tx, err := a.db.BeginTx(ctx, nil)
	if err != nil {
		return nil, err
	}
	defer func() {
		if err != nil {
			_ = tx.Rollback()
		} else {
			err = tx.Commit()
		}
	}()

	// 組織情報保存処理
	organization := entities.NewOrganization(input.OrganizationName, input.OrganizationCode)

	// 既存の組織を取得し、あればエラーを返却
	_, err = a.organizationRepo.FindByCode(ctx, tx, &organization.Code)
	if err == nil {
		return nil, entities.ErrOrganizationAlreadyExists
	}

	// 組織情報を作成する
	organizationID, err := a.organizationRepo.Create(ctx, tx, organization)
	if err != nil {
		return nil, err
	}

	// コンテキストに組織IDを格納
	ctx = context.WithValue(ctx, "organization_id", *organizationID)

	// ユーザー情報保存処理
	user, err := entities.NewUser(
		input.UserName,
		input.UserEmail,
		input.Password,
		entities.AdminRole.String(),
	)

	// 既存のユーザーを取得し、あればエラーを返却
	_, err = a.userRepo.FindByEmail(ctx, tx, &user.Email, organizationID)
	if err == nil {
		return nil, entities.ErrUserAlreadyExists
	}
	// ユーザー情報を作成する
	_, err = a.userRepo.Create(ctx, tx, user)
	if err != nil {
		return nil, err
	}

	return organization, nil
}
