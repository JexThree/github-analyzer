package models

import "go.mongodb.org/mongo-driver/v2/bson"

type User struct {
	ID bson.ObjectID `bson:"_id,omitempty" json:"id"`

	Username string `bson:"username" json:"username"`

	Password string `bson:"password" json:"password"`

	Name string `bson:"name" json:"name"`

	LastName string `bson:"lastname" json:"lastname"`

	GithubEmail string `bson:"githubEmail" json:"githubEmail"`

	BirthDate string `bson:"birthDate" json:"birthDate"`
}
