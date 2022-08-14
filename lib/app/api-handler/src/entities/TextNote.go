package entities

import "encoding/json"

type TextNote struct {
	id      ResourceID
	title   string
	content string
}

func (n TextNote) Id() ResourceID {
	return n.id
}

func (n *TextNote) GetTitle() string {
	return n.title
}

func (n *TextNote) GetContent() string {
	return n.content
}

func (n TextNote) MarshalJSON() ([]byte, error) {
	return json.Marshal(&struct {
		ID      string `json:"id"`
		Title   string `json:"title"`
		Content string `json:"content"`
	}{
		ID:      n.id,
		Title:   n.title,
		Content: n.content,
	})
}

func CreateTextNote(id ResourceID, title, content string) TextNote {
	return TextNote{id: id, title: title, content: content}
}
