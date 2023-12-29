package main

import (
	"awesomeProject/newpackage/dynamo"
	"context"
	"fmt"
	//"testing-golang/service"
)

func main() {
	fmt.Println("Hello")
	svc := DBServiceFactory(DynamoSvc)
	svc.Get(context.TODO(), "test0", dynamo.Keys{Id: "0", Name: "john"})

}
