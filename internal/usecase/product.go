package usecase

import (
	"context"
	"fmt"

	"github.com/ilyasbulat/rest_api/internal/entity"
)

// TranslationUseCase -.
type productUseCase struct {
	repo ProductRepo
}

// New -.
func NewProduct(r ProductRepo) *productUseCase {
	return &productUseCase{
		repo: r,
	}
}

// History - getting translate history from store.
func (uc *productUseCase) Get(ctx context.Context) ([]entity.Product, error) {
	products, err := uc.repo.All(ctx)
	if err != nil {
		return nil, fmt.Errorf("productUseCase - History - s.repo.GetHistory: %w", err)
	}

	return products, nil
}

func (uc *productUseCase) Translate(ctx context.Context, t entity.Product) (entity.Product, error) {
	// product, err := uc.webAPI.Translate(t)
	// if err != nil {
	// 	return entity.Product{}, fmt.Errorf("productUseCase - Translate - s.webAPI.Translate: %w", err)
	// }

	// err = uc.repo.Store(context.Background(), product)
	// if err != nil {
	// 	return entity.Product{}, fmt.Errorf("productUseCase - Translate - s.repo.Store: %w", err)
	// }

	return entity.Product{}, nil
}
