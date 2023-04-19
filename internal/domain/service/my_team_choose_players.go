package service

import (
	"errors"
	"github.com/fabioalmeida132/go-consolidador/internal/domain/entity"
)

func ChoosePlayers(myTeam entity.MyTeam, players []entity.Player) error {
	totalCost := 0.0
	totalEarned := 0.0

	for _, player := range players {
		if playerInMyTeam(player, myTeam) && !playerInPlayers(player, &players) {
			totalEarned += player.Price
		}
		if !playerInMyTeam(player, myTeam) && playerInPlayers(player, &players) {
			totalCost += player.Price
		}
	}

	if totalCost > totalEarned {
		return errors.New("total cost is greater than total earned")
	}

	myTeam.Score += totalEarned - totalCost
	myTeam.Players = []string{}

	for _, player := range players {
		myTeam.Players = append(myTeam.Players, player.ID)
	}
	return nil
}

func playerInMyTeam(player entity.Player, myTeam entity.MyTeam) bool {
	for _, playerID := range myTeam.Players {
		if playerID == player.ID {
			return true
		}
	}
	return false
}

func playerInPlayers(player entity.Player, players *[]entity.Player) bool {
	for _, p := range *players {
		if p.ID == player.ID {
			return true
		}
	}
	return false
}


// 55:44 - INTERFACE.GO