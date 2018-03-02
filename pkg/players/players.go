//Package for handling player data

package players

//Player is a structure that includes a name and a slice of struct seasons
type Player struct {
	Name    string
	Seasons []Season
}

//Season describes a full season worth of games, and a year the season took place durring
type Season struct {
	seasonYear int
	Games      []Game
}

//Game describes a game for a player: what game number of the season it was, what opponent they were playing, and how many rbis they hit
type Game struct {
	gameNumber int
	opponent   string
	rbi        int
}

//AppendSeasons appends a number of seasons to a player
func (p *Player) AppendSeasons(s ...Season) {
	//What happens if no seasons passed in?
	//Should error if a season is repeated
	//Should we be able to append games to an already created season?
	for _, season := range s {
		p.Seasons = append(p.Seasons, season)
	}
}

//AppendGames appends a number of games to a season
func (s *Season) AppendGames(g ...Game) {
	//What happens if no games passed in?
	//Should also error if we repeat a game number
	//Should we be able to append games to an already created season?
	for _, game := range g {
		s.Games = append(s.Games, game)
	}
}
