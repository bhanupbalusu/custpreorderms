package deal_settings

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type AcceptOverSubscriptions struct {
	AOSValue                          string `json:"aos_value,omitempty" bson:"aos_value,omitempty"`
	AOSPercentageOnBulkQuantityVolume string `json:"aos_percentage_on_bulk_quantity_volume,omitempty" bson:"aos_percentage_on_bulk_quantity_volume,omitempty"`
}

type AcceptUnderSubscriptions struct {
	AUSValue                          string `json:"aus_value,omitempty" bson:"value,omitempty"`
	AUSPercentageOnBulkQuantityVolume string `json:"aus_percentage_on_bulk_quantity_volume,omitempty" bson:"aus_percentage_on_bulk_quantity_volume,omitempty"`
}

type ByPost struct {
	BPValue string `json:"by_post_value,omitempty" bson:"value,omitempty"`
	Charges string `json:"charges,omitempty" bson:"charges,omitempty"`
}

type DeliveryOptions struct {
	Collection string `json:"collection,omitempty" bson:"collection,omitempty"`
	ByPost     ByPost `json:"by_post,omitempty" bson:"by_post,omitempty"`
}

type PaymentOptions struct {
	Cash    string `json:"cash,omitempty" bson:"cash,omitempty"`
	Digital string `json:"digital,omitempty" bson:"digital,omitempty"`
}

type DealSettings struct {
	MustSatisfyBulkQuantityVolume string                   `json:"must_satisfy_bulk_quantity_volume,omitempty" bson:"must_satisfy_bulk_quantity_volume,omitempty"`
	AcceptOverSubscriptions       AcceptOverSubscriptions  `json:"accept_over_subscriptions,omitempty" bson:"accept_over_subscriptions,omitempty"`
	AcceptUnderSubscriptions      AcceptUnderSubscriptions `json:"accept_under_subscriptions,omitempty" bson:"accept_under_subscriptions,omitempty"`
	DeliveryOptions               DeliveryOptions          `json:"delivery_options,omitempty" bson:"delivery_options,omitempty"`
	PaymentOptions                PaymentOptions           `json:"payment_options,omitempty" bson:"payment_options,omitempty"`
}

type DealSettingsModel struct {
	DealSettingsId    primitive.ObjectID `json:"deal_settings_id,omitempty" bson:"_id,omitempty"`
	PreOrderRequestId string             `json:"pre_order_request_id,omitempty" bson:"pre_order_request_id,omitempty"`
	CustomerId        string             `json:"customer_id,omitempty" bson:"customer_id,omitempty"`
	ProductId         string             `json:"product_id,omitempty" bson:"product_id,omitempty"`
	OrderTypesId      string             `json:"order_types_id,omitempty" bson:"order_types_id,omitempty"`
	DealSettings      DealSettings       `json:"deal_settings,omitempty" bson:"deal_settings,omitempty"`
}
