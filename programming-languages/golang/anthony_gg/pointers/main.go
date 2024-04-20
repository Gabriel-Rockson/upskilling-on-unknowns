package main

import "fmt"

type Player struct {
	health int
}

func (player *Player) takeDamageFromExplosion2(damage int) {
	fmt.Println("Player is taking damage from explosion")
	player.health -= damage
}

func takeDamageFromExplosion(player *Player, damage int) {
	fmt.Println("Player is taking damage from explosion")
	player.health -= damage
}

func main() {
	player := &Player{
		health: 100,
	}

	fmt.Printf("Before explosion %+v\n", player.health)
	// takeDamageFromExplosion(&player)
	player.takeDamageFromExplosion2(20)
	fmt.Printf("After explosion %+v\n", player.health)
}
