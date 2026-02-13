package category

import "context"

type Repository interface {
	Create(ctx context.Context, c *Category) error
	FindAll(ctx context.Context) ([]*Category, error)
	FindUsed(ctx context.Context) ([]*Category, error)
	ExistsByName(ctx context.Context, name string) (bool, error)
}
