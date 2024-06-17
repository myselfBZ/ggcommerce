package redis

import (
	"context"
	"log"
	"time"

	"github.com/anthdm/ggcommerce/store"
	"go.mongodb.org/mongo-driver/mongo"
    "github.com/go-redis/redis/v8"
	"go.mongodb.org/mongo-driver/mongo/options"
)




func PreloadProducts() {
    ctx := context.Background()
	client, err := mongo.Connect(context.TODO(), options.Client().ApplyURI("mongodb://localhost:27017"))
	if err != nil {
		panic(err)
	}
	productStore := store.NewMongoProductStore(client.Database("ggcommerce"))
    products, err := productStore.GetAll(ctx)
    for _, v := range products{
        err := Client.Set(ctx, v.ID, v.Name, time.Duration(10) * time.Minute).Err()
        if err != nil{
            log.Println("Error loading a product")
        }
    
    }
    log.Println("Preloaded products")
}



func GetProducts(productID string) (string, error) {
    ctx := context.Background()
    product, err := Client.Get(ctx, productID).Result()
    if err == redis.Nil{
        return "", err 
    }
    if err != nil{
        return "", err 
    }

    return product, nil 
}










