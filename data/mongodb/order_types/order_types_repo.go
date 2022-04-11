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
func (r *MongoRepository) Get() (*[]model.OrderTypesModel, error) {
	var results []model.OrderTypesModel

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
	return &results, nil
}

// Get single product using id
func (r *MongoRepository) GetByID(id string) (*model.OrderTypesModel, error) {
	var result model.OrderTypesModel
	fmt.Println("------- Inside repository.GetByID Before Calling r.GetCollection -----------")
	ctx, cancel := context.WithTimeout(context.Background(), r.Timeout)
	defer cancel()
	coll := r.Client.Database(r.DB).Collection("order_types_coll1")
	fmt.Println(coll)

	// newId, err := primitive.ObjectIDFromHex(id)
	// if err != nil {
	// 	log.Fatal(err)
	// }

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
			"pre_order_id": otm.PreOrderId,
			"customer_id":  otm.CustomerId,
			"product_id":   otm.ProductId,
			"order_types": bson.M{
				"ot_individual": bson.M{
					"i_min_order": bson.M{
						"min_order_volume": otm.OrderTypes.OTIndividual.IMinOrder.IMinOrderVolume,
						"min_order_units":  otm.OrderTypes.OTIndividual.IMinOrder.IMinUnits,
					},
					"i_max_order": bson.M{
						"max_order_volume": otm.OrderTypes.OTIndividual.IMaxOrder.IMaxOrderVolume,
						"max_order_units":  otm.OrderTypes.OTIndividual.IMaxOrder.IMaxUnits,
					},
				},
				"ot_group": bson.M{
					"g_individual_min_order": bson.M{
						"individual_min_order_volume": otm.OrderTypes.OTGroup.GIndividualMinOrder.IndividualMinOrderVolume,
						"individual_min_order_units":  otm.OrderTypes.OTGroup.GIndividualMinOrder.IndividualMinOrderUnits,
					},
					"g_max_order": bson.M{
						"group_max_order_volume": otm.OrderTypes.OTGroup.GGroupMaxOrder.GroupMaxOrderVolume,
						"group_max_order_units":  otm.OrderTypes.OTGroup.GGroupMaxOrder.GroupMaxOrderUnits,
						"group_max_order_discount": bson.M{
							"gmod_percentage": otm.OrderTypes.OTGroup.GGroupMaxOrder.GroupMaxOrderDiscount.GMODPercentage,
							"gmod_split_options": bson.M{
								"so_group_owner": otm.OrderTypes.OTGroup.GGroupMaxOrder.GroupMaxOrderDiscount.GMODSplitOptions.SOGroupOwner,
								"so_type":        otm.OrderTypes.OTGroup.GGroupMaxOrder.GroupMaxOrderDiscount.GMODSplitOptions.SOType,
							},
						},
					},
					"g_settings": bson.M{
						"s_must_satisfy_max_volume": otm.OrderTypes.OTGroup.GSettings.MustSatisfyMaxVolume,
						"s_accept_over_subscriptions": bson.M{
							"aos_value":                    otm.OrderTypes.OTGroup.GSettings.AcceptOverSubscriptions.AOSValue,
							"aos_percentage_on_max_volume": otm.OrderTypes.OTGroup.GSettings.AcceptOverSubscriptions.AOSPercentageOnMaxVolume,
						},
						"s_accept_under_subscriptions": bson.M{
							"aus_value":                    otm.OrderTypes.OTGroup.GSettings.AcceptUnderSubscriptions.AUSValue,
							"aus_percentage_on_max_volume": otm.OrderTypes.OTGroup.GSettings.AcceptUnderSubscriptions.AUSPercentageOnMaxVolume,
						},
						"s_accept_single_order": bson.M{
							"aso_value": otm.OrderTypes.OTGroup.GSettings.AcceptSingleOrder.ASOValue,
							"aso_single_order_discount": bson.M{
								"sod_percentage": otm.OrderTypes.OTGroup.GSettings.AcceptSingleOrder.ASOSingleOrderDiscount.ASOSODPercentage,
							},
						},
					},
				},
			},
			"created_at": otm.CreatedAt,
			"updated_at": otm.UpdatedAt,
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

	filter := bson.M{"pre_order_id": otm.PreOrderId}

	update := bson.M{
		"$set": bson.M{
			"pre_order_id": otm.PreOrderId,
			"customer_id":  otm.CustomerId,
			"product_id":   otm.ProductId,
			"order_types.ot_individual.i_min_order.min_order_volume":                                         otm.OrderTypes.OTIndividual.IMinOrder.IMinOrderVolume,
			"order_types.ot_individual.i_min_order.min_order_units":                                          otm.OrderTypes.OTIndividual.IMinOrder.IMinUnits,
			"order_types.ot_individual.i_max_order.max_order_volume":                                         otm.OrderTypes.OTIndividual.IMaxOrder.IMaxOrderVolume,
			"order_types.ot_individual.i_max_order.max_order_units":                                          otm.OrderTypes.OTIndividual.IMaxOrder.IMaxUnits,
			"order_types.ot_group.g_individual_min_order.individual_min_order_volume":                        otm.OrderTypes.OTGroup.GIndividualMinOrder.IndividualMinOrderVolume,
			"order_types.ot_group.g_individual_min_order.individual_min_order_units":                         otm.OrderTypes.OTGroup.GIndividualMinOrder.IndividualMinOrderUnits,
			"order_types.ot_group.g_max_order.group_max_order_volume":                                        otm.OrderTypes.OTGroup.GGroupMaxOrder.GroupMaxOrderVolume,
			"order_types.ot_group.g_max_order.group_max_order_units":                                         otm.OrderTypes.OTGroup.GGroupMaxOrder.GroupMaxOrderUnits,
			"order_types.ot_group.g_max_order.group_max_order_discount.gmod_percentage":                      otm.OrderTypes.OTGroup.GGroupMaxOrder.GroupMaxOrderDiscount.GMODPercentage,
			"order_types.ot_group.g_max_order.group_max_order_discount.gmod_split_options.so_group_owner":    otm.OrderTypes.OTGroup.GGroupMaxOrder.GroupMaxOrderDiscount.GMODSplitOptions.SOGroupOwner,
			"order_types.ot_group.g_max_order.group_max_order_discount.gmod_split_options.so_type":           otm.OrderTypes.OTGroup.GGroupMaxOrder.GroupMaxOrderDiscount.GMODSplitOptions.SOType,
			"order_types.ot_group.g_settings.s_must_satisfy_max_volume":                                      otm.OrderTypes.OTGroup.GSettings.MustSatisfyMaxVolume,
			"order_types.ot_group.g_settings.s_accept_over_subscriptions.aos_value":                          otm.OrderTypes.OTGroup.GSettings.AcceptOverSubscriptions.AOSValue,
			"order_types.ot_group.g_settings.s_accept_over_subscriptions.aos_percentage_on_max_volume":       otm.OrderTypes.OTGroup.GSettings.AcceptOverSubscriptions.AOSPercentageOnMaxVolume,
			"order_types.ot_group.g_settings.s_accept_under_subscriptions.aus_value":                         otm.OrderTypes.OTGroup.GSettings.AcceptUnderSubscriptions.AUSValue,
			"order_types.ot_group.g_settings.s_accept_under_subscriptions.aus_percentage_on_max_volume":      otm.OrderTypes.OTGroup.GSettings.AcceptUnderSubscriptions.AUSPercentageOnMaxVolume,
			"order_types.ot_group.g_settings.s_accept_single_order.aso_value":                                otm.OrderTypes.OTGroup.GSettings.AcceptSingleOrder.ASOValue,
			"order_types.ot_group.g_settings.s_accept_single_order.aso_single_order_discount.sod_percentage": otm.OrderTypes.OTGroup.GSettings.AcceptSingleOrder.ASOSingleOrderDiscount.ASOSODPercentage,
			"updated_at": otm.UpdatedAt,
		},
	}

	//fmt.Println(update)

	_, err := coll.UpdateOne(ctx, filter, update)

	return err
}

// Delete an existing product
func (r *MongoRepository) Delete(id string) error {

	ctx, cancel := context.WithTimeout(context.Background(), r.Timeout)
	defer cancel()
	coll := r.Client.Database(r.DB).Collection("order_types_coll1")
	fmt.Println(coll)

	//pid, err := primitive.ObjectIDFromHex(id)
	_, err := coll.DeleteOne(ctx, bson.M{"pre_order_id": id})
	if err != nil {
		return err
	}
	return nil
}
