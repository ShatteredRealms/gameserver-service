package gameserver

import (
	"time"

	"github.com/google/uuid"
)

type PendingConnection struct {
	// Id secret used by a server to lookup a pending connection
	Id uuid.UUID `db:"id" json:"id"`

	// Character id of the character attempting to connect
	CharacterId uuid.UUID `db:"character_id" json:"characterId"`

	// ServerName the name of the server the character is assigned to
	ServerName string `db:"server_name" json:"serverName"`

	// CreatedAt when the pending connection was created
	CreatedAt time.Time `db:"created_at" json:"createdAt"`
}
