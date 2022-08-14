package entities

import "encoding/json"

// A board is a collection of resources, while being a resource itself
type Board struct {
	id        ResourceID
	name      string
	resources []*Resource
}

func (b Board) Id() ResourceID {
	return b.id
}

func (b *Board) GetResources() []*Resource {
	return b.resources
}

func (b *Board) GetName() string {
	return b.name
}

func (b Board) MarshalJSON() ([]byte, error) {
	resourcesIds := []string{}
	for _, resPtr := range b.resources {
		res := *resPtr
		resourcesIds = append(resourcesIds, res.Id())
	}
	return json.Marshal(&struct {
		ID        string   `json:"id"`
		Name      string   `json:"name"`
		Resources []string `json:"resources"`
	}{
		ID:        b.id,
		Name:      b.name,
		Resources: resourcesIds,
	})
}

func CreateBoard(id ResourceID, resources []*Resource, name string) Board {
	return Board{id: id, resources: resources, name: name}
}
