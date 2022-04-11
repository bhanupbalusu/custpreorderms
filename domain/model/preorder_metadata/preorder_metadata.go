package preorder_metadata

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type PreOrderMetaDataModel struct {
	PreOrderId primitive.ObjectID `json:"pre_order_id,omitempty" bson:"_id,omitempty"`
	CustomerId string             `json:"customer_id,omitempty" bson:"customer_id,omitempty"`
	CreatedAt  time.Time          `json:"created_at" bson:"created_at"`
	UpdatedAt  time.Time          `json:"updated_at" bson:"updated_at"`
}
