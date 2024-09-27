package database

type DynamoDbClient struct {
	databaseStore string
}

func NewDynamoDbClient() DynamoDbClient {
	return DynamoDbClient{
		databaseStore: "dynamodb",
	}
}