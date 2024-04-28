package repositories

import (
	"context"
	"makino/models"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type PersonInfoRepository interface {
	Get() ([]models.PersonInfo, error)
	GetCustomerByPhoneNumber(filter bson.D) *models.Telephone
	GetCustomerByIdCard(idCard string) *models.PersonInfo
}

type mongoDbPersonInfoRepository struct {
	ctx      *context.Context
	database *mongo.Database
}

func NewPersonInfoRepository(ctx *context.Context, database *mongo.Database) PersonInfoRepository {
	return &mongoDbPersonInfoRepository{ctx: ctx, database: database}
}

func (repo *mongoDbPersonInfoRepository) Get() ([]models.PersonInfo, error) {
	filter := bson.D{}
	ctx := *repo.ctx
	collection := repo.database.Collection("person_info")
	cursor, err := collection.Find(ctx, filter)
	if err != nil {
		panic(err)
	}

	var persons []models.PersonInfo

	for cursor.Next(ctx) {
		var person models.PersonInfo
		err := cursor.Decode(&person)
		if err != nil {
			panic(err)
		}
		persons = append(persons, person)
	}
	defer cursor.Close(ctx)
	return persons, err
}

func (repo *mongoDbPersonInfoRepository) GetCustomerByPhoneNumber(filter bson.D) *models.Telephone {
	// var personInfo []models.PersonInfo
	var telephone *models.Telephone

	telephoneCollection := repo.database.Collection("telephone")
	// personInfoCollection := repo.database.Collection("person_info")

	cursor, err := telephoneCollection.Find(*repo.ctx, filter)
	if err != nil {
		panic(err)
	}
	for cursor.Next(*repo.ctx) {
		err := cursor.Decode(&telephone)
		if err != nil {
			return nil
		}
	}
	return telephone
}

func (repo *mongoDbPersonInfoRepository) GetCustomerByIdCard(idCard string) *models.PersonInfo {
	var personInfo *models.PersonInfo

	personInfoCollection := repo.database.Collection("person_info")

	filter := bson.D{{"idCard", idCard}}
	cursor, err := personInfoCollection.Find(*repo.ctx, filter)
	if err != nil {
		panic(err)
	}
	for cursor.Next(*repo.ctx) {
		err := cursor.Decode(&personInfo)
		if err != nil {
			return nil
		}
	}
	return personInfo
}
