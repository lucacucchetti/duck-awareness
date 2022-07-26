package entities

type ResourceID = string

type Resource interface {
	Id() ResourceID
}
