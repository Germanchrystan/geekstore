package admin

import (
	"context"
	"errors"
	"reflect"
	"strconv"

	"github.com/Germanchrystan/GeekStore/api/internal/domain"
)

type AdminService interface {
	GetAllUsers(ctx context.Context) ([]domain.User, error)
	//-----------------------------------------------------//
	PostProduct(ctx context.Context, product domain.Product) (domain.Product, error)
	DeleteProduct(ctx context.Context, product_id string) error
	UpdateProduct(ctx context.Context, product domain.Product) (domain.Product, error)
	//-----------------------------------------------------//
	ToggleUserBan(ctx context.Context, user_id string) error
	ToggleUserAdmin(ctx context.Context, user_id string) error
	//-----------------------------------------------------//
}

//===================================================================================================//

type service struct {
	repository AdminRepository
}

// ===========================================================//
func NewService(repository AdminRepository) AdminService {
	return &service{
		repository: repository,
	}
}

//===========================================================//

func (s *service) GetAllUsers(ctx context.Context) ([]domain.User, error) {
	return s.repository.GetAllUsers(ctx)
}

//===========================================================//

func (s *service) PostProduct(ctx context.Context, product domain.Product) (domain.Product, error) {
	return s.repository.PostProduct(ctx, product)
}

//===========================================================//

func (s *service) DeleteProduct(ctx context.Context, product_id string) error {
	return s.repository.DeleteProduct(ctx, product_id)
}

//===========================================================//

func (s *service) UpdateProduct(ctx context.Context, product domain.Product) (domain.Product, error) {
	previousProduct, notFoundErr := s.repository.GetProductById(ctx, strconv.Itoa(product.Id))
	if notFoundErr != nil {
		return domain.Product{}, errors.New("product could not be retrieved")
	}
	// Fetching values of previous product state
	previousProductValue := reflect.ValueOf(previousProduct)

	productTypes := reflect.TypeOf(product)
	referenceOfProduct := reflect.ValueOf(&product).Elem()

	productNumFields := productTypes.NumField()

	for i := 0; i < productNumFields; i++ {
		fieldType := referenceOfProduct.Field(i)

		previousValue := previousProductValue.Field(i)

		if fieldType.IsZero() {
			if fieldType.Kind() == reflect.Int {
				var PrevIntValue int = int(reflect.ValueOf(previousValue.Int()).Int())
				referenceOfProduct.Set(reflect.ValueOf(PrevIntValue))
			} else if fieldType.Kind() == reflect.Float32 {
				var PrevFloatValue float32 = float32(reflect.ValueOf(previousValue.Float()).Float())
				referenceOfProduct.Set(reflect.ValueOf(PrevFloatValue))
			} else if fieldType.Kind() == reflect.String {
				referenceOfProduct.Set(reflect.ValueOf(previousValue.String()))
			}
		}
	}
	updatedProduct, updateErr := s.repository.UpdateProduct(ctx, product)

	if updateErr != nil {
		return domain.Product{}, errors.New("product Could not be updated")
	} else {
		return updatedProduct, nil
	}
}

//===========================================================//

func (s *service) ToggleUserBan(ctx context.Context, user_id string) error {
	return s.repository.ToggleUserBan(ctx, user_id)
}

//===========================================================//

func (s *service) ToggleUserAdmin(ctx context.Context, user_id string) error {
	return s.repository.ToggleUserAdmin(ctx, user_id)
}

//===========================================================//
