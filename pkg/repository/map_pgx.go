package repository

import (
	"context"

	"github.com/ShatteredRealms/gameserver-service/pkg/model/game"
	"github.com/ShatteredRealms/go-common-service/pkg/repository"
	"github.com/google/uuid"
	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
)

type PgxMapRepository struct {
	conn *pgxpool.Pool
}

func NewPgxMapRepository(migrater *repository.PgxMigrater) MapRepository {
	return &PgxMapRepository{
		conn: migrater.Conn,
	}
}

// CreateMap implements MapRepository.
func (p *PgxMapRepository) CreateMap(ctx context.Context, m *game.Map) (*game.Map, error) {
	tx, err := p.conn.Begin(ctx)
	defer tx.Rollback(ctx)
	if err != nil {
		return nil, err
	}

	rows, _ := tx.Query(ctx,
		"INSERT INTO maps (name, map_path) VALUES ($1, $2) RETURNING *",
		m.Name, m.MapPath)
	outMap, err := pgx.CollectExactlyOneRow(rows, pgx.RowToAddrOfStructByName[game.Map])
	if err != nil {
		return nil, err
	}

	return outMap, tx.Commit(ctx)
}

// DeleteMap implements MapRepository.
func (p *PgxMapRepository) DeleteMap(ctx context.Context, mapId *uuid.UUID) (*game.Map, error) {
	tx, err := p.conn.Begin(ctx)
	defer tx.Rollback(ctx)
	if err != nil {
		return nil, err
	}

	rows, _ := tx.Query(ctx,
		"UPDATE maps SET deleted_at = now() WHERE id = $1 RETURNING *",
		mapId)
	outMap, err := pgx.CollectExactlyOneRow(rows, pgx.RowToAddrOfStructByName[game.Map])
	if err != nil {
		return nil, err
	}

	return outMap, tx.Commit(ctx)
}

// GetDeletedMaps implements MapRepository.
func (p *PgxMapRepository) GetDeletedMaps(ctx context.Context) (game.Maps, error) {
	tx, err := p.conn.Begin(ctx)
	defer tx.Rollback(ctx)
	if err != nil {
		return nil, err
	}

	rows, _ := tx.Query(ctx,
		"SELECT * FROM maps WHERE deleted_at IS NOT NULL")
	outMaps, err := pgx.CollectRows(rows, pgx.RowToAddrOfStructByName[game.Map])
	if err != nil {
		return nil, err
	}

	return outMaps, tx.Commit(ctx)
}

// GetMapById implements MapRepository.
func (p *PgxMapRepository) GetMapById(ctx context.Context, mapId *uuid.UUID) (*game.Map, error) {
	tx, err := p.conn.Begin(ctx)
	defer tx.Rollback(ctx)
	if err != nil {
		return nil, err
	}

	rows, _ := tx.Query(ctx,
		"SELECT * FROM maps WHERE id = $1",
		mapId)
	outMap, err := pgx.CollectExactlyOneRow(rows, pgx.RowToAddrOfStructByName[game.Map])
	if err != nil {
		return nil, err
	}

	return outMap, tx.Commit(ctx)
}

// GetMaps implements MapRepository.
func (p *PgxMapRepository) GetMaps(ctx context.Context) (game.Maps, error) {
	tx, err := p.conn.Begin(ctx)
	defer tx.Rollback(ctx)
	if err != nil {
		return nil, err
	}

	rows, _ := tx.Query(ctx,
		"SELECT * FROM maps")
	outMaps, err := pgx.CollectRows(rows, pgx.RowToAddrOfStructByName[game.Map])
	if err != nil {
		return nil, err
	}

	return outMaps, tx.Commit(ctx)
}

// UpdateMap implements MapRepository.
func (p *PgxMapRepository) UpdateMap(ctx context.Context, m *game.Map) (*game.Map, error) {
	tx, err := p.conn.Begin(ctx)
	defer tx.Rollback(ctx)
	if err != nil {
		return nil, err
	}

	rows, _ := tx.Query(ctx,
		"UPDATE maps SET name = $1, map_path = $2 WHERE id = $3 RETURNING *",
		m.Name, m.MapPath, m.Id)
	outMap, err := pgx.CollectExactlyOneRow(rows, pgx.RowToAddrOfStructByName[game.Map])
	if err != nil {
		return nil, err
	}

	return outMap, tx.Commit(ctx)
}
