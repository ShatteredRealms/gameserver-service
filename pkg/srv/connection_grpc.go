package srv

import (
	"context"
	"errors"
	"fmt"

	"github.com/ShatteredRealms/gameserver-service/pkg/pb"
	"github.com/ShatteredRealms/go-common-service/pkg/config"
	"github.com/ShatteredRealms/go-common-service/pkg/log"
	commonpb "github.com/ShatteredRealms/go-common-service/pkg/pb"
	"github.com/ShatteredRealms/go-common-service/pkg/util"
	"github.com/WilSimpson/gocloak/v13"
	"github.com/google/uuid"
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
		Description: gocloak.StringP("Allows to view the status of own character"),
	}, &ConnectionRoles)
	RoleConnectionStatusAll = util.RegisterRole(&gocloak.Role{
		Name:        gocloak.StringP("game.connections.status.all"),
		Description: gocloak.StringP("Allows to view the status of any and all characters"),
	}, &ConnectionRoles)
)

var (
	ErrCheckCharacterOwnership = errors.New("GS-CH-01")
	ErrGetCharacters           = errors.New("GS-CH-02")
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
	character, err := c.Context.IsOwnerWithRoleOrAllRole(ctx, request.Id, RolePlay, RolePlayOther)
	if err != nil {
		return nil, err
	}

	playing, err := c.isUserPlayer(ctx, character.OwnerId.String())
	if err != nil {
		return nil, err
	}

	if playing {
		return nil, status.Error(codes.FailedPrecondition, "account is already playing")
	}

	if c.Context.Config.Mode != config.ModeProduction {
		pc, err := c.Context.ConnectionService.CreatePendingConnection(ctx, request.Id, "sro-local-default")
		if err != nil {
			log.Logger.WithContext(ctx).Errorf("code %v: %v", ErrCreatePendingConnection, err)
			return nil, status.Error(codes.Internal, ErrCreatePendingConnection.Error())
		}
		return &pb.ConnectGameServerResponse{
			Address:      "127.0.0.1",
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
	_, err := c.Context.IsOwnerWithRoleOrAllRole(ctx, request.Id, RoleConnectionStatus, RoleConnectionStatusAll)
	if err != nil {
		return nil, err
	}

	playing, err := c.Context.GsmService.AnyCharactersConneted(ctx, []string{request.Id})
	if err != nil {
		log.Logger.WithContext(ctx).Errorf("code %v: %v", ErrCheckCharacterOwnership, err)
		return nil, status.Error(codes.Internal, ErrCheckCharacterOwnership.Error())
	}

	return &pb.ConnectionStatus{
		Online: playing,
	}, nil
}

// IsUserPlaying implements pb.ConnectionServiceServer.
func (c *connectionServiceServer) IsUserPlaying(
	ctx context.Context,
	request *commonpb.TargetId,
) (*pb.ConnectionStatus, error) {
	err := c.Context.IsSelfWithRoleOrAllRole(ctx, request.Id, RoleConnectionStatus, RoleConnectionStatusAll)
	if err != nil {
		return nil, err
	}

	playing, err := c.isUserPlayer(ctx, request.Id)
	if err != nil {
		return nil, err
	}

	return &pb.ConnectionStatus{
		Online: playing,
	}, nil
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
	err := c.Context.validateRole(ctx, RoleManageConnections)
	if err != nil {
		return nil, err
	}

	id, err := uuid.Parse(request.ConnectionId)
	if err != nil {
		return nil, status.Error(codes.InvalidArgument, fmt.Sprintf("invalid connection id: %v", err))
	}

	pc, err := c.Context.ConnectionService.CheckPlayerConnection(ctx, &id, request.ServerName)
	if err != nil {
		return nil, status.Error(codes.InvalidArgument, err.Error())
	}

	return &commonpb.TargetId{Id: pc.CharacterId.String()}, nil
}

func NewConnectionServiceServer(ctx context.Context, srvCtx *GameServerContext) (pb.ConnectionServiceServer, error) {
	return &connectionServiceServer{Context: srvCtx}, srvCtx.CreateRoles(ctx, &ConnectionRoles)
}

func (c *connectionServiceServer) isUserPlayer(ctx context.Context, userId string) (bool, error) {
	playing, err := c.Context.GsmService.AnyUsersConneted(ctx, []string{userId})
	if err != nil {
		log.Logger.WithContext(ctx).Errorf("code %v: %v", ErrCheckCharacterOwnership, err)
		return false, status.Error(codes.Internal, ErrCheckCharacterOwnership.Error())
	}

	return playing, nil
}
