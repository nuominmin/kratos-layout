package biz

import (
	"context"
)

// Repo .
type Repo interface {
	Save(context.Context, *Entity) (*Entity, error)
}

// Entity is a entity model.
type Entity struct {
	Hello string
}

// Create .
func (s *Service) Create(ctx context.Context, g *Entity) (*Entity, error) {
	s.log.WithContext(ctx).Infof("Create: %v", g.Hello)
	return s.repo.Save(ctx, g)
}
