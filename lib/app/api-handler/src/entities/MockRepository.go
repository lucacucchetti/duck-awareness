package entities

import (
	"math/rand"
)

type MockRepository struct {
}

func (_ MockRepository) GetBoard(id ResourceID) (Board, error) {
	return CreateBoard(id, []*Resource{}, "title: "+id), nil
}

func (_ MockRepository) GetTextNote(id ResourceID) (TextNote, error) {
	content := id
	reps := rand.Intn(20)
	for i := 0; i < reps; i++ {
		content += " " + id
	}
	return CreateTextNote(id, "title: "+id, content), nil
}
