package domain

type Entity struct {
	Key   string
	Value string
}

func NewEntity(key string, value string) *Entity {
	return &Entity{Key: key, Value: value}
}

type EntityList struct {
	Entities []Entity
}

func NewEntityList(entities []Entity) *EntityList {
	return &EntityList{Entities: entities}
}
