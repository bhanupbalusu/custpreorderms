package product_details

import (
	"context"
	"fmt"
	"log"

	model "github.com/bhanupbalusu/custpreorderms/domain/model/product_details"
	"github.com/pkg/errors"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func (r *MongoRepository) Get() (*[]model.ProductDetailsModel, error) {
	var results []model.ProductDetailsModel

	ctx, cancel := context.WithTimeout(context.Background(), r.Timeout)
	defer cancel()
	coll := r.Client.Database(r.DB).Collection("product_details_coll1")
	fmt.Println(coll)

	cursor, err := coll.Find(ctx, bson.M{})
	if err != nil {
		return nil, err
	}

	if err = cursor.All(ctx, &results); err != nil {
		errors.Wrap(err, "db.repository.Get.cursor.All")
		log.Fatal(err)
	}
	return &results, nil
}

func (r *MongoRepository) GetByID(id string) (*model.ProductDetailsModel, error) {
	var result model.ProductDetailsModel
	fmt.Println("------- Inside repository.GetByID Before Calling r.GetCollection -----------")
	ctx, cancel := context.WithTimeout(context.Background(), r.Timeout)
	defer cancel()
	coll := r.Client.Database(r.DB).Collection("product_details_coll1")
	fmt.Println(coll)

	filter := bson.M{"pre_order_id": id}

	fmt.Println(id)
	fmt.Println("------- Inside repository.GetByID Before Calling coll.FindOne -----------")
	err := coll.FindOne(ctx, filter).Decode(&result)
	if err != nil {
		return &result, err
	}
	return &result, nil
}

// Create or insert a new product
func (r *MongoRepository) Create(pdm *model.ProductDetailsModel) (string, error) {
	fmt.Println("------- Inside repository.Create Before Calling r.GetCollection -----------")

	ctx, cancel := context.WithTimeout(context.Background(), r.Timeout)
	defer cancel()
	coll := r.Client.Database(r.DB).Collection("product_details_coll1")
	fmt.Println(coll)

	fmt.Println("------- Inside db/repository.Create Before Calling coll.InsertOne -----------")
	result, err := coll.InsertOne(
		ctx,
		bson.M{
			"pre_order_id": pdm.PreOrderId,
			"customer_id":  pdm.CustomerId,
			"product_details": bson.M{
				"product_name": pdm.ProductDetails.ProductName,
				"description":  pdm.ProductDetails.Description,
				"image_url":    pdm.ProductDetails.ImageURL,
			},
			"quantity_details": bson.M{
				"bulk_quantity": bson.M{
					"volume": pdm.QuantityDetails.BulkQuantity.Volume,
					"units":  pdm.QuantityDetails.BulkQuantity.Units,
				},
				"price": bson.M{
					"amount":   pdm.QuantityDetails.Price.Amount,
					"currency": pdm.QuantityDetails.Price.Currency,
					"per_unit": pdm.QuantityDetails.Price.PerUnit,
					"units":    pdm.QuantityDetails.Price.Units,
				},
			},
			"schedular": bson.M{
				"start_date": pdm.Schedular.StartDate,
				"end_date":   pdm.Schedular.EndDate,
			},
			"created_at": pdm.CreatedAt,
			"updated_at": pdm.UpdatedAt,
		},
	)
	if err != nil {
		return "", errors.Wrap(err, "repository.Create")
	}
	pid := (result.InsertedID).(primitive.ObjectID).Hex()
	return pid, nil
}

func (r *MongoRepository) Update(pdm *model.ProductDetailsModel) error {

	ctx, cancel := context.WithTimeout(context.Background(), r.Timeout)
	defer cancel()
	coll := r.Client.Database(r.DB).Collection("product_details_coll1")
	fmt.Println(coll, ctx)

	filter := bson.M{"pre_order_id": pdm.PreOrderId}

	update := bson.M{
		"$set": bson.M{
			"pre_order_id":                          pdm.PreOrderId,
			"customer_id":                           pdm.CustomerId,
			"product_details.product_name":          pdm.ProductDetails.ProductName,
			"product_details.description":           pdm.ProductDetails.Description,
			"product_details.image_url":             pdm.ProductDetails.ImageURL,
			"quantity_details.bulk_quantity.volume": pdm.QuantityDetails.BulkQuantity.Volume,
			"quantity_details.bulk_quantity.units":  pdm.QuantityDetails.BulkQuantity.Units,
			"quantity_details.price.amount":         pdm.QuantityDetails.Price.Amount,
			"quantity_details.price.currency":       pdm.QuantityDetails.Price.Currency,
			"quantity_details.price.per_unit":       pdm.QuantityDetails.Price.PerUnit,
			"quantity_details.price.units":          pdm.QuantityDetails.Price.Units,
			"schedular.start_date":                  pdm.Schedular.StartDate,
			"schedular.end_date":                    pdm.Schedular.EndDate,
			"updated_at":                            pdm.UpdatedAt,
		},
	}

	fmt.Println(update)

	_, err := coll.UpdateOne(ctx, filter, update)

	return err
}

func (r *MongoRepository) Delete(id string) error {

	ctx, cancel := context.WithTimeout(context.Background(), r.Timeout)
	defer cancel()
	coll := r.Client.Database(r.DB).Collection("product_details_coll1")
	fmt.Println(coll)

	_, err := coll.DeleteOne(ctx, bson.M{"pre_order_id": id})
	if err != nil {
		return err
	}
	return nil
}
