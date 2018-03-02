package main

import (
	spew "github.com/davecgh/go-spew/spew"
)

type Player struct {
	Name    string
	seasons []season
}

type season struct {
	seasonYear int
	games      []game
}

type game struct {
	gameNumber int
	opponent   string
	rbi        int
}

func (p *Player) appendSeasons(s ...season) {
	//What happens if no seasons passed in?
	//Should error if a season is repeated
	//Should we be able to append games to an already created season?
	for _, season := range s {
		p.seasons = append(p.seasons, season)
	}
}

func (s *season) appendGames(g ...game) {
	//What happens if no games passed in?
	//Should also error if we repeat a game number
	//Should we be able to append games to an already created season?
	for _, game := range g {
		s.games = append(s.games, game)
	}
}

func main() {
	cal := Player{Name: "Cal"}
	//fmt.Println(cal.Name)
	s := season{seasonYear: 2016}
	g1 := game{gameNumber: 1, opponent: "Yankies", rbi: 5}
	g2 := game{gameNumber: 2, opponent: "Seattle", rbi: 1}
	s.appendGames(g1)
	s.appendGames(g2)
	cal.appendSeasons(s)
	//spew.Dump(s)
	spew.Dump(cal)
}
