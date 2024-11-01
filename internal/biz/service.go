package biz

import (
	"context"
)

// Entity is a entity model.
type Entity struct {
	Hello string
}

// Create .
func (uc *Service) Create(ctx context.Context, g *Entity) (*Entity, error) {
	uc.log.WithContext(ctx).Infof("Create: %v", g.Hello)
	return uc.repo.Save(ctx, g)
}
