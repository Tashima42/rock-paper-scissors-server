package rockpaperscissors

import (
	"errors"
	"log"
	"math/rand"
)

var ErrMatchFull error = errors.New("match is full")
var ErrInvalidPlayerSize error = errors.New("match must have 2 players to start")
var ErrPlayerAlreadyMoved error = errors.New("player already made his move, wait for the other player")
var ErrMatchNotStarted error = errors.New("match didn't started yet")

type MoveType uint8
type MoveResult uint8
type PlayerNumber uint8

const (
	PlayerNumberOne PlayerNumber = 1
	PlayerNumberTwo PlayerNumber = 2
)

const (
	MoveTypeRock     MoveType = 0
	MoveTypePaper    MoveType = 1
	MoveTypeScissors MoveType = 2
)

const (
	MoveResultNotCompleted MoveResult = 0
	MoveResultWon          MoveResult = 1
	MoveResultLost         MoveResult = 2
	MoveResultDraw         MoveResult = 3
)

var moveResults map[MoveType]map[MoveType]MoveResult = map[MoveType]map[MoveType]MoveResult{
	MoveTypeRock: {
		MoveTypePaper:    MoveResultLost,
		MoveTypeRock:     MoveResultDraw,
		MoveTypeScissors: MoveResultWon,
	},
	MoveTypePaper: {
		MoveTypePaper:    MoveResultDraw,
		MoveTypeRock:     MoveResultWon,
		MoveTypeScissors: MoveResultLost,
	},
	MoveTypeScissors: {
		MoveTypePaper:    MoveResultWon,
		MoveTypeRock:     MoveResultLost,
		MoveTypeScissors: MoveResultDraw,
	},
}

type Game struct {
	matches []*Match
}

type Match struct {
	id            string
	playerOne     *Player
	playerTwo     *Player
	playerOneMove *MoveType
	playerTwoMove *MoveType
	playerIDs     map[string]PlayerNumber
	matchStarted  bool
	scores        map[string]int // "playerID": score
	maxScore      int
}

type Player struct {
	id   string
	name string
}

func NewGame() *Game {
	return new(Game)
}

func (g *Game) NewMatch(maxScore int) *Match {
	match := &Match{
		id:            pseudoRandomID(5),
		playerOne:     nil,
		playerTwo:     nil,
		playerOneMove: nil,
		playerTwoMove: nil,
		matchStarted:  false,
		playerIDs:     map[string]PlayerNumber{},
		scores:        map[string]int{},
		maxScore:      maxScore,
	}
	g.matches = append(g.matches, match)
	return match
}

func NewPlayer(name string) *Player {
	return &Player{
		id:   pseudoRandomID(5),
		name: name,
	}
}

func (m *Match) Join(player *Player) error {
	if m.playerOne == nil {
		log.Println("registering player one")
		m.playerOne = player
		m.playerIDs[player.id] = PlayerNumberOne
		return nil
	} else if m.playerTwo == nil {
		log.Println("registering player two")
		m.playerTwo = player
		m.playerIDs[player.id] = PlayerNumberTwo
		return nil
	}
	log.Println("all places are occupied")
	return ErrMatchFull
}

func (m *Match) Start() error {
	if m.playerOne == nil || m.playerTwo == nil {
		return ErrInvalidPlayerSize
	}
	m.matchStarted = true
	return nil
}

func (m *Match) Play(move MoveType, playerID string) (MoveResult, error) {
	log.Println("playing")
	if !m.matchStarted {
		log.Println("match not started")
		return MoveResultNotCompleted, ErrMatchNotStarted
	}
	log.Println("getting player by id: " + playerID)
	player, err := m.getPlayerNumberByID(playerID)
	if err != nil {
		return MoveResultNotCompleted, err
	}
	log.Printf("player number: %d\n", player)
	if player == PlayerNumberOne {
		if m.playerOneMove != nil {
			log.Println("player one already moved")
			return MoveResultNotCompleted, ErrPlayerAlreadyMoved
		}
		log.Printf("player one registering move: %d\n", move)
		m.playerOneMove = &move
	} else if player == PlayerNumberTwo {
		if m.playerTwoMove != nil {
			log.Println("player one already moved")
			return MoveResultNotCompleted, ErrPlayerAlreadyMoved
		}
		log.Printf("player two registering move: %d\n", move)
		m.playerTwoMove = &move
	}
	if m.playerOneMove != nil && m.playerTwoMove != nil {
		log.Println("both players already played, calculating results")
		result := moveResults[*m.playerOneMove][*m.playerTwoMove]
		log.Printf("result: %d", result)
		switch result {
		case MoveResultWon:
			log.Println("player one won")
			m.scores[m.playerOne.id] += 1
		case MoveResultLost:
			log.Println("player two won")
			m.scores[m.playerTwo.id] += 1
		}
		m.playerOneMove = nil
		m.playerTwoMove = nil
		return result, nil
	}
	return MoveResultNotCompleted, nil
}

func (m *Match) getPlayerNumberByID(id string) (PlayerNumber, error) {
	p, ok := m.playerIDs[id]
	if !ok {
		return 0, errors.New("invalid player id")
	}
	return p, nil
}

func (p *Player) GetID() string {
	return p.id
}

func pseudoRandomID(length int) string {
	letters := []rune("abcdefghijklmnopqrstuvwxyz123456789=*%#@!?")
	b := make([]rune, length)
	for i := range b {
		b[i] = letters[rand.Intn(len(letters))]
	}
	return string(b)
}
