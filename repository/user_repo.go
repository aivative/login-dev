package repository

import (
	"context"
	"fmt"

	"github.com/aivative/login-dev/config"
	"github.com/aivative/login-dev/repository/model"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

func NewUserRepo(client *mongo.Client) *UserMongoRepo {
	mongoConfig := config.MongoConf["mongo-user"]

	collection := client.Database(mongoConfig.DBName).Collection(mongoConfig.CollName)
	return &UserMongoRepo{
		coll: collection,
	}
}

type UserMongoRepo struct {
	coll *mongo.Collection
}

func (umr *UserMongoRepo) GetUserLoginInfo(ctx context.Context, loginReq model.TLoginReq, userResult *model.TUser) (err error) {
	if err = umr.coll.FindOne(ctx, gin.H{"email": loginReq.Email}).
		Decode(userResult); err != nil {
		return

	} else if err != nil {
		return
	}
	return
}

func (umr *UserMongoRepo) CreateUser(ctx context.Context, userPayload model.TCreateUserQuery) error {
	ir, err := umr.coll.InsertOne(ctx, userPayload)
	if err != nil {
		return err
	}

	if ir.InsertedID == primitive.NilObjectID {
		return fmt.Errorf("can't insert user")
	}

	return nil
}

func (umr *UserMongoRepo) GetUser(ctx context.Context, uid string) (result model.TGetUserResp, err error) {
	err = umr.coll.FindOne(ctx, model.TUserID{UserID: uid}).Decode(&result)
	return
}

func (umr *UserMongoRepo) DeleteUser(ctx context.Context, uid string) error {
	dr, err := umr.coll.DeleteOne(ctx, model.TUserID{UserID: uid})
	if err != nil {
		return err
	}

	if dr.DeletedCount == 0 {
		return fmt.Errorf("can't delete user")
	}

	return nil
}

func (umr *UserMongoRepo) UpdateUser(ctx context.Context, uid string, profile model.TUpdateUserReq) error {
	ur, err := umr.coll.UpdateOne(ctx, model.TUserID{UserID: uid}, profile)
	if err != nil {
		return err
	}

	if ur.MatchedCount == 0 {
		return fmt.Errorf("user not found")
	}

	nilStruct := model.TUpdateUserReq{}

	if ur.ModifiedCount == 0 && profile != nilStruct {
		return fmt.Errorf("user not updated")
	}

	return nil
}

// // GetUsers
// func (r *UserMongoRepo) GetUsers(offset int64, result interface{}) error {
// 	var opt options.FindOptions
// 	opt.SetLimit(10).SetSkip(offset * 10)
//
// 	cursor, err := r.coll.Find(r.ctx, bson.M{"user_type": bson.M{"$ne": "ADMIN"}}, &opt)
// 	if err != nil {
// 		return err
// 	}
//
// 	if cursor.RemainingBatchLength() == 0 {
// 		return mongo.ErrNoDocuments
// 	}
//
// 	if err := cursor.All(r.ctx, result); err != nil {
// 		return err
// 	}
//
// 	return nil
// }
//
// // GetUsersByDistrictID
// func (r *UserMongoRepo) GetUsersByDistrictID(districtID string, result interface{}) error {
// 	cursor, err := r.coll.Find(r.ctx, bson.M{"district_id": districtID, "user_type": "GARBO"})
// 	if err != nil {
// 		return err
// 	}
//
// 	if cursor.RemainingBatchLength() == 0 {
// 		return mongo.ErrNoDocuments
// 	}
//
// 	if err := cursor.All(r.ctx, result); err != nil {
// 		return err
// 	}
//
// 	return nil
// }
