// package user

// import (
// 	"context"
// 	"testing"

// 	"github.com/DATA-DOG/go-sqlmock"
// 	"github.com/Germanchrystan/GeekStore/api/internal/domain"
// 	"github.com/stretchr/testify/assert"
// )

// func Test_AddAddress_Ok(t *testing.T) {
// 	db, mock, err := sqlmock.New()
// 	assert.NoError(t, err)
// 	defer db.Close()

// 	mock.ExpectPrepare("INSERT INTO addresses")
// 	mock.ExpectExec("INSERT INTO addresses").WillReturnResult(sqlmock.NewResult(1, 1))

// 	r := NewRepository(db)
// 	ctx := context.TODO()
// 	savedAddress := domain.Address{
// 		ID:           "1",
// 		Street:       "street",
// 		StreetNumber: 123,
// 		State:        "Buenos Aires",
// 		Country:      "Argentina",
// 		Zipcode:      "123456",
// 	}
// 	_, err := r.AddAddress()
// }