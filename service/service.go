package service

import "context"

type DBService interface {
	Get(ctx context.Context, tableName, Keys any) (any, error)
}
