package models

type Player interface {
	ProcessWin(opponent Player)
	ProcessLoss(opponent Player)
}
