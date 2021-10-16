package main

import (
	"context"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/dynamodb"
	"time"
)

var db = dynamodb.New(session.Must(session.NewSession()))

func Update(ctx context.Context) {

	_, err := db.UpdateItemWithContext(ctx, &dynamodb.UpdateItemInput{
		ExpressionAttributeNames: map[string]*string{
			"#AT":  aws.String("AlbumTitle"),
			"#Y":   aws.String("Year"),
			"#REV": aws.String("Revision"),
			"#UPA": aws.String("UpdatedAt"),
		},
		ExpressionAttributeValues: map[string]*dynamodb.AttributeValue{
			":t": {
				S: aws.String("Louder Than Ever"),
			},
			":y": {
				N: aws.String("2015"),
			},
			":inc": {
				N: aws.String("1"),
			},
			":upa": {
				S: aws.String(time.Now().UTC().Format(time.RFC3339)),
			},
		},
		Key: map[string]*dynamodb.AttributeValue{
			"Artist": {
				S: aws.String("Acme Band"),
			},
			"SongTitle": {
				S: aws.String("Happy Day"),
			},
		},
		TableName:        aws.String("Music"),
		UpdateExpression: aws.String("SET #Y = :y, #AT = :t, #UPA = :upa ADD #REV :inc"),
	})

	if err != nil {
		// エラーハンドリング
	}

}
