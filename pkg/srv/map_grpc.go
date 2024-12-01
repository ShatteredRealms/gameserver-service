package srv

import (
	"context"
	"errors"

	"github.com/ShatteredRealms/gameserver-service/pkg/model/game"
	"github.com/ShatteredRealms/gameserver-service/pkg/pb"
	"github.com/ShatteredRealms/go-common-service/pkg/bus/gameserver/mapbus"
	"github.com/ShatteredRealms/go-common-service/pkg/log"
	commonpb "github.com/ShatteredRealms/go-common-service/pkg/pb"
	"github.com/ShatteredRealms/go-common-service/pkg/util"
	"github.com/WilSimpson/gocloak/v13"
	"github.com/google/uuid"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/types/known/emptypb"
)

var (
	MapRoles = make([]*gocloak.Role, 0)

	RoleMapManage = util.RegisterRole(&gocloak.Role{
		Name:        gocloak.StringP("map.manage"),
		Description: gocloak.StringP("Manage maps"),
	}, &MapRoles)
)

var (
	ErrMapCreate   = errors.New("failed to create map")
	ErrMapDelete   = errors.New("failed to delete map")
	ErrMapEdit     = errors.New("failed to edit map")
	ErrMapLookup   = errors.New("failed to lookup map")
	ErrMapNotExist = errors.New("map does not exist")
	ErrMapId       = errors.New("invalid map id")
)

type mapServiceServer struct {
	pb.UnimplementedMapServiceServer
	Context *GameServerContext
}

// CreateMap implements pb.MapServiceServer.
func (s *mapServiceServer) CreateMap(ctx context.Context, request *pb.CreateMapRequest) (*pb.Map, error) {
	err := s.Context.validateRole(ctx, RoleMapManage)
	if err != nil {
		return nil, err
	}

	m, err := s.Context.MapService.CreateMap(ctx, request.Name, request.MapPath)
	if err != nil {
		if errors.Is(err, game.ErrValidation) {
			return nil, status.Error(codes.InvalidArgument, err.Error())
		}

		return nil, status.Error(codes.Internal, ErrMapCreate.Error())
	}

	s.Context.MapBusWriter.Publish(ctx, mapbus.Message{
		Id:      m.Id.String(),
		Deleted: false,
	})

	return m.ToPb(), nil
}

// DeleteMap implements pb.MapServiceServer.
func (s *mapServiceServer) DeleteMap(ctx context.Context, request *commonpb.TargetId) (*emptypb.Empty, error) {
	err := s.Context.validateRole(ctx, RoleMapManage)
	if err != nil {
		return nil, err
	}

	id, err := uuid.Parse(request.Id)
	if err != nil {
		return nil, status.Error(codes.InvalidArgument, ErrMapId.Error())
	}

	m, err := s.Context.MapService.DeleteMap(ctx, &id)
	if err != nil {
		log.Logger.WithContext(ctx).Errorf("%v: %v", ErrMapDelete, err)
		return nil, status.Error(codes.Internal, ErrMapDelete.Error())
	}
	if m == nil {
		return nil, status.Error(codes.NotFound, ErrMapNotExist.Error())
	}

	s.Context.MapBusWriter.Publish(ctx, mapbus.Message{
		Id:      m.Id.String(),
		Deleted: true,
	})

	return &emptypb.Empty{}, nil
}

// EditMap implements pb.MapServiceServer.
func (s *mapServiceServer) EditMap(ctx context.Context, request *pb.EditMapRequest) (*pb.Map, error) {
	err := s.Context.validateRole(ctx, RoleMapManage)
	if err != nil {
		return nil, err
	}

	m, err := s.getMapById(ctx, request.TargetId)
	if err != nil {
		return nil, err
	}

	if request.OptionalName != nil {
		m.Name = request.GetName()
	}
	if request.OptionalMapPath != nil {
		m.MapPath = request.GetMapPath()
	}

	m, err = s.Context.MapService.EditMap(ctx, m)
	if err != nil {
		if errors.Is(err, game.ErrValidation) {
			return nil, status.Error(codes.InvalidArgument, err.Error())
		}

		log.Logger.WithContext(ctx).Errorf("%v: %v", ErrMapEdit, err)
		return nil, status.Error(codes.Internal, ErrMapEdit.Error())
	}

	return m.ToPb(), nil
}

// GetMap implements pb.MapServiceServer.
func (s *mapServiceServer) GetMap(ctx context.Context, request *commonpb.TargetId) (*pb.Map, error) {
	err := s.Context.validateRole(ctx, RoleMapManage)
	if err != nil {
		return nil, err
	}

	m, err := s.getMapById(ctx, request.Id)
	if err != nil {
		return nil, err
	}

	return m.ToPb(), nil
}

// GetMaps implements pb.MapServiceServer.
func (s *mapServiceServer) GetMaps(ctx context.Context, request *emptypb.Empty) (*pb.Maps, error) {
	err := s.Context.validateRole(ctx, RoleMapManage)
	if err != nil {
		return nil, err
	}

	maps, err := s.Context.MapService.GetMaps(ctx)
	if err != nil {
		log.Logger.WithContext(ctx).Errorf("%v: %v", ErrMapLookup, err)
		return nil, status.Error(codes.Internal, ErrMapLookup.Error())
	}

	return maps.ToPb(), nil
}

func NewMapServiceServer(ctx context.Context, srvCtx *GameServerContext) (pb.MapServiceServer, error) {
	err := srvCtx.CreateRoles(ctx, &MapRoles)
	if err != nil {
		return nil, err
	}
	return &mapServiceServer{Context: srvCtx}, nil
}

func (s *mapServiceServer) getMapById(ctx context.Context, mapId string) (*game.Map, error) {
	id, err := uuid.Parse(mapId)
	if err != nil {
		return nil, status.Error(codes.InvalidArgument, ErrMapId.Error())
	}

	m, err := s.Context.MapService.GetMapById(ctx, &id)
	if err != nil {
		log.Logger.WithContext(ctx).Errorf("%v: %v", ErrMapLookup, err)
		return nil, status.Error(codes.Internal, ErrMapLookup.Error())
	}
	if m == nil {
		return nil, status.Error(codes.NotFound, ErrMapNotExist.Error())
	}

	return m, nil
}
