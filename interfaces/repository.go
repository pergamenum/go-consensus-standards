package interfaces

import (
	"context"

	t "github.com/pergamenum/go-consensus-standards/types"
)

type Repository[Model any] interface {
	Create(ctx context.Context, id string, model Model) error
	Read(ctx context.Context, id string) (Model, error)
	Update(ctx context.Context, id string, update t.Update) error
	Delete(ctx context.Context, id string) error
	Search(ctx context.Context, query []t.Query) ([]Model, error)
}

// RepositoryMapper
//
//	DU = DAO Update
type RepositoryMapper[Model, Entity any] interface {
	ToEntity(Model) Entity
	FromEntity(Entity) Model
}
