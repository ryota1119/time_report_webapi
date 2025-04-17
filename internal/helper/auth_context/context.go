package auth_context

import (
	"context"
	"github.com/ryota1119/time_resport/internal/domain/entities"
)

type contextKey string

const userIDKey contextKey = "user_id"
const organizationIDKey contextKey = "organization_id"
const roleKey contextKey = "role"

// ContextUserID は context.Context から UserID を取得する
func ContextUserID(ctx context.Context) entities.UserID {
	v, ok := ctx.Value(string(userIDKey)).(entities.UserID)
	if !ok {
		return entities.UserID(0)
	}
	return v
}

// SetContextOrganizationID は context.Context から OrganizationID を取得する
func SetContextOrganizationID(ctx context.Context, orgId entities.OrganizationID) context.Context {
	return context.WithValue(ctx, "organization_id", orgId.String())
}

// ContextOrganizationID は context.Context から OrganizationID を取得する
func ContextOrganizationID(ctx context.Context) entities.OrganizationID {
	v := ctx.Value(string(organizationIDKey))

	orgID, ok := v.(entities.OrganizationID)
	if !ok {
		return entities.OrganizationID(0)
	}

	return orgID
}

// ContextUserRole は context.Context から UserRole を取得する
func ContextUserRole(ctx context.Context) entities.Role {
	v, ok := ctx.Value(string(roleKey)).(entities.Role)
	if !ok {
		return entities.UnknownRole
	}
	return v
}
