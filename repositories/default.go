package repositories

import (
	"context"

	i "github.com/pergamenum/go-consensus-standards/interfaces"
	"github.com/pergamenum/go-consensus-standards/reflection"
	t "github.com/pergamenum/go-consensus-standards/types"
)

type Repo[M any, E any] struct {
	dao i.DAO[E]
}

type RepoConfig[M any, E any] struct {
	DAO i.DAO[E]
}

func NewRepo[M any, E any](conf RepoConfig[M, E]) *Repo[M, E] {

	return &Repo[M, E]{
		dao: conf.DAO,
	}
}

func (r *Repo[M, E]) Create(ctx context.Context, id string, model M) error {

	entity, err := reflection.AutoMap[E](model)
	if err != nil {
		return err
	}

	err = r.dao.Create(ctx, id, entity)
	if err != nil {
		return err
	}

	return nil
}

func (r *Repo[M, E]) Read(ctx context.Context, id string) (model M, err error) {

	entity, err := r.dao.Read(ctx, id)
	if err != nil {
		return model, err
	}

	model, err = reflection.AutoMap[M](entity)
	if err != nil {
		return model, err
	}

	return model, nil
}

func (r *Repo[M, E]) Update(ctx context.Context, id string, update t.Update) error {

	if len(update) == 0 {
		return nil
	}

	err := r.dao.Update(ctx, id, update)
	if err != nil {
		return err
	}

	return nil
}

func (r *Repo[M, E]) Delete(ctx context.Context, id string) error {

	err := r.dao.Delete(ctx, id)
	if err != nil {
		return err
	}

	return nil
}

func (r *Repo[M, E]) Search(ctx context.Context, query []t.Query) ([]M, error) {

	es, err := r.dao.Search(ctx, query)
	if err != nil {
		return nil, err
	}

	var ms []M
	for _, e := range es {
		m, err := reflection.AutoMap[M](e)
		if err != nil {
			return nil, err
		}
		ms = append(ms, m)
	}

	return ms, nil
}
