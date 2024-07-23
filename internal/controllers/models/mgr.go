package models

type MapGroupRelation struct {
	ID         string `json:"id"`
	GroupID    string `json:"group_id"`
	MapID      string `json:"map_id"`
	GroupOrder int32  `json:"group_order"`
}
