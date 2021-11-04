package models

type Role struct {
	Id   uint64 `json:"id,string,omitempty"`
	Name string `json:"name"`
}
