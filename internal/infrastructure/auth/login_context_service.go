package auth

import (
	"context"

	"github.com/ryota1119/time_resport_webapi/internal/domain/entities"
)

// GetOrganizationLoginService 組織情報を取得する
func (a *authService) GetOrganizationLoginService(ctx context.Context, orgCode string) (*entities.Organization, error) {
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

	organizationCode := entities.OrganizationCode(orgCode)

	organization, err := a.organizationRepo.FindByCode(ctx, tx, &organizationCode)
	if err != nil {
		return nil, err
	}
	return organization, nil
}
