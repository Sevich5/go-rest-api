package converter

import (
	"app/internal/domain/entity"
	"app/internal/infrastructure/persistence/model"
)

type DomainModelConverter[D entity.Root, M model.Base] interface {
	FromDomainToModel(d *D) M
	FromModelToDomain(m *M) D
}
