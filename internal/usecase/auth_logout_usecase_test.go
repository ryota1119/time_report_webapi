package usecase

//
//import (
//	"context"
//
//	"github.com/ryota1119/time_resport_webapi/internal/domain/entities"
//	mockrepo "github.com/ryota1119/time_resport_webapi/internal/domain/repository/mock"
//	"github.com/ryota1119/time_resport_webapi/internal/helper/auth_context"
//
//	"testing"
//)
//
//func TestAuthLogoutUsecase_Logout(t *testing.T) {
//	const (
//		organizationID = 1
//	)
//	orgID := entities.OrganizationID(organizationID)
//	ctx := auth_context.SetContextOrganizationID(context.Background(), orgID)
//
//	usecase := NewAuthLogoutUsecase(
//		&mockrepo.AuthRepositoryMock{
//			DeleteTokenFunc: func(ctx context.Context) error {
//				return nil
//			},
//		},
//	)
//
//	err := usecase.Logout(ctx)
//	if err != nil {
//		t.Fatalf("%+v", err)
//	}
//}
