package store

import (
	"context"
	"errors"

	"github.com/anthdm/ggcommerce/types"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type MongoProductStore struct {
	db   *mongo.Database
	coll string
}







func NewMongoProductStore(db *mongo.Database) *MongoProductStore {
	return &MongoProductStore{
		db:   db,
		coll: "products",
	}
}

func (s *MongoProductStore) Insert(ctx context.Context, p *types.Product) error {
	res, err := s.db.Collection(s.coll).InsertOne(ctx, p)
	if err != nil {
		return err
	}
	p.ID = res.InsertedID.(primitive.ObjectID).Hex()

	return err
}

func (s *MongoProductStore) GetAll(ctx context.Context) ([]*types.Product, error) {
	cursor, err := s.db.Collection(s.coll).Find(ctx, map[string]any{})
	if err != nil {
		return nil, err
	}

	products := []*types.Product{}
	err = cursor.All(ctx, &products)
	return products, err
}

func (s *MongoProductStore) GetByID(ctx context.Context, id string) (*types.Product, error) {
	var (
		objID, _ = primitive.ObjectIDFromHex(id)
		res      = s.db.Collection(s.coll).FindOne(ctx, bson.M{"_id": objID})
		p        = &types.Product{}
		err      = res.Decode(p)
	)
	return p, err
}


func (s *MongoProductStore) Update(ctx context.Context, newP *types.Product, id string) error{
    update := bson.M{
        "$set":newP,
    }

    result, err := s.db.Collection(s.coll).UpdateByID(ctx, id, update)
    if err != nil{
        return err 
    }
    if result.MatchedCount == 0 {
        return errors.New("Product not found")
    }
    return nil 

}
