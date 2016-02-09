package dynamo

import (
	"errors"

	"github.com/aws/aws-sdk-go/aws/request"
	"github.com/aws/aws-sdk-go/service/dynamodb"
)

var (
	ErrNotImplemented = errors.New("not implemented")
)

type Dynamo struct{}

func (d *Dynamo) BatchGetItemRequest(*dynamodb.BatchGetItemInput) (*request.Request, *dynamodb.BatchGetItemOutput) {
	return nil, nil
}

func (d *Dynamo) BatchGetItem(*dynamodb.BatchGetItemInput) (*dynamodb.BatchGetItemOutput, error) {
	return nil, ErrNotImplemented
}

func (d *Dynamo) BatchGetItemPages(*dynamodb.BatchGetItemInput, func(*dynamodb.BatchGetItemOutput, bool) bool) error {
	return ErrNotImplemented
}

func (d *Dynamo) BatchWriteItemRequest(*dynamodb.BatchWriteItemInput) (*request.Request, *dynamodb.BatchWriteItemOutput) {
	return nil, nil
}

func (d *Dynamo) BatchWriteItem(*dynamodb.BatchWriteItemInput) (*dynamodb.BatchWriteItemOutput, error) {
	return nil, ErrNotImplemented
}

func (d *Dynamo) CreateTableRequest(*dynamodb.CreateTableInput) (*request.Request, *dynamodb.CreateTableOutput) {
	return nil, nil
}

func (d *Dynamo) CreateTable(*dynamodb.CreateTableInput) (*dynamodb.CreateTableOutput, error) {
	return nil, ErrNotImplemented
}

func (d *Dynamo) DeleteItemRequest(*dynamodb.DeleteItemInput) (*request.Request, *dynamodb.DeleteItemOutput) {
	return nil, nil
}

func (d *Dynamo) DeleteItem(*dynamodb.DeleteItemInput) (*dynamodb.DeleteItemOutput, error) {
	return nil, ErrNotImplemented
}

func (d *Dynamo) DeleteTableRequest(*dynamodb.DeleteTableInput) (*request.Request, *dynamodb.DeleteTableOutput) {
	return nil, nil
}

func (d *Dynamo) DeleteTable(*dynamodb.DeleteTableInput) (*dynamodb.DeleteTableOutput, error) {
	return nil, nil
}

func (d *Dynamo) DescribeTableRequest(*dynamodb.DescribeTableInput) (*request.Request, *dynamodb.DescribeTableOutput) {
	return nil, nil
}

func (d *Dynamo) DescribeTable(*dynamodb.DescribeTableInput) (*dynamodb.DescribeTableOutput, error) {
	return nil, ErrNotImplemented
}

func (d *Dynamo) GetItemRequest(*dynamodb.GetItemInput) (*request.Request, *dynamodb.GetItemOutput) {
	return nil, nil
}

func (d *Dynamo) GetItem(*dynamodb.GetItemInput) (*dynamodb.GetItemOutput, error) {
	return nil, ErrNotImplemented
}

func (d *Dynamo) ListTablesRequest(*dynamodb.ListTablesInput) (*request.Request, *dynamodb.ListTablesOutput) {
	return nil, nil
}

func (d *Dynamo) ListTables(*dynamodb.ListTablesInput) (*dynamodb.ListTablesOutput, error) {
	return nil, ErrNotImplemented
}

func (d *Dynamo) ListTablesPages(*dynamodb.ListTablesInput, func(*dynamodb.ListTablesOutput, bool) bool) error {
	return ErrNotImplemented
}

func (d *Dynamo) PutItemRequest(*dynamodb.PutItemInput) (*request.Request, *dynamodb.PutItemOutput) {
	return nil, nil
}

func (d *Dynamo) PutItem(*dynamodb.PutItemInput) (*dynamodb.PutItemOutput, error) {
	return nil, ErrNotImplemented
}

func (d *Dynamo) QueryRequest(*dynamodb.QueryInput) (*request.Request, *dynamodb.QueryOutput) {
	return nil, nil
}

func (d *Dynamo) Query(*dynamodb.QueryInput) (*dynamodb.QueryOutput, error) {
	return nil, ErrNotImplemented
}

func (d *Dynamo) QueryPages(*dynamodb.QueryInput, func(*dynamodb.QueryOutput, bool) bool) error {
	return ErrNotImplemented
}

func (d *Dynamo) ScanRequest(*dynamodb.ScanInput) (*request.Request, *dynamodb.ScanOutput) {
	return nil, nil
}

func (d *Dynamo) Scan(*dynamodb.ScanInput) (*dynamodb.ScanOutput, error) {
	return nil, ErrNotImplemented
}

func (d *Dynamo) ScanPages(*dynamodb.ScanInput, func(*dynamodb.ScanOutput, bool) bool) error {
	return ErrNotImplemented
}

func (d *Dynamo) UpdateItemRequest(*dynamodb.UpdateItemInput) (*request.Request, *dynamodb.UpdateItemOutput) {
	return nil, nil
}

func (d *Dynamo) UpdateItem(*dynamodb.UpdateItemInput) (*dynamodb.UpdateItemOutput, error) {
	return nil, ErrNotImplemented
}

func (d *Dynamo) UpdateTableRequest(*dynamodb.UpdateTableInput) (*request.Request, *dynamodb.UpdateTableOutput) {
	return nil, nil
}

func (d *Dynamo) UpdateTable(*dynamodb.UpdateTableInput) (*dynamodb.UpdateTableOutput, error) {
	return nil, ErrNotImplemented
}
