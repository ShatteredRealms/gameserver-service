package srv

import (
	"context"

	"github.com/ShatteredRealms/gameserver-service/pkg/pb"
	"github.com/ShatteredRealms/go-common-service/pkg/util"
	"github.com/WilSimpson/gocloak/v13"
	"google.golang.org/protobuf/types/known/emptypb"
)

var (
	GameServerDataRoles = make([]*gocloak.Role, 0)

	RoleGameServerDataView = util.RegisterRole(&gocloak.Role{
		Name:        gocloak.StringP("gameserver.data.view"),
		Description: gocloak.StringP("View game server data"),
	}, &GameServerDataRoles)
)

type gameServerDataServiceServer struct {
	pb.UnimplementedGameServerDataServiceServer
	Context *GameServerContext
}

// GetGameServerDetails implements pb.GameServerDataServiceServer.
func (g *gameServerDataServiceServer) GetGameServerDetails(ctx context.Context, _ *emptypb.Empty) (*pb.GameServerDetails, error) {
	if err := g.Context.validateRole(ctx, RoleGameServerDataView); err != nil {
		return nil, err
	}

	if !g.Context.UsingAgones() {
		return &pb.GameServerDetails{Count: -1}, nil
	}

	details, err := g.Context.GsmService.CountGameServers(ctx)
	if err != nil {
		return nil, err
	}

	return &pb.GameServerDetails{
		Count: int32(details),
	}, nil
}

func NewGameServerDataServiceServer(ctx context.Context, srvCtx *GameServerContext) (pb.GameServerDataServiceServer, error) {
	err := srvCtx.CreateRoles(ctx, &GameServerDataRoles)
	if err != nil {
		return nil, err
	}
	return &gameServerDataServiceServer{Context: srvCtx}, nil
}
