package product_details

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type ProductDetails struct {
	ProductName string `json:"product_name,omitempty" bson:"product_name,omitempty"`
	Description string `json:"description,omitempty" bson:"description,omitempty"`
	ImageURL    string `json:"image_url,omitempty" bson:"ImageUrl,omitempty"`
}

type BulkQuantity struct {
	Volume string `json:"volume,omitempty" bson:"volume,omitempty"`
	Units  string `json:"units,omitempty" bson:"units,omitempty"`
}

type Price struct {
	Amount   string `json:"amount,omitempty" bson:"amount,omitempty"`
	Currency string `json:"currency,omitempty" bson:"currency,omitempty"`
	PerUnit  string `json:"per_unit,omitempty" bson:"per_unit,omitempty"`
	Units    string `json:"units,omitempty" bson:"units,omitempty"`
}

type QuantityDetails struct {
	BulkQuantity BulkQuantity `json:"bulk_quantity,omitempty" bson:"bulk_quantity,omitempty"`
	Price        Price        `json:"price,omitempty" bson:"price,omitempty"`
}

type Schedular struct {
	StartDate string `json:"start_date,omitempty" bson:"start_date,omitempty"`
	EndDate   string `json:"end_date,omitempty" bson:"end_date,omitempty"`
}

type ProductDetailsModel struct {
	ProductId       primitive.ObjectID `json:"product_id,omitempty" bson:"_id,omitempty"`
	PreOrderId      string             `json:"pre_order_id,omitempty" bson:"pre_order_request_id,omitempty"`
	CustomerId      string             `json:"customer_id,omitempty" bson:"customer_id,omitempty"`
	ProductDetails  ProductDetails     `json:"product_details,omitempty" bson:"product_details,omitempty"`
	QuantityDetails QuantityDetails    `json:"quantity_details,omitempty" bson:"quantity_details,omitempty"`
	Schedular       Schedular          `json:"schedular,omitempty" bson:"schedular,omitempty"`
	CreatedAt       time.Time          `json:"created_at" bson:"created_at"`
	UpdatedAt       time.Time          `json:"updated_at" bson:"updated_at"`
}
