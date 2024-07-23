package models

type LayerStyleRelation struct {
	ID      string `json:"id"`
	StyleID string `json:"style_id"`
	LayerID string `json:"layer_id"`
}
