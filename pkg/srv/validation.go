package srv

import (
	"context"

	"github.com/ShatteredRealms/go-common-service/pkg/auth"
	commonsrv "github.com/ShatteredRealms/go-common-service/pkg/srv"
	"github.com/WilSimpson/gocloak/v13"
)

func (c *GameServerContext) validateRole(ctx context.Context, role *gocloak.Role) error {
	claims, ok := auth.RetrieveClaims(ctx)
	if !ok {
		return commonsrv.ErrPermissionDenied
	}
	if !claims.HasResourceRole(role, c.Config.Keycloak.ClientId) {
		return commonsrv.ErrPermissionDenied
	}
	return nil
}
