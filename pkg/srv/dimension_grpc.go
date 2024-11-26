package srv

import (
	"context"
	"errors"
	"fmt"

	"github.com/ShatteredRealms/gameserver-service/pkg/model/game"
	"github.com/ShatteredRealms/gameserver-service/pkg/pb"
	"github.com/ShatteredRealms/go-common-service/pkg/bus"
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
	DimensionRoles = make([]*gocloak.Role, 0)

	RoleDimensionManage = util.RegisterRole(&gocloak.Role{
		Name:        gocloak.StringP("dimension.manage"),
		Description: gocloak.StringP("Manage dimensions"),
	}, &DimensionRoles)
)

var (
	ErrDimensionCreate   = errors.New("failed to create dimension")
	ErrDimensionDelete   = errors.New("failed to delete dimension")
	ErrDimensionEdit     = errors.New("failed to edit dimension")
	ErrDimensionLookup   = errors.New("failed to lookup dimension")
	ErrDimensionNotExist = errors.New("dimension does not exist")
	ErrDimensionId       = errors.New("invalid dimension id")
)

type dimensionServiceServer struct {
	pb.UnimplementedDimensionServiceServer
	Context *GameServerContext
}

func NewDimensionServiceServer(ctx context.Context, srvCtx *GameServerContext) (pb.DimensionServiceServer, error) {
	err := srvCtx.CreateRoles(ctx, &DimensionRoles)
	if err != nil {
		return nil, err
	}
	return &dimensionServiceServer{Context: srvCtx}, nil
}

// CreateDimension implements pb.DimensionServiceServer.
func (c *dimensionServiceServer) CreateDimension(ctx context.Context, request *pb.CreateDimensionRequest) (*pb.Dimension, error) {
	err := c.Context.validateRole(ctx, RoleDimensionManage)
	if err != nil {
		return nil, err
	}

	mapIds := make([]*uuid.UUID, len(request.MapIds))
	for idx, mapId := range request.MapIds {
		id, err := uuid.Parse(mapId)
		if err != nil {
			return nil, status.Error(codes.InvalidArgument, ErrMapId.Error())
		}
		mapIds[idx] = &id

		m, err := c.Context.MapService.GetMapById(ctx, &id)
		if err != nil {
			log.Logger.WithContext(ctx).Errorf("%v: %v", ErrMapLookup, err)
			return nil, status.Error(codes.Internal, fmt.Sprintf("map id %v: %s", mapId, ErrMapLookup))
		}
		if m == nil {
			return nil, status.Error(codes.InvalidArgument, fmt.Sprintf("map id %v: %s", mapId, ErrMapNotExist))
		}
	}

	dimension, err := c.Context.DimensionService.CreateDimension(ctx, request.Name, request.Version, request.Location, mapIds)
	if err != nil {
		if errors.Is(err, game.ErrValidation) {
			return nil, status.Error(codes.InvalidArgument, err.Error())
		}

		return nil, status.Error(codes.Internal, ErrDimensionCreate.Error())
	}

	c.Context.DimensionBusWriter.Publish(ctx, bus.DimensionMessage{
		Id:      dimension.Id.String(),
		Deleted: false,
	})

	return dimension.ToPb(), nil
}

// DeleteDimension implements pb.DimensionServiceServer.
func (c *dimensionServiceServer) DeleteDimension(ctx context.Context, request *commonpb.TargetId) (*emptypb.Empty, error) {
	err := c.Context.validateRole(ctx, RoleDimensionManage)
	if err != nil {
		return nil, err
	}

	id, err := uuid.Parse(request.Id)
	if err != nil {
		return nil, status.Error(codes.InvalidArgument, ErrDimensionId.Error())
	}

	dimension, err := c.Context.DimensionService.DeleteDimension(ctx, &id)
	if err != nil {
		log.Logger.WithContext(ctx).Errorf("%v: %v", ErrDimensionDelete, err)
		return nil, status.Error(codes.Internal, ErrDimensionDelete.Error())
	}
	if dimension == nil {
		return nil, status.Error(codes.NotFound, ErrDimensionNotExist.Error())
	}

	c.Context.DimensionBusWriter.Publish(ctx, bus.DimensionMessage{
		Id:      dimension.Id.String(),
		Deleted: true,
	})

	return &emptypb.Empty{}, nil
}

