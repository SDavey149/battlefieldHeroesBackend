package matchmaking

import (
	"github.com/Synaxis/battlefieldHeroesBackend/backend/network"
)

type Game struct {
	ID         int
	LobbyID    int
	GameServer *network.Client

	PlayersJoining int
	PlayersPlaying int
	PlayerSlots    int
	// VipSlots       int
}
