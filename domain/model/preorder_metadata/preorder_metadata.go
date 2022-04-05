package preorder_metadata

import "go.mongodb.org/mongo-driver/bson/primitive"

type PreOrderMetaDataModel struct {
	PreOrderId primitive.ObjectID `json:"preorder_id,omitempty" bson:"_id,omitempty"`
	CustomerId string             `json:"customer_id,omitempty" bson:"customer_id,omitempty"`
}

type PreOrderMetaDataModelList *[]PreOrderMetaDataModel
