package redis

import (
	"context"
	"log"
	"time"

	"github.com/anthdm/ggcommerce/store"
	"github.com/anthdm/ggcommerce/types"
	"github.com/go-redis/redis/v8"
	"go.mongodb.org/mongo-driver/mongo"
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
        err := Client.HSet(ctx, v.ID, serializeProduct(v), time.Duration(10) * time.Minute).Err()
        if err != nil{
            log.Println("Error loading a product")
        }
    
    }
    log.Println("Preloaded products")
}

func serializeProduct(product *types.Product) map[string]string{
    return map[string]string{
        "id":product.ID,
        "name":product.Name,
        "sku":product.SKU,
    }
} 


func deserialize(product map[string]string) *types.Product{
    return &types.Product{
        ID: product["id"],
        SKU: product["sku"],
        Name: product["name"],
    }
}


func GetProducts(productID string) (*types.Product, error) {
    ctx := context.Background()
    productCache, err := Client.HGetAll(ctx, productID).Result()
    if err == redis.Nil{
        return nil, err 
    }
    if err != nil{
        return nil, err 
    }
    product := deserialize(productCache)
    return product, nil 
}










