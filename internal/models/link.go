package models

type Link struct {
	Short    string `bson:"_id" json:"short,omitempty"`
	Original string `bson:"original" json:"original,omitempty"`
}
