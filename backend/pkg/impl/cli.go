package impl

import (
	"bufio"
	"fmt"
	"io"
	"strconv"
	"strings"

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
	fmt.Fprintf(c.out, "Let us start comparing")
	c.game.GetTwoOpps()

	fmt.Fprintf(c.out, "Choose the better player, input 1 or 2")
	choiceInput := c.readline()

	_, err := strconv.Atoi(strings.Trim(choiceInput, "\n"))
	if err != nil {
		fmt.Fprintf(c.out, "Bad input")
	}
}

func (c *CLI) readline() string {
	c.in.Scan()
	return c.in.Text()
}
