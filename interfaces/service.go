package interfaces

import (
	"context"

	t "github.com/pergamenum/go-consensus-standards/types"
)

type Service[Model any] interface {
	Create(ctx context.Context, model Model) error
	Read(ctx context.Context, id string) (Model, error)
	Update(ctx context.Context, update t.Update) error
	Delete(ctx context.Context, id string) error
	Search(ctx context.Context, query []t.Query) ([]Model, error)
}
