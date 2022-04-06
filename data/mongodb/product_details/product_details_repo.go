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

// Get all products
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

// Get single product using id
func (r *MongoRepository) GetByID(id string) (*model.ProductDetailsModel, error) {
	var result model.ProductDetailsModel
	fmt.Println("------- Inside repository.GetByID Before Calling r.GetCollection -----------")
	ctx, cancel := context.WithTimeout(context.Background(), r.Timeout)
	defer cancel()
	coll := r.Client.Database(r.DB).Collection("product_details_coll1")
	fmt.Println(coll)

	newId, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		log.Fatal(err)
	}
	filter := bson.M{"_id": newId}

	fmt.Println(newId)
	fmt.Println("------- Inside repository.GetByID Before Calling coll.FindOne -----------")
	err = coll.FindOne(ctx, filter).Decode(&result)
	if err != nil {
		return &result, err
	}
	return &result, nil
}

// Create or insert a new product
func (r *MongoRepository) Create(pm *model.ProductDetailsModel) (string, error) {
	fmt.Println("------- Inside repository.Create Before Calling r.GetCollection -----------")

	ctx, cancel := context.WithTimeout(context.Background(), r.Timeout)
	defer cancel()
	coll := r.Client.Database(r.DB).Collection("product_details_coll1")
	fmt.Println(coll)

	fmt.Println("------- Inside db/repository.Create Before Calling coll.InsertOne -----------")
	result, err := coll.InsertOne(
		ctx,
		bson.M{
			"pre_order_request_id": pm.PreOrderRequestId,
			"customer_id":          pm.CustomerId,
			"product_details": bson.M{
				"product_name": pm.ProductDetails.ProductName,
				"description":  pm.ProductDetails.Description,
				"image_url":    pm.ProductDetails.ImageURL,
			},
			"quantity_details": bson.M{
				"bulk_quantity": bson.M{
					"volume": pm.QuantityDetails.BulkQuantity.Volume,
					"units":  pm.QuantityDetails.BulkQuantity.Units,
				},
				"price": bson.M{
					"amount":   pm.QuantityDetails.Price.Amount,
					"currency": pm.QuantityDetails.Price.Currency,
					"per_unit": pm.QuantityDetails.Price.PerUnit,
					"units":    pm.QuantityDetails.Price.Units,
				},
			},
			"schedular": bson.M{
				"start_date": pm.Schedular.StartDate,
				"end_date":   pm.Schedular.EndDate,
			},
			"created_at": pm.CreatedAt,
			"updated_at": pm.UpdatedAt,
		},
	)
	if err != nil {
		return "", errors.Wrap(err, "repository.Create")
	}
	pid := (result.InsertedID).(primitive.ObjectID).Hex()
	return pid, nil
}

// Update existing product
func (r *MongoRepository) Update(pm *model.ProductDetailsModel) error {

	ctx, cancel := context.WithTimeout(context.Background(), r.Timeout)
	defer cancel()
	coll := r.Client.Database(r.DB).Collection("product_details_coll1")
	fmt.Println(coll, ctx)

	// pid, err := primitive.ObjectIDFromHex(id)
	// if err != nil {
	// 	log.Fatal(err)
	// }

	filter := bson.M{"_id": pm.ProductId}

	update := bson.M{
		"$set": bson.M{
			"pre_order_request_id":                  pm.PreOrderRequestId,
			"customer_id":                           pm.CustomerId,
			"product_details.product_name":          pm.ProductDetails.ProductName,
			"product_details.description":           pm.ProductDetails.Description,
			"product_details.image_url":             pm.ProductDetails.ImageURL,
			"quantity_details.bulk_quantity.volume": pm.QuantityDetails.BulkQuantity.Volume,
			"quantity_details.bulk_quantity.units":  pm.QuantityDetails.BulkQuantity.Units,
			"quantity_details.price.amount":         pm.QuantityDetails.Price.Amount,
			"quantity_details.price.currency":       pm.QuantityDetails.Price.Currency,
			"quantity_details.price.per_unit":       pm.QuantityDetails.Price.PerUnit,
			"quantity_details.price.units":          pm.QuantityDetails.Price.Units,
			"schedular.start_date":                  pm.Schedular.StartDate,
			"schedular.end_date":                    pm.Schedular.EndDate,
		},
	}

	fmt.Println(update)

	_, err := coll.UpdateOne(ctx, filter, update)

	return err
}

// Delete an existing product
func (r *MongoRepository) Delete(id string) error {

	ctx, cancel := context.WithTimeout(context.Background(), r.Timeout)
	defer cancel()
	coll := r.Client.Database(r.DB).Collection("product_details_coll1")
	fmt.Println(coll)

	pid, err := primitive.ObjectIDFromHex(id)
	_, err = coll.DeleteOne(ctx, bson.M{"_id": pid})
	if err != nil {
		return err
	}
	return nil
}
