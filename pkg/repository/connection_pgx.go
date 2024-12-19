package repository

import (
	"context"

	"github.com/ShatteredRealms/gameserver-service/pkg/model/gameserver"
	"github.com/ShatteredRealms/go-common-service/pkg/repository"
	"github.com/google/uuid"
	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
)

type pgxConnectionRepository struct {
	conn *pgxpool.Pool
}

func NewPgxConnectionRepository(migrater *repository.PgxMigrater) ConnectionRepository {
	return &pgxConnectionRepository{
		conn: migrater.Conn,
	}
}

// CreatePendingConnection implements ConnectionRepository.
func (p *pgxConnectionRepository) CreatePendingConnection(ctx context.Context, character string, serverName string) (*gameserver.PendingConnection, error) {
	tx, err := p.conn.Begin(ctx)
	defer tx.Rollback(ctx)
	if err != nil {
		return nil, err
	}

	rows, err := tx.Query(ctx,
		"INSERT INTO pending_connections (character, server_name) VALUES ($1, $2) RETURNING *",
		character, serverName)
	if err != nil {
		return nil, err
	}

	outConn, err := pgx.CollectExactlyOneRow(rows, pgx.RowToStructByName[gameserver.PendingConnection])
	if err != nil {
		return nil, err
	}

	err = tx.Commit(ctx)
	if err != nil {
		return nil, err
	}

	return &outConn, nil
}

// DeletePendingConnection implements ConnectionRepository.
func (p *pgxConnectionRepository) DeletePendingConnection(ctx context.Context, id *uuid.UUID) error {
	tx, err := p.conn.Begin(ctx)
	defer tx.Rollback(ctx)
	if err != nil {
		return err
	}

	ct, err := tx.Exec(ctx,
		"DELETE FROM pending_connections WHERE id = $1",
		id)
	if err != nil {
		return err
	}

	if ct.RowsAffected() == 0 {
		return pgx.ErrNoRows
	}

	return tx.Commit(ctx)
}

// FindPendingConnection implements ConnectionRepository.
func (p *pgxConnectionRepository) FindPendingConnection(ctx context.Context, id *uuid.UUID) (*gameserver.PendingConnection, error) {
	tx, err := p.conn.Begin(ctx)
	defer tx.Rollback(ctx)
	if err != nil {
		return nil, err
	}

	rows, err := tx.Query(ctx,
		"SELECT * FROM pending_connections WHERE id = $1",
		id)
	if err != nil {
		return nil, err
	}
	out, err := pgx.CollectExactlyOneRow(rows, pgx.RowToStructByName[gameserver.PendingConnection])
	if err != nil {
		return nil, err
	}
	err = tx.Commit(ctx)
	if err != nil {
		return nil, err
	}
	return &out, nil

}
