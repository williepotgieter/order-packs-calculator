package ports

import "github.com/williepotgieter/order-packs-calculator/internal/core/entities"

type Calculator interface {
	GetOrder(packs ...uint) (entities.Order, error)
}
