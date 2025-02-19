package database

import (
	"lambda-func/types"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/dynamodb"
)

const (
	tableName = "userTable"
)
type DynamoDbClient struct {
	databaseStore *dynamodb.DynamoDB
}

func NewDynamoDbClient() DynamoDbClient {
	dbSession := session.Must(session.NewSession())
	db := dynamodb.New(dbSession)
	return DynamoDbClient{
		databaseStore: db,
	}
}

func (u DynamoDbClient) DoesUserExist(username string) (bool, error) {
	result, err := u.databaseStore.GetItem(&dynamodb.GetItemInput{
		TableName: aws.String(tableName),
		Key: map[string]*dynamodb.AttributeValue{
			"username": {
				S: aws.String(username),
			},
		},
	})

	if err != nil {
		return true, err
	}
	if result.Item == nil {
		return false, nil
	}
	return true, nil
}


func (u DynamoDbClient) CreateUser(user types.RegisterUser) error {

	item := &dynamodb.PutItemInput{
		TableName: aws.String(tableName),
		Item: map[string]*dynamodb.AttributeValue{
			"username": {
				S: aws.String(user.Username),
			},
			"email": {
				S: aws.String(user.Password),
			},
		},
	}
	_, err := u.databaseStore.PutItem(item)
	if err != nil {
		return err
	}
	return nil  
}