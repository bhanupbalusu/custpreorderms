package preorder_metadata

import (
	"context"
	"fmt"
	"log"

	model "github.com/bhanupbalusu/custpreorderms/domain/model/preorder_metadata"
	"github.com/pkg/errors"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func (r *MongoRepository) Get() (*[]model.PreOrderMetaDataModel, error) {
	var results []model.PreOrderMetaDataModel

	ctx, cancel := context.WithTimeout(context.Background(), r.Timeout)
	defer cancel()
	coll := r.Client.Database(r.DB).Collection("preorder_metadata_coll1")
	fmt.Println(coll)

	cursor, err := coll.Find(ctx, bson.M{})
	fmt.Println("--------Execution After Find method-----")
	if err != nil {
		return nil, err
	}
	fmt.Println("--------Execution After Find method Before Iteration-----")
	if err = cursor.All(ctx, &results); err != nil {
		errors.Wrap(err, "db.repository.Get.cursor.All")
		log.Fatal(err)
	}
	return &results, nil
}

func (r *MongoRepository) GetByID(id string) (*model.PreOrderMetaDataModel, error) {
	var result model.PreOrderMetaDataModel
	fmt.Println("------- Inside repository.GetByID Before Calling r.GetCollection -----------")
	ctx, cancel := context.WithTimeout(context.Background(), r.Timeout)
	defer cancel()
	coll := r.Client.Database(r.DB).Collection("preorder_metadata_coll1")
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

func (r *MongoRepository) Create(pomd *model.PreOrderMetaDataModel) (string, error) {
	fmt.Println("------- Inside repository.Create Before Calling r.GetCollection -----------")

	ctx, cancel := context.WithTimeout(context.Background(), r.Timeout)
	defer cancel()
	coll := r.Client.Database(r.DB).Collection("preorder_metadata_coll1")
	fmt.Println(coll)

	fmt.Println("------- Inside db/repository.Create Before Calling coll.InsertOne -----------")
	result, err := coll.InsertOne(
		ctx,
		bson.M{
			"customer_id": pomd.CustomerId,
			"created_at":  pomd.CreatedAt,
			"updated_at":  pomd.UpdatedAt,
		},
	)
	if err != nil {
		return "", errors.Wrap(err, "repository.Create")
	}
	pid := (result.InsertedID).(primitive.ObjectID).Hex()
	return pid, nil
}

func (r *MongoRepository) Update(pomd *model.PreOrderMetaDataModel) error {
	ctx, cancel := context.WithTimeout(context.Background(), r.Timeout)
	defer cancel()
	coll := r.Client.Database(r.DB).Collection("preorder_metadata_coll1")
	filter := bson.M{"_id": pomd.PreOrderId}
	update := bson.M{
		"$set": bson.M{
			"customer_id": pomd.CustomerId,
			"updated_at":  pomd.UpdatedAt,
		},
	}
	_, err := coll.UpdateOne(ctx, filter, update)
	return err
}

func (r *MongoRepository) Delete(id string) error {

	ctx, cancel := context.WithTimeout(context.Background(), r.Timeout)
	defer cancel()
	coll := r.Client.Database(r.DB).Collection("preorder_metadata_coll1")
	fmt.Println(coll)

	pid, err := primitive.ObjectIDFromHex(id)
	_, err = coll.DeleteOne(ctx, bson.M{"_id": pid})
	if err != nil {
		return err
	}
	return nil
}
