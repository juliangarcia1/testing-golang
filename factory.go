package main

import (
	"awesomeProject/newpackage/dynamo"
	"awesomeProject/service"
)

type ServiceType int

const (
	DynamoSvc = iota
)

func DBServiceFactory(svc ServiceType, svcConfig any) service.DBService {

	switch svc {
	case DynamoSvc:
		return dynamo.NewDB()
	}

}
