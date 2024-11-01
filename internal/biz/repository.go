package biz

import "context"

// Repo .
type Repo interface {
	Save(context.Context, *Entity) (*Entity, error)
}
