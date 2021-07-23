package rps

import (
	"math/rand"
	"time"
)

const (
	ROCK         = 0 // beats SCISSORS. (SCISSORS + 1) % 3 = 0
	PAPER        = 1 // beats ROCK. (ROCK + 1) % 3 = 1
	SCISSORS     = 2 // beats PAPER. (PAPER + 1) % 3 = 2
	PLAYWINS     = 1
	COMPUTERWINS = 2
	DRAW         = 3
)

type Round struct {
	Winner         int    `json:"winner"`
	ComputerChoice string `json:"computer_choice"`
	Result         string `json:"round_result"`
}

func PlayRound(playerValue int) Round {
	rand.Seed(time.Now().UnixNano())
	computerValue := rand.Intn(3)
	computerChoice := "Computer chose "

	switch computerValue {
	case ROCK:
		computerChoice += "ROCK"
	case PAPER:
		computerChoice += "PAPER"
	case SCISSORS:
		computerChoice += "SCISSOR"
	}

	result := ""
	winner := 0
	if playerValue == computerValue {
		result = "It's a draw"
		winner = DRAW
	} else if playerValue == (computerValue+1)%3 {
		result = "Player wins!"
		winner = PLAYWINS
	} else {
		result = "Computer wins!"
		winner = COMPUTERWINS
	}

	var roundResult Round
	roundResult.Winner = winner
	roundResult.ComputerChoice = computerChoice
	roundResult.Result = result

	return roundResult
}
