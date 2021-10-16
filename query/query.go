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

func QueryTable(ctx context.Context, deviceID string, start, end time.Time) {

	keyCond := expression.Key("DeviceID").Equal(expression.Value(deviceID)).
		And(expression.Key("Timestamp").Between(expression.Value(start.Format(time.RFC3339)), expression.Value(end.Format(time.RFC3339))))

	filterCond := expression.Name("DeviceType").Equal(expression.Value("Normal")).
		And(expression.Name("CreatedYear").Equal(expression.Value(2021)))

	expr, err := expression.NewBuilder().WithKeyCondition(keyCond).WithFilter(filterCond).Build()
	if err != nil {
		// エラーハンドリング
	}

	result, err := db.QueryWithContext(ctx, &dynamodb.QueryInput{
		KeyConditionExpression:    expr.KeyCondition(),
		ProjectionExpression:      expr.Projection(),
		ExpressionAttributeNames:  expr.Names(),
		ExpressionAttributeValues: expr.Values(),
		FilterExpression:          expr.Filter(),
		TableName:                 aws.String("DeviceLog"),
	})
	if err != nil {
		// エラーハンドリング
	}

}
