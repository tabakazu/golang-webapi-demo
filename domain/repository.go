package domain

import (
	"github.com/tabakazu/golang-webapi-demo/domain/collection"
	"github.com/tabakazu/golang-webapi-demo/domain/entity"
	"github.com/tabakazu/golang-webapi-demo/domain/value"
)

type ItemRepository interface {
	FindAll() (collection.Items, error)
	Find(value.ItemID) (entity.Item, error)
	Create(*entity.Item) error
	UpdateAttributes(*entity.Item, map[string]interface{}) error
	Delete(*entity.Item) error
}
