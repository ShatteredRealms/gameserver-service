package game

import (
	"time"

	"github.com/ShatteredRealms/gameserver-service/pkg/pb"
	"github.com/ShatteredRealms/go-common-service/pkg/model"
)

type Map struct {
	model.Model
	Name      string `gorm:"not null" json:"name"`
	MapPath   string `gorm:"not null" json:"map_path"`
	CreatedAt time.Time
	UpdatedAt time.Time
}

type Maps []*Map

func (m *Map) ToPb() *pb.Map {
	return &pb.Map{
		Id:      m.Id.String(),
		Name:    m.Name,
		MapPath: m.MapPath,
	}
}

func (maps Maps) ToPb() *pb.Maps {
	out := make([]*pb.Map, len(maps))
	for idx, m := range maps {
		out[idx] = m.ToPb()
	}

	return &pb.Maps{
		Maps: out,
	}
}
