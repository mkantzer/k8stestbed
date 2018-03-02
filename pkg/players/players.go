//Here will be all of the logic for describing and manipulating
//player information

package players

//Player is a structure that includes a name and a slice of struct Seasons
type Player struct {
	name    string
	seasons []struct {
		seasonYear int
		games      []struct {
			gameNumber int
			opponent   string
			rbi        int
			//maybe will need to have a date in here for the graphing?
		}
	}
}

func AppendSeason