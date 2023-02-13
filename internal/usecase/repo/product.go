package repo

import (
	"context"
	"fmt"

	"github.com/ilyasbulat/rest_api/internal/entity"
	"github.com/ilyasbulat/rest_api/pkg/postgres"
)

const (
	_defaultProductCap = 64
	schema             = "public"
	table              = "product"
)

// productRepo -.
type productRepo struct {
	*postgres.Postgres
}

// NewProductRepo -.
func NewProduct(pg *postgres.Postgres) *productRepo {
	return &productRepo{pg}
}

// GetAllProducts -.
func (r *productRepo) All(ctx context.Context) ([]entity.Product, error) {
	sql, _, err := r.Builder.
		Select(
			"id",
			"name",
			"description",
			"price",
			"currency_id",
			"rating",
			"category_id",
			"specification",
			"image",
			"created_at",
			"updated_at",
		).
		From("product").
		ToSql()
	if err != nil {
		return nil, fmt.Errorf("TranslationRepo - GetHistory - r.Builder: %w", err)
	}

	rows, err := r.Pool.Query(ctx, sql)
	if err != nil {
		return nil, fmt.Errorf("TranslationRepo - GetHistory - r.Pool.Query: %w", err)
	}
	defer rows.Close()

	products := make([]entity.Product, 0, _defaultProductCap)

	for rows.Next() {
		p := entity.Product{}

		err = rows.Scan(
			&p.ID,
			&p.Name,
			&p.Description,
			&p.Price,
			&p.CurrencyID,
			&p.Rating,
			&p.CategoryID,
			&p.Specification,
			&p.Image,
			&p.CreatedAt,
			&p.UpdatedAt,
		)
		if err != nil {
			return nil, fmt.Errorf("TranslationRepo - GetHistory - rows.Scan: %w", err)
		}

		products = append(products, p)
	}

	return products, nil
}

// Store -.
func (r *productRepo) Store(ctx context.Context, p entity.Product) error {
	// sql, args, err := r.Builder.
	// 	Insert("history").
	// 	Columns("source, destination, original, translation").
	// 	Values(t.Source, t.Destination, t.Original, t.Translation).
	// 	ToSql()
	// if err != nil {
	// 	return fmt.Errorf("TranslationRepo - Store - r.Builder: %w", err)
	// }

	// _, err = r.Pool.Exec(ctx, sql, args...)
	// if err != nil {
	// 	return fmt.Errorf("TranslationRepo - Store - r.Pool.Exec: %w", err)
	// }

	return nil
}
