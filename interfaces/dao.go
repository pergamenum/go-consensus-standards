package interfaces

import (
	"context"

	t "github.com/pergamenum/go-consensus-standards/types"
)

type DAO[Entity any] interface {
	Create(ctx context.Context, id string, entity Entity) error
	Read(ctx context.Context, id string) (Entity, error)
	Update(ctx context.Context, id string, update t.Update) error
	Delete(ctx context.Context, id string) error
	Search(ctx context.Context, queries []t.Query) ([]Entity, error)
}
