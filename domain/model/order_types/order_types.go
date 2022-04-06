package order_types

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type MinOrder struct {
	Volume string `json:"volume,omitempty" bson:"volume,omitempty"`
	Units  string `json:"units,omitempty" bson:"units,omitempty"`
}

type MaxOrder struct {
	Volume string `json:"volume,omitempty" bson:"volume,omitempty"`
	Units  string `json:"units,omitempty" bson:"units,omitempty"`
}

type Individual struct {
	MinOrder MinOrder `json:"min_order,omitempty" bson:"min_order,omitempty"`
	MaxOrder MaxOrder `json:"max_order,omitempty" bson:"max_order,omitempty"`
}

type IndividualMinOrder struct {
	Volume string `json:"volume,omitempty" bson:"volume,omitempty"`
	Units  string `json:"units,omitempty" bson:"units,omitempty"`
}

type SplitOptions struct {
	GroupOwner string `json:"group_owner,omitempty" bson:"group_owner,omitempty"`
	Type       string `json:"type,omitempty" bson:"type,omitempty"`
}

type GroupDiscount struct {
	Percentage   string       `json:"percentage,omitempty" bson:"percentage,omitempty"`
	SplitOptions SplitOptions `json:"split_options,omitempty" bson:"split_options,omitempty"`
}

type GroupMaxOrder struct {
	Volume   string        `json:"volume,omitempty" bson:"volume,omitempty"`
	Units    string        `json:"units,omitempty" bson:"units,omitempty"`
	Discount GroupDiscount `json:"discount,omitempty" bson:"group_discount,omitempty"`
}

type AcceptOverSubscriptions struct {
	Value                 string `json:"value,omitempty" bson:"value,omitempty"`
	PercentageOnMaxVolume string `json:"percentage_on_max_volume,omitempty" bson:"percentage_on_max_volume,omitempty"`
}

type AcceptUnderSubscriptions struct {
	Value                 string `json:"value,omitempty" bson:"value,omitempty"`
	PercentageOnMaxVolume string `json:"percentage_on_max_volume,omitempty" bson:"percentage_on_max_volume,omitempty"`
}

type Discount struct {
	Percentage string `json:"percentage,omitempty" bson:"percentage,omitempty"`
}

type AcceptSingleOrder struct {
	Value    string   `json:"value,omitempty" bson:"value,omitempty"`
	Discount Discount `json:"discount,omitempty" bson:"discount,omitempty"`
}

type Settings struct {
	MustSatisfyMaxVolume     string                   `json:"must_satisfy_max_volume,omitempty" bson:"must_satisfy_max_volume,omitempty"`
	AcceptOverSubscriptions  AcceptOverSubscriptions  `json:"accept_over_subscriptions,omitempty" bson:"accept_over_subscriptions,omitempty"`
	AcceptUnderSubscriptions AcceptUnderSubscriptions `json:"accept_under_subscriptions,omitempty" bson:"accept_under_subscriptions,omitempty"`
	AcceptSingleOrder        AcceptSingleOrder        `json:"accept_single_order,omitempty" bson:"accept_single_order,omitempty"`
}

type Group struct {
	IndividualMinOrder IndividualMinOrder `json:"individual_min_order,omitempty" bson:"individual_min_order,omitempty"`
	MaxOrder           GroupMaxOrder      `json:"max_order,omitempty" bson:"max_order,omitempty"`
	Settings           Settings           `json:"settings,omitempty" bson:"settings,omitempty"`
}

type OrderTypes struct {
	Individual Individual `json:"individual,omitempty" bson:"individual,omitempty"`
	Group      Group      `json:"group,omitempty" bson:"group,omitempty"`
}

type OrderTypesModel struct {
	OrderTypesId      primitive.ObjectID `json:"order_types_id,omitempty" bson:"_id,omitempty"`
	PreOrderRequestId string             `json:"pre_order_request_id,omitempty" bson:"pre_order_request_id,omitempty"`
	CustomerId        string             `json:"customer_id,omitempty" bson:"customer_id,omitempty"`
	ProductId         string             `json:"product_id,omitempty" bson:"product_id,omitempty"`
	OrderTypes        OrderTypes         `json:"order_types,omitempty" bson:"order_types,omitempty"`
}
