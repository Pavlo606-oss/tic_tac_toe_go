package models

type Board struct {
	Condition [3][3]int8 `json:"condition"`
}
