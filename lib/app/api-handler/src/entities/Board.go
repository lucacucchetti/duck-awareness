package entities

// A board is a collection of resources, while being a resource itself
type Board struct {
	id        ResourceID
	resources []Resource
}

func (b Board) Id() ResourceID {
	return b.id
}

func (b *Board) GetResources() []Resource {
	return b.resources
}
