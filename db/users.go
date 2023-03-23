//   Approver Bot
//   Copyright (C) 2022 Reeshuxd (@reeshuxd)

//   This program is distributed in the hope that it will be useful,
//   but WITHOUT ANY WARRANTY; without even the implied warranty of
//   MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
//   GNU Affero General Public License for more details.

package db

import (
	"context"
	"fmt"

	"go.mongodb.org/mongo-driver/bson"
)

var (
	udb = database.Collection("Users")
)

func CheckUser(user_id int64) bool {
	find := udb.FindOne(context.TODO(), bson.M{"user_id": user_id})
	if find.Err() == nil {
		return true
	} else {
		return false
	}
}

func AddUser(user_id int64) bool {
	check := CheckUser(user_id)
	if !check {
		_, err := udb.InsertOne(
			context.TODO(),
			bson.M{"user_id": user_id},
		)
		if err != nil {
			fmt.Println("Error while adding user:" + err.Error())
		}
	}
	return true
}

func GetUsers() []bson.M {
	var users []bson.M
	u, _ := udb.Find(context.TODO(), bson.M{})
	u.All(context.TODO(), &users)
	return users
}
