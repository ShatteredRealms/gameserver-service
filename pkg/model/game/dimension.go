package game

import (
	"fmt"
	"regexp"
	"time"

	"github.com/ShatteredRealms/gameserver-service/pkg/pb"
	"github.com/google/uuid"
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
	Id        uuid.UUID  `db:"id" json:"id"`
	Name      string     `db:"name" json:"name"`
	Version   string     `db:"version" json:"version"`
	Location  string     `db:"location" json:"location"`
	CreatedAt time.Time  `db:"created_at" json:"createdAt"`
	UpdatedAt time.Time  `db:"updated_at" json:"updatedAt"`
	DeletedAt *time.Time `db:"deleted_at" json:"deletedAt"`

	Maps Maps `db:"maps" json:"maps"`
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
	for idx, m := range dimension.Maps {
		maps[idx] = m.Id.String()
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
