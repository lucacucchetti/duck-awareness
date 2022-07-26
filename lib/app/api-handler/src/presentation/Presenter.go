package presentation

import (
	"example/hello/src/entities"
)

type DisplayType int8

const (
	Text DisplayType = iota
	HTML
)

type Presenter interface {
	DisplayBoard(board entities.Board) (DisplayType, string, error)
	DisplayWidget(board entities.Resource) (DisplayType, string, error)
}
