package gameserver

import (
	"time"

	"github.com/google/uuid"
)

type PendingConnection struct {
	// Id secret used by a server to lookup a pending connection
	Id *uuid.UUID `gorm:"type:uuid;default:gen_random_uuid()"`

	// Character id of the character attempting to connect
	CharacterId string `gorm:"not null;uniqueIndex"`

	// ServerName the name of the server the character is assigned to
	ServerName string `gorm:"not null"`

	// CreatedAt when the pending connection was created
	CreatedAt time.Time
}
