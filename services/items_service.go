package services

import (
	"github.com/leslesnoa/bookstore_items-api/domain/items"
	restErr "github.com/leslesnoa/bookstore_items-api/utils/errors"
)

var (
	ItemsService ItemsServiceInterface = &itemsService{}
)

type ItemsServiceInterface interface {
	Create(item items.Item) (*items.Item, *restErr.RestErr)
	Get(string) (*items.Item, *restErr.RestErr)
}

type itemsService struct{}

func (s *itemsService) Create(items.Item) (*items.Item, *restErr.RestErr) {

	return nil, restErr.NewInternalServerError("not_implemented")
}

func (s *itemsService) Get(string) (*items.Item, *restErr.RestErr) {
	return nil, restErr.NewInternalServerError("not_implemented")
}
