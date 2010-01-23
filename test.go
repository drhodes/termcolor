package main

import . "./termcolor"
import . "fmt"

func main(){
	Printf(Colorize("text\n", "red"))
	Printf(Colorize("text\n", "blue"))
	Printf(Colorize("text\n", "black"))
	Printf(Colorize("text\n", "magenta"))
	Printf(Colorize("text\n", "white"))
	Printf(Colorize("text\n", "yellow"))
	Printf(Colorize("text\n", "Marmalade Fuscia")) // default test.
}
