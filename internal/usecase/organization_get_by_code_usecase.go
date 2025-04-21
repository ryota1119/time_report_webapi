package usecase

import (
	"context"
	"database/sql"

	"github.com/ryota1119/time_resport_webapi/internal/domain/repository"

	"github.com/ryota1119/time_resport_webapi/internal/domain/entities"
)

// OrganizationGetByCodeUsecase organizationGetByCodeUsecase のインターフェースを定義
type OrganizationGetByCodeUsecase interface {
	// GetByCode は組織情報を作成する
	GetByCode(ctx context.Context, input OrganizationGetByCodeUsecaseInput) (*entities.Organization, error)
}

// organizationGetByCodeUsecase ユースケース
type organizationGetByCodeUsecase struct {
	db               *sql.DB
	organizationRepo repository.OrganizationRepository
}

var _ OrganizationGetByCodeUsecase = (*organizationGetByCodeUsecase)(nil)

func NewOrganizationGetByCodeUsecase(
	db *sql.DB,
	organizationRepo repository.OrganizationRepository,
) OrganizationGetByCodeUsecase {
	return &organizationGetByCodeUsecase{
		db:               db,
		organizationRepo: organizationRepo,
	}
}

// OrganizationGetByCodeUsecaseInput は organizationUsecase.GetByCode のインプット
type OrganizationGetByCodeUsecaseInput struct {
	OrganizationCode string
}

// GetByCode は組織情報を作成する
func (a *organizationGetByCodeUsecase) GetByCode(ctx context.Context, input OrganizationGetByCodeUsecaseInput) (*entities.Organization, error) {
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
	organizationCode := entities.OrganizationCode(input.OrganizationCode)

	// 既存の組織を取得し、あればエラーを返却
	organization, err := a.organizationRepo.FindByCode(ctx, tx, &organizationCode)
	if err != nil {
		return nil, err
	}

	return organization, nil
}
