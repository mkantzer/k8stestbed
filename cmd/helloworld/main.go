// starting out with a small program that just prints
// to the command line, as proof of repo structure

//I'm pretty sure I've broken this all, because of a missunderstanding of how structs and packages works.
// I may need to work on a simpler data project to figure this stuff out

package main

import (
	"fmt"

	spew "github.com/davecgh/go-spew/spew"

	players "github.com/mkantzer/k8stestbed/pkg/players"
)

func main() {
	cal := players.Player{Name: "Cal Ripken Jr"}
	fmt.Println(cal.Name)
	spew.Dump(cal)
	g1 := players.Game{gameNumber: 1, opponent: "Yankies", rbi: 5}
	g2 := players.Game{gameNumber: 2, opponent: "Seattle", rbi: 1}
	s := players.Season{season}
}
