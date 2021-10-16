package main

import (
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/dynamodb"
	"github.com/aws/aws-sdk-go/service/dynamodb/expression"
)

func main() {

	filt := expression.Name("Artist").Equal(expression.Value("No One You Know"))
	proj := expression.NamesList(expression.Name("SongTitle"), expression.Name("AlbumTitle"))
	expr, err := expression.NewBuilder().WithFilter(filt).WithProjection(proj).Build()
	if err != nil {
		// エラーハンドリング
	}

	db := dynamodb.New(session.Must(session.NewSession()))
	result, err := db.Scan(&dynamodb.ScanInput{
		ExpressionAttributeNames:  expr.Names(),
		ExpressionAttributeValues: expr.Values(),
		FilterExpression:          expr.Filter(),
		ProjectionExpression:      expr.Projection(),
		TableName:                 aws.String("Music"),
	})
	if err != nil {
		// エラーハンドリング
	}

}
