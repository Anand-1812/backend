package auth

import (
	"context"
	"errors"
	"time"

	"go.mongodb.org/mongo-driver/v2/bson"
	"go.mongodb.org/mongo-driver/v2/mongo"
)

var ErrUserNotFound = errors.New("user not found")

type UserRepository interface {
	Create(ctx context.Context, user *User) error
	GetByEmail(ctx context.Context, email string) (*User, error)
}

type MongoUserRepository struct {
	collection *mongo.Collection
}

func NewMongoUserRepository(client *mongo.Client) *MongoUserRepository {
	return &MongoUserRepository{
		collection: client.Database("db").Collection("users"),
	}
}

func (r *MongoUserRepository) Create(ctx context.Context, user *User) error {
	user.CreatedAt = time.Now().UTC()

	result, err := r.collection.InsertOne(ctx, user)
	if err != nil {
		return err
	}

	if id, ok := result.InsertedID.(bson.ObjectID); ok {
		user.ID = id
	}

	return nil
}

func (r *MongoUserRepository) GetByEmail(ctx context.Context, email string) (*User, error) {
	var user User

	err := r.collection.FindOne(ctx, bson.M{"email": email}).Decode(&user)
	if errors.Is(err, mongo.ErrNoDocuments) {
		return nil, ErrUserNotFound
	}
	if err != nil {
		return nil, err
	}

	return &user, nil
}
