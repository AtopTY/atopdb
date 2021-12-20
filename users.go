package atopdb

import (
	"context"
	"fmt"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

// UserCollection collection name of users
var UserCollection = "users"

// type UserRequest struct {
// 	Name     string   `bson:"name"`
// 	Password string   `bson:"password"`
// 	Tags     []string `bson:"tags"`
// }

//ValidUser check name and password
func (d *DB) ValidUser(body bson.M) (string, error) {

	ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
	defer cancel()
	name := body["name"]
	colleciont := d.db.Collection(UserCollection)
	var item bson.M
	if err := colleciont.FindOne(ctx, bson.M{"name": name}).Decode(&item); err != nil {
		return "", err
	}

	password := []byte(body["password"].(string))
	body["password"].sha256.New().Sum(password)

	if item["password"] == body["password"] {
		id := item["_id"].(primitive.ObjectID)
		return id.Hex(), nil
	}
	return "", fmt.Errorf("Wrong account or password")
}

type UserRole struct {
	Name string `bson:"name"`
	Role string `bson:"role"`
}

// GetRoleByID  get user role
func (d *DB) GetRoleByID(idHex string) (UserRole, error) {
	id, err := primitive.ObjectIDFromHex(idHex)
	if err != nil {
		return UserRole{}, err
	}
	ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
	defer cancel()
	colleciont := d.db.Collection(UserCollection)

	var item UserRole
	if err := colleciont.FindOne(ctx, bson.M{"_id": id}).Decode(&item); err != nil {
		return UserRole{}, err
	}
	// role, ok := item["role"].(string)
	// if !ok {
	// 	return "", errors.New("Missing role property")
	// }
	return item, nil
}
