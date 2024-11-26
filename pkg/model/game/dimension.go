package game

import (
	"errors"
	"fmt"

	"github.com/ShatteredRealms/gameserver-service/pkg/pb"
	"github.com/ShatteredRealms/go-common-service/pkg/model"
)

const (
	MinNameLength = 1
	MaxNameLength = 64
)

var (
	// ErrValidation thrown when a validation error occurs
	ErrValidation = errors.New("validation")

	// ErrNameToShort thrown when a dimension name is too short
	ErrNameToShort = fmt.Errorf("%w: name must be at least %d characters", ErrValidation, MinNameLength)

	// ErrNameToLong thrown when a dimension name is too long
	ErrNameToLong = fmt.Errorf("%w: name can be at most %d characters", ErrValidation, MaxNameLength)
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
	if len(dimension.Name) < MinNameLength {
		return ErrNameToShort
	}

	if len(dimension.Name) > MaxNameLength {
		return ErrNameToLong
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
