package deal_settings

import (
	"context"
	"fmt"
	"log"

	model "github.com/bhanupbalusu/custpreorderms/domain/model/deal_settings"
	"github.com/pkg/errors"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

// Get all products
func (r *MongoRepository) Get() (model.DealSettingsModelList, error) {
	var results model.DealSettingsModelList

	ctx, cancel := context.WithTimeout(context.Background(), r.Timeout)
	defer cancel()
	coll := r.Client.Database(r.DB).Collection("deal_settings_coll1")
	fmt.Println(coll)

	cursor, err := coll.Find(ctx, bson.M{})
	if err != nil {
		return nil, err
	}

	if err = cursor.All(ctx, &results); err != nil {
		errors.Wrap(err, "db.repository.Get.cursor.All")
		log.Fatal(err)
	}
	return results, nil
}

// Get single product using id
func (r *MongoRepository) GetByID(id string) (*model.DealSettingsModel, error) {
	var result model.DealSettingsModel
	fmt.Println("------- Inside repository.GetByID Before Calling r.GetCollection -----------")
	ctx, cancel := context.WithTimeout(context.Background(), r.Timeout)
	defer cancel()
	coll := r.Client.Database(r.DB).Collection("deal_settings_coll1")
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
func (r *MongoRepository) Create(dsm *model.DealSettingsModel) (string, error) {
	fmt.Println("------- Inside repository.Create Before Calling r.GetCollection -----------")

	ctx, cancel := context.WithTimeout(context.Background(), r.Timeout)
	defer cancel()
	coll := r.Client.Database(r.DB).Collection("deal_settings_coll1")
	fmt.Println(coll)

	fmt.Println("------- Inside db/repository.Create Before Calling coll.InsertOne -----------")
	result, err := coll.InsertOne(
		ctx,
		bson.M{
			"pre_order_request_id": dsm.PreOrderRequestId,
			"customer_id":          dsm.CustomerId,
			"product_id":           dsm.ProductId,
			"order_types_id":       dsm.DealSettingsId,
			"deal_settings": bson.M{
				"must_satisfy_bulk_quantity_volume": dsm.DealSettings.MustSatisfyBulkQuantityVolume,
				"accept_over_subscriptions": bson.M{
					"value":                              dsm.DealSettings.AcceptOverSubscriptions.Value,
					"percentage_on_bulk_quantity_volume": dsm.DealSettings.AcceptOverSubscriptions.PercentageOnBulkQuantityVolume,
				},
				"accept_under_subscriptions": bson.M{
					"value":                              dsm.DealSettings.AcceptUnderSubscriptions.Value,
					"percentage_on_bulk_quantity_volume": dsm.DealSettings.AcceptUnderSubscriptions.PercentageOnBulkQuantityVolume,
				},
				"delivery_options": bson.M{
					"collection": dsm.DealSettings.DeliveryOptions.Collection,
					"by_post": bson.M{
						"value":   dsm.DealSettings.DeliveryOptions.ByPost.Value,
						"charges": dsm.DealSettings.DeliveryOptions.ByPost.Charges,
					},
				},
				"payment_options": bson.M{
					"cash":    dsm.DealSettings.PaymentOptions.Cash,
					"digital": dsm.DealSettings.PaymentOptions.Digital,
				},
			},
		},
	)
	if err != nil {
		return "", errors.Wrap(err, "repository.Create")
	}
	pid := (result.InsertedID).(primitive.ObjectID).Hex()
	return pid, nil
}

// Update existing product
func (r *MongoRepository) Update(dsm *model.DealSettingsModel) error {

	ctx, cancel := context.WithTimeout(context.Background(), r.Timeout)
	defer cancel()
	coll := r.Client.Database(r.DB).Collection("deal_settings_coll1")
	fmt.Println(coll, ctx)

	filter := bson.M{"_id": dsm.OrderTypesId}

	update := bson.M{
		"$set": bson.M{
			"pre_order_request_id": dsm.PreOrderRequestId,
			"customer_id":          dsm.CustomerId,
			"product_id":           dsm.ProductId,
			"order_types_id":       dsm.DealSettingsId,
			"deal_settings.must_satisfy_bulk_quantity_volume":                             dsm.DealSettings.MustSatisfyBulkQuantityVolume,
			"deal_settings.accept_over_subscriptions.value":                               dsm.DealSettings.AcceptOverSubscriptions.Value,
			"deal_settings.accept_over_subscriptions.percentage_on_bulk_quantity_volume":  dsm.DealSettings.AcceptOverSubscriptions.PercentageOnBulkQuantityVolume,
			"deal_settings.accept_under_subscriptions.value":                              dsm.DealSettings.AcceptUnderSubscriptions.Value,
			"deal_settings.accept_under_subscriptions.percentage_on_bulk_quantity_volume": dsm.DealSettings.AcceptUnderSubscriptions.PercentageOnBulkQuantityVolume,
			"deal_settings.delivery_options.collection":                                   dsm.DealSettings.DeliveryOptions.Collection,
			"deal_settings.delivery_options.by_post.value":                                dsm.DealSettings.DeliveryOptions.ByPost.Value,
			"deal_settings.delivery_options.by_post.charges":                              dsm.DealSettings.DeliveryOptions.ByPost.Charges,
			"deal_settings.payment_options.cash":                                          dsm.DealSettings.PaymentOptions.Cash,
			"deal_settings.payment_options.digital":                                       dsm.DealSettings.PaymentOptions.Digital,
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
	coll := r.Client.Database(r.DB).Collection("deal_settings_coll1")
	fmt.Println(coll)

	pid, err := primitive.ObjectIDFromHex(id)
	_, err = coll.DeleteOne(ctx, bson.M{"_id": pid})
	if err != nil {
		return err
	}
	return nil
}
