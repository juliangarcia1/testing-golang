package dynamo

import (
	"awesomeProject/service"
	"context"
	"github.com/aws/aws-sdk-go-v2/feature/dynamodb/attributevalue"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb"
	"log"

	"github.com/aws/aws-sdk-go-v2/aws"
)

type Keys struct {
	Id   string `json:"id" dynamodbav:"id"`
	Name string `json:"name" dynamodbav:"name"`
}

func NewDB(configfile any) service.DBService {
	//Config aws dynamo and return a dynamo client
}

type DBClient struct {
	client any
}

func (c DBClient) Get(ctx context.Context, tableName string, keys Keys) (service.DBService, error) {
	av := attributevalue.UnmarshallMap(keys)

	input := &dynamodb.GetItemInput{
		TableName: aws.String(tableName),
		Key:       av,
	}
	result, err := c.client.GetItem(ctx, input)
	if err != nil {
		log.Println("Error reading dynamo table.")
		return nill, err
	}
	return result, nil
}
