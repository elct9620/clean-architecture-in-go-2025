package repository

import (
	"context"
	"encoding/json"

	"github.com/elct9620/clean-architecture-in-go-2025/internal/entity/orders"
	bolt "go.etcd.io/bbolt"
)

const BoltOrdersTableName = "orders"

type BoltOrderItemSchema struct {
	Name      string
	Quantity  int
	UnitPrice int
}

type BoltOrderSchema struct {
	Id           string
	CustomerName string
	Items        []BoltOrderItemSchema
}

type BoltOrderRepository struct {
	db        *bolt.DB
	tableName string
}

func NewBoltOrderRepository(db *bolt.DB) *BoltOrderRepository {
	return &BoltOrderRepository{
		db:        db,
		tableName: BoltOrdersTableName,
	}
}

func (r *BoltOrderRepository) Find(ctx context.Context, id string) (*orders.Order, error) {
	var order BoltOrderSchema
	err := r.db.View(func(tx *bolt.Tx) error {
		bucket := tx.Bucket([]byte(r.tableName))
		if bucket == nil {
			return orders.ErrOrderNotFound
		}

		boltValue := bucket.Get([]byte(id))
		if boltValue == nil {
			return orders.ErrOrderNotFound
		}

		return json.Unmarshal(boltValue, &order)
	})

	if err != nil {
		return nil, err
	}

	orderEntity := orders.New(order.Id, order.CustomerName)
	for _, item := range order.Items {
		err := orderEntity.AddItem(item.Name, item.Quantity, item.UnitPrice)
		if err != nil {
			return nil, err
		}
	}

	return orderEntity, nil
}

func (r *BoltOrderRepository) Save(ctx context.Context, order *orders.Order) error {
	boltOrder := BoltOrderSchema{
		Id:           order.Id(),
		CustomerName: order.CustomerName(),
		Items:        []BoltOrderItemSchema{},
	}

	for _, item := range order.Items() {
		boltOrder.Items = append(boltOrder.Items, BoltOrderItemSchema{
			Name:      item.Name(),
			Quantity:  item.Quantity(),
			UnitPrice: item.UnitPrice(),
		})
	}

	return r.db.Update(func(tx *bolt.Tx) error {
		bucket, err := tx.CreateBucketIfNotExists([]byte(r.tableName))
		if err != nil {
			return err
		}

		boltValue, err := json.Marshal(boltOrder)
		if err != nil {
			return err
		}

		return bucket.Put([]byte(order.Id()), boltValue)
	})
}
