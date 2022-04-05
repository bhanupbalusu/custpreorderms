package order_types

import (
	"context"
	"fmt"
	"log"

	model "github.com/bhanupbalusu/custpreorderms/domain/model/order_types"
	"github.com/pkg/errors"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

// Get all products
func (r *MongoRepository) Get() (model.OrderTypesModelList, error) {
	var results model.OrderTypesModelList

	ctx, cancel := context.WithTimeout(context.Background(), r.Timeout)
	defer cancel()
	coll := r.Client.Database(r.DB).Collection("order_types_coll1")
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
func (r *MongoRepository) GetByID(id string) (*model.OrderTypesModel, error) {
	var result model.OrderTypesModel
	fmt.Println("------- Inside repository.GetByID Before Calling r.GetCollection -----------")
	ctx, cancel := context.WithTimeout(context.Background(), r.Timeout)
	defer cancel()
	coll := r.Client.Database(r.DB).Collection("order_types_coll1")
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
func (r *MongoRepository) Create(otm *model.OrderTypesModel) (string, error) {
	fmt.Println("------- Inside repository.Create Before Calling r.GetCollection -----------")

	ctx, cancel := context.WithTimeout(context.Background(), r.Timeout)
	defer cancel()
	coll := r.Client.Database(r.DB).Collection("order_types_coll1")
	fmt.Println(coll)

	fmt.Println("------- Inside db/repository.Create Before Calling coll.InsertOne -----------")
	result, err := coll.InsertOne(
		ctx,
		bson.M{
			"pre_order_request_id": otm.PreOrderRequestId,
			"customer_id":          otm.CustomerId,
			"product_id":           otm.ProductId,
			"order_types": bson.M{
				"individual": bson.M{
					"min_order": bson.M{
						"volume": otm.OrderTypes.Individual.MinOrder.Volume,
						"units":  otm.OrderTypes.Individual.MinOrder.Units,
					},
					"max_order": bson.M{
						"volume": otm.OrderTypes.Individual.MaxOrder.Volume,
						"units":  otm.OrderTypes.Individual.MaxOrder.Units,
					},
				},
				"group": bson.M{
					"individual_min_order": bson.M{
						"volume": otm.OrderTypes.Group.IndividualMinOrder.Volume,
						"units":  otm.OrderTypes.Group.IndividualMinOrder.Units,
					},
					"max_order": bson.M{
						"volume": otm.OrderTypes.Group.MaxOrder.Volume,
						"units":  otm.OrderTypes.Group.MaxOrder.Units,
						"discount": bson.M{
							"percentage": otm.OrderTypes.Group.MaxOrder.Discount.Percentage,
							"split_options": bson.M{
								"group_owner": otm.OrderTypes.Group.MaxOrder.Discount.SplitOptions.GroupOwner,
								"type":        otm.OrderTypes.Group.MaxOrder.Discount.SplitOptions.Type,
							},
						},
					},
					"settings": bson.M{
						"must_satisfy_max_volume": otm.OrderTypes.Group.Settings.MustSatisfyMaxVolume,
						"accept_over_subscriptions": bson.M{
							"value":                    otm.OrderTypes.Group.Settings.AcceptOverSubscriptions.Value,
							"percentage_on_max_volume": otm.OrderTypes.Group.Settings.AcceptOverSubscriptions.PercentageOnMaxVolume,
						},
						"accept_under_subscriptions": bson.M{
							"value":                    otm.OrderTypes.Group.Settings.AcceptUnderSubscriptions.Value,
							"percentage_on_max_volume": otm.OrderTypes.Group.Settings.AcceptUnderSubscriptions.PercentageOnMaxVolume,
						},
						"accept_single_order": bson.M{
							"value": otm.OrderTypes.Group.Settings.AcceptSingleOrder.Value,
							"discount": bson.M{
								"percentage": otm.OrderTypes.Group.Settings.AcceptSingleOrder.Discount.Percentage,
							},
						},
					},
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
func (r *MongoRepository) Update(otm *model.OrderTypesModel) error {

	ctx, cancel := context.WithTimeout(context.Background(), r.Timeout)
	defer cancel()
	coll := r.Client.Database(r.DB).Collection("order_types_coll1")
	fmt.Println(coll, ctx)

	filter := bson.M{"_id": otm.OrderTypesId}

	update := bson.M{
		"$set": bson.M{
			"pre_order_request_id": otm.PreOrderRequestId,
			"customer_id":          otm.CustomerId,
			"product_id":           otm.ProductId,
			"order_types.individual.min_order.volume":                                        otm.OrderTypes.Individual.MinOrder.Volume,
			"order_types.individual.min_order.units":                                         otm.OrderTypes.Individual.MinOrder.Units,
			"order_types.individual.max_order.volume":                                        otm.OrderTypes.Individual.MaxOrder.Volume,
			"order_types.individual.max_order.units":                                         otm.OrderTypes.Individual.MaxOrder.Units,
			"order_types.group.individual_min_order.volume":                                  otm.OrderTypes.Group.IndividualMinOrder.Volume,
			"order_types.group.individual_min_order.units":                                   otm.OrderTypes.Group.IndividualMinOrder.Units,
			"order_types.group.max_order.volume":                                             otm.OrderTypes.Group.MaxOrder.Volume,
			"order_types.group.max_order.units":                                              otm.OrderTypes.Group.MaxOrder.Units,
			"order_types.group.max_order.discount.percentage":                                otm.OrderTypes.Group.MaxOrder.Discount.Percentage,
			"order_types.group.max_order.discount.split_options.group_owner":                 otm.OrderTypes.Group.MaxOrder.Discount.SplitOptions.GroupOwner,
			"order_types.group.max_order.discount.split_options.type":                        otm.OrderTypes.Group.MaxOrder.Discount.SplitOptions.Type,
			"order_types.group.settings.must_satisfy_max_volume":                             otm.OrderTypes.Group.Settings.MustSatisfyMaxVolume,
			"order_types.group.settings.accept_over_subscriptions.value":                     otm.OrderTypes.Group.Settings.AcceptOverSubscriptions.Value,
			"order_types.group.settings.accept_over_subscriptions.percentage_on_max_volume":  otm.OrderTypes.Group.Settings.AcceptOverSubscriptions.PercentageOnMaxVolume,
			"order_types.group.settings.accept_under_subscriptions.value":                    otm.OrderTypes.Group.Settings.AcceptUnderSubscriptions.Value,
			"order_types.group.settings.accept_under_subscriptions.percentage_on_max_volume": otm.OrderTypes.Group.Settings.AcceptUnderSubscriptions.PercentageOnMaxVolume,
			"order_types.group.settings.accept_single_order.value":                           otm.OrderTypes.Group.Settings.AcceptSingleOrder.Value,
			"order_types.group.settings.accept_single_order.discount.percentage":             otm.OrderTypes.Group.Settings.AcceptSingleOrder.Discount.Percentage,
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
	coll := r.Client.Database(r.DB).Collection("order_types_coll1")
	fmt.Println(coll)

	pid, err := primitive.ObjectIDFromHex(id)
	_, err = coll.DeleteOne(ctx, bson.M{"_id": pid})
	if err != nil {
		return err
	}
	return nil
}
