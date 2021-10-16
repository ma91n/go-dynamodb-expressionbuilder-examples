package main

import (
	"context"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/dynamodb"
	"github.com/aws/aws-sdk-go/service/dynamodb/expression"
	"time"
)

var db = dynamodb.New(session.Must(session.NewSession()))

func Update(ctx context.Context) {

	update := expression.Set(expression.Name("AlbumTitle"), expression.Value("Louder Than Ever")).
		Set(expression.Name("Year"), expression.Value("2015")).
		Set(expression.Name("UpdatedAt"), expression.Value(time.Now())).
		Add(expression.Name("Revision"), expression.Value(1))

	expr, err := expression.NewBuilder().WithUpdate(update).Build()
	if err != nil {
		// エラーハンドリング
	}

	_, err := db.UpdateItemWithContext(ctx, &dynamodb.UpdateItemInput{
		TableName: aws.String("Music"),
		Key: map[string]*dynamodb.AttributeValue{
			"Artist": {
				S: aws.String("Acme Band"),
			},
			"SongTitle": {
				S: aws.String("Happy Day"),
			},
		},
		ExpressionAttributeNames:  expr.Names(),
		ExpressionAttributeValues: expr.Values(),
		UpdateExpression:          expr.Update(),
	})

	if err != nil {
		// エラーハンドリング
	}

}
