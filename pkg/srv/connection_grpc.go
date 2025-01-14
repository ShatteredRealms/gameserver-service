package srv

import (
	"context"
	"errors"

	"github.com/ShatteredRealms/gameserver-service/pkg/pb"
	"github.com/ShatteredRealms/go-common-service/pkg/auth"
	"github.com/ShatteredRealms/go-common-service/pkg/config"
	"github.com/ShatteredRealms/go-common-service/pkg/log"
	commonpb "github.com/ShatteredRealms/go-common-service/pkg/pb"
	commonsrv "github.com/ShatteredRealms/go-common-service/pkg/srv"
	"github.com/ShatteredRealms/go-common-service/pkg/util"
	"github.com/WilSimpson/gocloak/v13"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

var (
	ConnectionRoles = make([]*gocloak.Role, 0)

	RoleManageConnections = util.RegisterRole(&gocloak.Role{
		Name:        gocloak.StringP("game.connections.manage"),
		Description: gocloak.StringP("Manage gamesever connections"),
	}, &ConnectionRoles)

	RolePlay = util.RegisterRole(&gocloak.Role{
		Name:        gocloak.StringP("game.connections.play"),
		Description: gocloak.StringP("Allows user to connect to servers and play the game as any character they own"),
	}, &ConnectionRoles)

	RolePlayOther = util.RegisterRole(&gocloak.Role{
		Name:        gocloak.StringP("game.connections.play.other"),
		Description: gocloak.StringP("Allows user to connect to servers and play the game as any character that exists"),
	}, &ConnectionRoles)

	RoleConnectionStatus = util.RegisterRole(&gocloak.Role{
		Name:        gocloak.StringP("game.connections.status"),
		Description: gocloak.StringP("Allows to view the status of any character"),
	}, &ConnectionRoles)
)

var (
	ErrCheckCharacterOwnership = errors.New("GS-CH-01")
	ErrCreatePendingConnection = errors.New("GS-CO-01")
)

type connectionServiceServer struct {
	pb.UnimplementedConnectionServiceServer
	Context *GameServerContext
}

// ConnectGameServer implements pb.ConnectionServiceServer.
func (c *connectionServiceServer) ConnectGameServer(
	ctx context.Context,
	request *commonpb.TargetId,
) (*pb.ConnectGameServerResponse, error) {
	claims, ok := auth.RetrieveClaims(ctx)
	if !ok {
		return nil, commonsrv.ErrPermissionDenied
	}
	if !claims.HasResourceRole(RolePlay, c.Context.Config.Keycloak.ClientId) {
		return nil, commonsrv.ErrPermissionDenied
	}

	ok, err := c.Context.CharacterService.DoesOwnCharacter(ctx, request.Id, claims.Subject)
	if err != nil {
		log.Logger.WithContext(ctx).Errorf("code %v: %v", ErrCheckCharacterOwnership, err)
		return nil, status.Error(codes.Internal, ErrCheckCharacterOwnership.Error())
	}
	if !ok && !claims.HasResourceRole(RolePlayOther, c.Context.Config.Keycloak.ClientId) {
		return nil, commonsrv.ErrPermissionDenied
	}

	if c.Context.Config.Mode == config.ModeLocal {
		pc, err := c.Context.ConnectionService.CreatePendingConnection(ctx, "localhost", request.Id)
		if err != nil {
			log.Logger.WithContext(ctx).Errorf("code %v: %v", ErrCreatePendingConnection, err)
			return nil, status.Error(codes.Internal, ErrCreatePendingConnection.Error())
		}
		return &pb.ConnectGameServerResponse{
			Address:      "localhost",
			Port:         7777,
			ConnectionId: pc.Id.String(),
		}, nil
	}

	return nil, status.Error(codes.Unimplemented, "remote connections not supported")
}

// IsCharacterPlaying implements pb.ConnectionServiceServer.
func (c *connectionServiceServer) IsCharacterPlaying(
	ctx context.Context,
	request *commonpb.TargetId,
) (*pb.ConnectionStatus, error) {
	panic("unimplemented")
}

// IsUserPlaying implements pb.ConnectionServiceServer.
func (c *connectionServiceServer) IsUserPlaying(
	ctx context.Context,
	request *commonpb.TargetId,
) (*pb.ConnectionStatus, error) {
	panic("unimplemented")
}

// TransferPlayer implements pb.ConnectionServiceServer.
func (c *connectionServiceServer) TransferPlayer(
	ctx context.Context,
	request *pb.TransferPlayerRequest,
) (*pb.ConnectGameServerResponse, error) {
	panic("unimplemented")
}

// VerifyConnect implements pb.ConnectionServiceServer.
func (c *connectionServiceServer) VerifyConnect(
	ctx context.Context,
	request *pb.VerifyConnectRequest,
) (*commonpb.TargetId, error) {
	panic("unimplemented")
}

func NewConnectionServiceServer(ctx context.Context, srvCtx *GameServerContext) (pb.ConnectionServiceServer, error) {
	return &connectionServiceServer{Context: srvCtx}, srvCtx.CreateRoles(ctx, &DimensionRoles)
}
