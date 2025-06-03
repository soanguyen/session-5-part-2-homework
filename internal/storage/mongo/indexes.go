package mongostore

import (
	"context"
	"go.mongodb.org/mongo-driver/mongo"
)

func CreateUserCollectionIndex(userColl *userCollection) error {
	c := userColl.client

	// TODO 2: Create unique index for 'username'

	indexModel := mongo.IndexModel{}
	_, err := c.Indexes().CreateOne(context.TODO(), indexModel)
	if err != nil {
		return err
	}

	return nil
}
