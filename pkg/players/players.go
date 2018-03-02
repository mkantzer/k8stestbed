//Package for handling player data

package players

//Player is a structure that includes a name and a slice of struct Seasons
type Player struct {
	Name string
	season
}

type season struct {
	seasonYear int
	games      []struct {
		gameNumber int
		opponent   string
		rbi        int
	}
}

func (p *Player) appendSeason(s season) {

}
