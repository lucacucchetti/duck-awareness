package entities

type ResourceID = string

type Resource interface {
	id() ResourceID
}