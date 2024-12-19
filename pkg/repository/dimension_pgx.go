package repository

import (
	"context"
	"strings"

	"github.com/ShatteredRealms/gameserver-service/pkg/model/game"
	"github.com/ShatteredRealms/go-common-service/pkg/repository"
	"github.com/google/uuid"
	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
)

type pgxDimensionRepository struct {
	conn *pgxpool.Pool
}

func NewPgxDimensionRepository(migrater *repository.PgxMigrater) DimensionRepository {
	return &pgxDimensionRepository{
		conn: migrater.Conn,
	}
}

// CreateDimension implements DimensionRepository.
func (p *pgxDimensionRepository) CreateDimension(ctx context.Context, dimension *game.Dimension) (*game.Dimension, error) {
	tx, err := p.conn.Begin(ctx)
	defer tx.Rollback(ctx)
	if err != nil {
		return nil, err
	}

	rows, _ := tx.Query(ctx,
		"INSERT INTO dimensions (name, location, version) VALUES ($1, $2, $3) RETURNING *",
		dimension.Name, dimension.Location, dimension.Version)
	outDimension, err := pgx.CollectExactlyOneRow(rows, pgx.RowToAddrOfStructByName[game.Dimension])
	if err != nil {
		return nil, err
	}

	if len(dimension.Maps) > 0 {
		for _, m := range dimension.Maps {
			_, err := tx.Exec(ctx,
				"INSERT INTO dimension_maps (dimension_id, map_id) VALUES ($1, $2)",
				outDimension.Id, m.Id)
			if err != nil {
				return nil, err
			}
		}
		maps, err := p.findMaps(tx, ctx, dimension.Maps)
		if err != nil {
			return nil, err
		}

		outDimension.Maps = maps
	}

	return outDimension, tx.Commit(ctx)
}

// DeleteDimension implements DimensionRepository.
func (p *pgxDimensionRepository) DeleteDimension(ctx context.Context, dimensionId *uuid.UUID) (*game.Dimension, error) {
	tx, err := p.conn.Begin(ctx)
	defer tx.Rollback(ctx)
	if err != nil {
		return nil, err
	}

	rows, _ := tx.Query(ctx,
		"UPDATE dimensions SET DELETED_AT = now() WHERE id = $1 RETURNING *",
		dimensionId)
	outDimension, err := pgx.CollectExactlyOneRow(rows, pgx.RowToAddrOfStructByName[game.Dimension])
	if err != nil {
		return nil, err
	}

	return outDimension, tx.Commit(ctx)
}

// GetDeletedDimensions implements DimensionRepository.
func (p *pgxDimensionRepository) GetDeletedDimensions(ctx context.Context) (game.Dimensions, error) {
	tx, err := p.conn.Begin(ctx)
	defer tx.Rollback(ctx)
	if err != nil {
		return nil, err
	}

	rows, _ := tx.Query(ctx,
		"SELECT * FROM dimensions WHERE deleted_at IS NOT NULL")
	dimensions, err := pgx.CollectRows(rows, pgx.RowToAddrOfStructByName[game.Dimension])
	if err != nil {
		return nil, err
	}

	return dimensions, tx.Commit(ctx)
}

// GetDimensionById implements DimensionRepository.
func (p *pgxDimensionRepository) GetDimensionById(ctx context.Context, dimensionId *uuid.UUID) (*game.Dimension, error) {
	tx, err := p.conn.Begin(ctx)
	defer tx.Rollback(ctx)
	if err != nil {
		return nil, err
	}

	rows, _ := tx.Query(ctx, `
		SELECT d.*, 
			(select array_agg(row(m.id)) as maps from maps m where m.id in (select map_id from dimension_maps dm where dm.dimension_id = d.id))
		FROM dimensions d
		WHERE d.id = $1`, dimensionId)
	outDimension, err := pgx.CollectExactlyOneRow(rows, pgx.RowToAddrOfStructByName[game.Dimension])
	if err != nil {
		return nil, err
	}
	return outDimension, tx.Commit(ctx)
}

// GetDimensions implements DimensionRepository.
func (p *pgxDimensionRepository) GetDimensions(ctx context.Context) (game.Dimensions, error) {
	tx, err := p.conn.Begin(ctx)
	defer tx.Rollback(ctx)
	if err != nil {
		return nil, err
	}

	rows, _ := tx.Query(ctx, `
		SELECT d.*, 
			(select array_agg(row(m.id)) as maps from maps m where m.id in (select map_id from dimension_maps dm where dm.dimension_id = d.id))
		FROM dimensions d
		WHERE d.deleted_at IS NULL`)
	dimensions, err := pgx.CollectRows(rows, pgx.RowToAddrOfStructByName[game.Dimension])
	if err != nil {
		return nil, err
	}
	return dimensions, tx.Commit(ctx)
}

// UpdateDimension implements DimensionRepository.
func (p *pgxDimensionRepository) UpdateDimension(ctx context.Context, dimension *game.Dimension) (*game.Dimension, error) {
	tx, err := p.conn.Begin(ctx)
	defer tx.Rollback(ctx)
	if err != nil {
		return nil, err
	}

	rows, _ := tx.Query(ctx,
		"UPDATE dimensions SET name = $1, location = $2, version = $3 WHERE id = $4 RETURNING *",
		dimension.Name, dimension.Location, dimension.Version, dimension.Id)
	outDimension, err := pgx.CollectExactlyOneRow(rows, pgx.RowToAddrOfStructByName[game.Dimension])
	if err != nil {
		return nil, err
	}

	if len(dimension.Maps) > 0 {
		_, err := tx.Exec(ctx,
			"DELETE FROM dimension_maps WHERE dimension_id = $1",
			dimension.Id)
		if err != nil {
			return nil, err
		}

		for _, m := range dimension.Maps {
			_, err := tx.Exec(ctx,
				"INSERT INTO dimension_maps (dimension_id, map_id) VALUES ($1, $2)",
				outDimension.Id, m.Id)
			if err != nil {
				return nil, err
			}
		}
	}

	maps, err := p.findMaps(tx, ctx, dimension.Maps)
	if err != nil {
		return nil, err
	}

	outDimension.Maps = maps
	return outDimension, tx.Commit(ctx)
}

func (p *pgxDimensionRepository) findMaps(tx pgx.Tx, ctx context.Context, maps []*game.Map) (game.Maps, error) {
	queryBuilder := strings.Builder{}
	queryBuilder.WriteString("SELECT * FROM maps WHERE id IN (")

	paramBuilder := strings.Builder{}
	paramBuilder.WriteString("")
	for _, m := range maps {
		paramBuilder.WriteString("'")
		paramBuilder.WriteString(m.Id.String())
		paramBuilder.WriteString("', ")
	}
	paramStr := paramBuilder.String()
	queryBuilder.WriteString(paramStr[:len(paramStr)-2])
	queryBuilder.WriteString(")")

	rows, _ := tx.Query(ctx, queryBuilder.String())
	maps, err := pgx.CollectRows(rows, pgx.RowToAddrOfStructByName[game.Map])
	if err != nil {
		return nil, err
	}

	return maps, nil
}
