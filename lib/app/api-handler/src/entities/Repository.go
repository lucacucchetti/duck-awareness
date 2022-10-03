package entities

type Repository interface {
	GetTextNote(id ResourceID) (TextNote, error)
	GetBoard(id ResourceID) (Board, error)
}
