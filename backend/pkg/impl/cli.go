package impl

import (
	"bufio"
	"fmt"
	"io"
	"strconv"
	"strings"

	"github.com/runattekky/footyrate/models"
	"github.com/runattekky/footyrate/pkg"
)

type CLI struct {
	in   *bufio.Scanner
	out  io.Writer
	game pkg.Game
}

func NewCLI(in io.Reader, out io.Writer, game pkg.Game) *CLI {
	return &CLI{
		in:   bufio.NewScanner(in),
		out:  out,
		game: game,
	}
}

func (c *CLI) Start() {
	fmt.Fprintln(c.out, "Let us start comparing")
	for true {
		fmt.Fprintln(c.out, "Getting two new players")
		p1, p2, err := c.game.GetTwoOpps()
		if err != nil {
			fmt.Fprintln(c.out, "Error getting two opponents")
		}

		c.displayPlayers(p1, p2)

		fmt.Fprintln(c.out, "Choose the better player, input 1 or 2")
		choiceInput := c.readline()

		choice, err := strconv.Atoi(strings.Trim(choiceInput, "\n"))
		if err != nil {
			fmt.Fprintln(c.out, "Bad input")
		}

		switch choice {
		case 1:
			c.game.SaveChoice(p1, p2)
		case 2:
			c.game.SaveChoice(p2, p1)
		default:
			fmt.Fprintln(c.out, "Choose correct option")
		}

		fmt.Fprintln(c.out, "Players after vote")
		c.displayPlayers(p1, p2)
		fmt.Fprint(c.out, "\n\n")
	}
}

func (c *CLI) readline() string {
	c.in.Scan()
	return c.in.Text()
}

func (c *CLI) displayPlayers(p1, p2 *models.Footballer) {
	fmt.Fprint(c.out, "Player 1: ")
	fmt.Fprintln(c.out, p1)
	fmt.Fprint(c.out, "Player 2: ")
	fmt.Fprintln(c.out, p2)
}
