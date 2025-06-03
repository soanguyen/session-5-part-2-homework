package mongostore

import (
	"context"
	"ct-backend-course-baonguyen/internal/entity"
	"ct-backend-course-baonguyen/pkg/hashpass"
	"go.mongodb.org/mongo-driver/mongo"
	"time"
)

func NewUserCollection(db *mongo.Database, collName string) *userCollection {
	return &userCollection{
		client:  db.Collection(collName),
		timeout: 3 * time.Second,
	}
}

type userCollection struct {
	client  *mongo.Collection
	timeout time.Duration
}

func (c *userCollection) Create(info entity.UserInfo) error {
	doc := NewUserDocument(info)

	ctx, cancelFn := context.WithTimeout(context.Background(), c.timeout)
	defer cancelFn()

	if _, err := c.client.InsertOne(ctx, doc); err != nil {
		return err
	}

	return nil
}

func (c *userCollection) ChangePassword(username string, newPassword string) error {
	// #TODO 3: implement ChangePassword
	// oldPass and newPassword should not be duplicate

	return nil
}

func IsPasswordDuplicate(newPassword, hashPassword string) bool {
	return hashPassword == hashpass.HashPasswordLogin(newPassword, hashPassword)
}

func (c *userCollection) Query(username string) (entity.UserInfo, error) {
	panic("TODO implement it")

	//ctx, cancelFn := context.WithTimeout(context.Background(), c.timeout)
	//defer cancelFn()
	//
	//panic("TODO implement it")
	//
	//return  entity.UserInfo{}, nil
}
