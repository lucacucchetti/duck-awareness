package presentation

import (
	"errors"
	"example/hello/src/entities"
	"reflect"
)

type TextPresenter struct {
}

// Returns the concatenation of the representation of each of the widgets contained in this board
func (presenter TextPresenter) DisplayBoard(board entities.Board) (DisplayType, string, error) {
	final_representation := "Board: " + board.GetName() + "\n[\n"
	for _, resource := range board.GetResources() {
		_, representation, error := presenter.DisplayWidget(*resource)
		if error != nil {
			return Text, "", error
		}
		final_representation += "\t" + representation + "\n"
	}
	final_representation += "]"
	return Text, final_representation, nil
}

// Returns the representation of a resource in it's widget format (when it is being shown inside a board)
func (presenter TextPresenter) DisplayWidget(resource entities.Resource) (DisplayType, string, error) {
	displayType := Text
	if board, ok := resource.(*entities.Board); ok {
		return displayType, "Board: " + board.GetName(), nil
	}
	return displayType, "", errors.New("Unexpected resource type " +
		reflect.TypeOf(resource).Name())
}
