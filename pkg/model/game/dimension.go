package game

import (
	"fmt"
	"regexp"

	"github.com/ShatteredRealms/gameserver-service/pkg/pb"
	"github.com/ShatteredRealms/go-common-service/pkg/model"
)

const (
	MinDimensionNameLength = 1
	MaxDimensionNameLength = 64
)

var (
	DimensionNameRegex = "^[a-zA-Z0-9_-]+$"

	// ErrDimensionNameToShort thrown when a dimension name is too short
	ErrDimensionNameToShort = fmt.Errorf("%w: name must be at least %d characters", ErrValidation, MinDimensionNameLength)

	// ErrDimensionNameToLong thrown when a dimension name is too long
	ErrDimensionNameToLong = fmt.Errorf("%w: name can be at most %d characters", ErrValidation, MaxDimensionNameLength)

	// ErrNameToLong thrown when a dimension name is too long
	ErrDimensionRegex = fmt.Errorf("%w: name can be alphanumeric with spaces, dashes and underscores", ErrValidation)
)

type Dimension struct {
	model.Model
	Name     string `gorm:"index:idx_deleted,unique;not null" json:"name"`
	Version  string `gorm:"not null" json:"version"`
	Maps     Maps   `gorm:"many2many:dimension_maps" json:"maps"`
	Location string `gorm:"not null" json:"location"`
}
type Dimensions []*Dimension

func (dimension *Dimension) Validate() error {
	return dimension.ValidateName()
}

func (dimension *Dimension) ValidateName() error {
	if len(dimension.Name) < MinDimensionNameLength {
		return ErrDimensionNameToShort
	}

	if len(dimension.Name) > MaxDimensionNameLength {
		return ErrDimensionNameToLong
	}

	ok, err := regexp.MatchString(DimensionNameRegex, dimension.Name)
	if !ok {
		return ErrDimensionRegex
	}
	if err != nil {
		return fmt.Errorf("%w: %w", ErrRegex, err)
	}

	return nil
}

func (dimension *Dimension) ToPb() *pb.Dimension {
	maps := make([]string, len(dimension.Maps))
	for _, m := range dimension.Maps {
		maps = append(maps, m.Id.String())
	}

	return &pb.Dimension{
		Id:       dimension.Id.String(),
		Name:     dimension.Name,
		Version:  dimension.Version,
		MapIds:   maps,
		Location: dimension.Location,
	}
}

func (dimensions Dimensions) ToPb() *pb.Dimensions {
	out := make([]*pb.Dimension, len(dimensions))
	for idx, c := range dimensions {
		out[idx] = c.ToPb()
	}

	return &pb.Dimensions{
		Dimensions: out,
	}
}

func (dimension *Dimension) GetVersionOrDefault() string {
	version := "latest"
	if dimension.Version != "" {
		version = dimension.Version
	}

	return version
}
