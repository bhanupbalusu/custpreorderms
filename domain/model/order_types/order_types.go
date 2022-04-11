package order_types

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type IMinOrder struct {
	IMinOrderVolume string `json:"min_order_volume,omitempty" bson:"min_order_volume,omitempty"`
	IMinUnits       string `json:"min_order_units,omitempty" bson:"min_order_units,omitempty"`
}

type IMaxOrder struct {
	IMaxOrderVolume string `json:"max_order_volume,omitempty" bson:"max_order_volume,omitempty"`
	IMaxUnits       string `json:"max_order_units,omitempty" bson:"max_order_units,omitempty"`
}

type OTIndividual struct {
	IMinOrder IMinOrder `json:"i_min_order,omitempty" bson:"i_min_order,omitempty"`
	IMaxOrder IMaxOrder `json:"i_max_order,omitempty" bson:"i_max_order,omitempty"`
}

type GIndividualMinOrder struct {
	IndividualMinOrderVolume string `json:"individual_min_order_volume,omitempty" bson:"individual_min_order_volume,omitempty"`
	IndividualMinOrderUnits  string `json:"individual_min_order_units,omitempty" bson:"individual_min_order_units,omitempty"`
}

type GMODSplitOptions struct {
	SOGroupOwner string `json:"so_group_owner,omitempty" bson:"so_group_owner,omitempty"`
	SOType       string `json:"so_type,omitempty" bson:"so_type,omitempty"`
}

type GroupMaxOrderDiscount struct {
	GMODPercentage   string           `json:"gmod_percentage,omitempty" bson:"gmod_percentage,omitempty"`
	GMODSplitOptions GMODSplitOptions `json:"gmod_split_options,omitempty" bson:"gmod_split_options,omitempty"`
}

type GGroupMaxOrder struct {
	GroupMaxOrderVolume   string                `json:"group_max_order_volume,omitempty" bson:"group_max_order_volume,omitempty"`
	GroupMaxOrderUnits    string                `json:"group_max_order_units,omitempty" bson:"group_max_order_units,omitempty"`
	GroupMaxOrderDiscount GroupMaxOrderDiscount `json:"group_max_order_discount,omitempty" bson:"group_max_order_discount,omitempty"`
}

type AcceptOverSubscriptions struct {
	AOSValue                 string `json:"aos_value,omitempty" bson:"aos_value,omitempty"`
	AOSPercentageOnMaxVolume string `json:"aos_percentage_on_max_volume,omitempty" bson:"aos_percentage_on_max_volume,omitempty"`
}

type AcceptUnderSubscriptions struct {
	AUSValue                 string `json:"aus_value,omitempty" bson:"aus_value,omitempty"`
	AUSPercentageOnMaxVolume string `json:"aus_percentage_on_max_volume,omitempty" bson:"aus_percentage_on_max_volume,omitempty"`
}

type ASOSingleOrderDiscount struct {
	ASOSODPercentage string `json:"sod_percentage,omitempty" bson:"sod_percentage,omitempty"`
}

type AcceptSingleOrder struct {
	ASOValue               string                 `json:"aso_value,omitempty" bson:"aso_value,omitempty"`
	ASOSingleOrderDiscount ASOSingleOrderDiscount `json:"aso_single_order_discount,omitempty" bson:"aso_single_order_discount,omitempty"`
}

type GSettings struct {
	MustSatisfyMaxVolume     string                   `json:"s_must_satisfy_max_volume,omitempty" bson:"s_must_satisfy_max_volume,omitempty"`
	AcceptOverSubscriptions  AcceptOverSubscriptions  `json:"s_accept_over_subscriptions,omitempty" bson:"s_accept_over_subscriptions,omitempty"`
	AcceptUnderSubscriptions AcceptUnderSubscriptions `json:"s_accept_under_subscriptions,omitempty" bson:"s_accept_under_subscriptions,omitempty"`
	AcceptSingleOrder        AcceptSingleOrder        `json:"s_accept_single_order,omitempty" bson:"s_accept_single_order,omitempty"`
}

type OTGroup struct {
	GIndividualMinOrder GIndividualMinOrder `json:"g_individual_min_order,omitempty" bson:"g_individual_min_order,omitempty"`
	GGroupMaxOrder      GGroupMaxOrder      `json:"g_max_order,omitempty" bson:"g_max_order,omitempty"`
	GSettings           GSettings           `json:"g_settings,omitempty" bson:"g_settings,omitempty"`
}

type OrderTypes struct {
	OTIndividual OTIndividual `json:"ot_individual,omitempty" bson:"ot_individual,omitempty"`
	OTGroup      OTGroup      `json:"ot_group,omitempty" bson:"ot_group,omitempty"`
}

type OrderTypesModel struct {
	OrderTypesId primitive.ObjectID `json:"order_types_id,omitempty" bson:"_id,omitempty"`
	PreOrderId   string             `json:"pre_order_id,omitempty" bson:"pre_order_request_id,omitempty"`
	CustomerId   string             `json:"customer_id,omitempty" bson:"customer_id,omitempty"`
	ProductId    string             `json:"product_id,omitempty" bson:"product_id,omitempty"`
	OrderTypes   OrderTypes         `json:"order_types,omitempty" bson:"order_types,omitempty"`
	CreatedAt    time.Time          `json:"created_at" bson:"created_at"`
	UpdatedAt    time.Time          `json:"updated_at" bson:"updated_at"`
}