// DuplicateDimension implements pb.DimensionServiceServer.
func (c *dimensionServiceServer) DuplicateDimension(ctx context.Context, request *pb.DuplicateDimensionRequest) (*pb.Dimension, error) {
	err := c.Context.validateRole(ctx, RoleDimensionManage)
	if err != nil {
		return nil, err
	}

	dimension, err := c.getDimensionById(ctx, request.TargetId)
	if err != nil {
		return nil, err
	}

	dimension.Id = nil
	dimension.Name = request.Name

	mapIds := make([]*uuid.UUID, len(dimension.Maps))
	for idx, m := range dimension.Maps {
		mapIds[idx] = m.Id
	}

	newDimension, err := c.Context.DimensionService.CreateDimension(ctx, dimension.Name, dimension.Version, dimension.Location, mapIds)
	if err != nil {
		if errors.Is(err, game.ErrValidation) {
			return nil, status.Error(codes.InvalidArgument, err.Error())
		}

		log.Logger.WithContext(ctx).Errorf("%v: %v", ErrDimensionCreate, err)
		return nil, status.Error(codes.Internal, ErrDimensionCreate.Error())
	}

	c.Context.DimensionBusWriter.Publish(ctx, bus.DimensionMessage{
		Id:      newDimension.Id.String(),
		Deleted: false,
	})

	return newDimension.ToPb(), nil
}

// EditDimension implements pb.DimensionServiceServer.
func (c *dimensionServiceServer) EditDimension(ctx context.Context, request *pb.EditDimensionRequest) (*pb.Dimension, error) {
	err := c.Context.validateRole(ctx, RoleDimensionManage)
	if err != nil {
		return nil, err
	}

	dimension, err := c.getDimensionById(ctx, request.TargetId)
	if err != nil {
		return nil, err
	}

	if request.OptionalName != nil {
		dimension.Name = request.GetName()
	}
	if request.OptionalLocation != nil {
		dimension.Location = request.GetLocation()
	}
	if request.OptionalVersion != nil {
		dimension.Version = request.GetVersion()
	}
	if request.EditMaps {
		maps := make([]*game.Map, len(request.MapIds))
		for idx, mapId := range request.MapIds {
			id, err := uuid.Parse(mapId)
			if err != nil {
				return nil, status.Error(codes.InvalidArgument, ErrMapId.Error())
			}

			m, err := c.Context.MapService.GetMapById(ctx, &id)
			if err != nil {
				log.Logger.WithContext(ctx).Errorf("%v: %v", ErrMapLookup, err)
				return nil, status.Error(codes.Internal, fmt.Sprintf("map id %v: %s", mapId, ErrMapLookup))
			}
			if m == nil {
				return nil, status.Error(codes.InvalidArgument, fmt.Sprintf("map id %v: %s", mapId, ErrMapNotExist))
			}
			maps[idx] = m
		}
		dimension.Maps = maps
	}

	dimension, err = c.Context.DimensionService.EditDimension(ctx, dimension)
	if err != nil {
		if errors.Is(err, game.ErrValidation) {
			return nil, status.Error(codes.InvalidArgument, err.Error())
		}

		log.Logger.WithContext(ctx).Errorf("%v: %v", ErrDimensionEdit, err)
		return nil, status.Error(codes.Internal, ErrDimensionEdit.Error())
	}

	return dimension.ToPb(), nil
}

// GetDimension implements pb.DimensionServiceServer.
func (c *dimensionServiceServer) GetDimension(ctx context.Context, request *commonpb.TargetId) (*pb.Dimension, error) {
	err := c.Context.validateRole(ctx, RoleDimensionManage)
	if err != nil {
		return nil, err
	}

	dimension, err := c.getDimensionById(ctx, request.Id)
	if err != nil {
		return nil, err
	}

	return dimension.ToPb(), nil
}

// GetDimensions implements pb.DimensionServiceServer.
func (c *dimensionServiceServer) GetDimensions(ctx context.Context, request *emptypb.Empty) (*pb.Dimensions, error) {
	err := c.Context.validateRole(ctx, RoleDimensionManage)
	if err != nil {
		return nil, err
	}

	dimensions, err := c.Context.DimensionService.GetDimensions(ctx)
	if err != nil {
		log.Logger.WithContext(ctx).Errorf("%v: %v", ErrDimensionLookup, err)
		return nil, status.Error(codes.Internal, ErrDimensionLookup.Error())
	}

	return dimensions.ToPb(), nil
}

func (c *dimensionServiceServer) getDimensionById(ctx context.Context, dimensionId string) (*game.Dimension, error) {
	id, err := uuid.Parse(dimensionId)
	if err != nil {
		return nil, status.Error(codes.InvalidArgument, ErrDimensionId.Error())
	}

	dimension, err := c.Context.DimensionService.GetDimensionById(ctx, &id)
	if err != nil {
		log.Logger.WithContext(ctx).Errorf("%v: %v", ErrDimensionLookup, err)
		return nil, status.Error(codes.Internal, ErrDimensionLookup.Error())
	}
	if dimension == nil {
		return nil, status.Error(codes.NotFound, ErrDimensionNotExist.Error())
	}

	return dimension, nil
}
