package interfaces

import t "github.com/pergamenum/go-consensus-standards/types"

type Controller[C any] interface {
	Create(ctx C)
	Read(ctx C)
	Update(ctx C)
	Delete(ctx C)
	Search(ctx C)
}

type ControllerMapper[Model, DTO any] interface {
	ToDTO(Model) DTO
	FromDTO(DTO) Model
	ToUpdate(DTO) t.Update
}
