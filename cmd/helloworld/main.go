// starting out with a small program that just prints
// to the command line, as proof of repo structure

package main

import (
	"fmt"

	players "github.com/mkantzer/k8stestbed/pkg/players"
)

func main() {
	cal := players.Player{Name: "Cal Ripken Jr"}
	fmt.Println(cal.Name)
}
