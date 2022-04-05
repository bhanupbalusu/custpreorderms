package user_auth

import (
	"context"
	"fmt"

	m "github.com/bhanupbalusu/custpreorderms/domain/model/user_auth"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func (r *MongoRepository) Get() (m.UserList, error) {
	var users m.UserList
	ctx, cancel := context.WithTimeout(context.Background(), r.Timeout)
	defer cancel()
	coll := r.Client.Database(r.DB).Collection("user_auth_coll1")
	cursor, err := coll.Find(ctx, bson.M{})
	if err = cursor.All(ctx, &users); err != nil {
		return nil, err
	}
	return users, err
}

func (r *MongoRepository) GetByID(id string) (*m.User, error) {
	var user m.User
	ctx, cancel := context.WithTimeout(context.Background(), r.Timeout)
	defer cancel()
	coll := r.Client.Database(r.DB).Collection("user_auth_coll1")
	newId, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return nil, err
	}
	filter := bson.M{"_id": newId}
	err = coll.FindOne(ctx, filter).Decode(&user)
	if err != nil {
		return nil, err
	}
	return &user, nil
}

func (r *MongoRepository) GetByEmail(email string) (*m.User, error) {
	var user m.User
	ctx, cancel := context.WithTimeout(context.Background(), r.Timeout)
	defer cancel()
	coll := r.Client.Database(r.DB).Collection("user_auth_coll1")
	filter := bson.M{"email": email}
	fmt.Println(coll)
	fmt.Println(filter)
	fmt.Println("---------Mongodb.GetBYEmail before calling coll.FindOne---------")

	err := coll.FindOne(ctx, filter).Decode(&user)
	if err != nil {
		fmt.Println("!!!!!!!!!! error after calling FindOne")
		return nil, err
	}
	return &user, nil
}

func (r *MongoRepository) Create(u *m.User) error {
	ctx, cancel := context.WithTimeout(context.Background(), r.Timeout)
	defer cancel()
	coll := r.Client.Database(r.DB).Collection("user_auth_coll1")
	fmt.Println("---------Mongodb.Create before calling coll.InsertOne---------")
	_, err := coll.InsertOne(
		ctx,
		bson.M{
			"id":         u.Id,
			"email":      u.Email,
			"password":   u.Password,
			"created_at": u.CreatedAt,
			"updated_at": u.UpdatedAt,
		},
	)
	return err
}

func (r *MongoRepository) Update(u *m.User) error {
	ctx, cancel := context.WithTimeout(context.Background(), r.Timeout)
	defer cancel()
	coll := r.Client.Database(r.DB).Collection("user_auth_coll1")
	filter := bson.M{"_id": u.Id}
	update := bson.M{
		"$set": bson.M{
			"id":         u.Id,
			"email":      u.Email,
			"password":   u.Password,
			"created_at": u.CreatedAt,
			"updated_at": u.UpdatedAt,
		},
	}
	_, err := coll.UpdateOne(ctx, filter, update)
	return err
}

func (r *MongoRepository) Delete(id string) error {
	ctx, cancel := context.WithTimeout(context.Background(), r.Timeout)
	defer cancel()
	coll := r.Client.Database(r.DB).Collection("user_auth_coll1")
	pid, err := primitive.ObjectIDFromHex(id)
	_, err = coll.DeleteOne(ctx, bson.M{"_id": pid})
	return err
}
