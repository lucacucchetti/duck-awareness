package entities

// A board is a collection of resources, while being a resource itself
type Board struct {
	id        ResourceID
	name      string
	resources []*Resource
}

func NewBoard(id ResourceID, name string, resources []*Resource) *Board {
	return &Board{id: id, name: name, resources: resources}
}

func (b Board) Id() ResourceID {
	return b.id
}

func (b *Board) GetResources() []*Resource {
	return b.resources
}

func (b *Board) AddResource(resource Resource) {
	b.resources = append(b.resources, &resource)
}

func (b Board) GetName() string {
	return b.name
}
