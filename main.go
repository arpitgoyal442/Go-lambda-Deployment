package main

import (

	"fmt"
	"github.com/aws/aws-lambda-go/lambda"
	
	// "github.com/aws/aws-lambda-go/lambda"
)

type MyEvent struct{

	Name string `json:"name"`
	// Age  int  `json:"What is your age"` 
}

type MyResponse struct{

	Message string `json:"Answer:"`
}

func HandleLambdaEvent(event MyEvent) (MyResponse,error){

	return MyResponse{Message: fmt.Sprintf( "Hello %s ",event.Name)},nil



}

func main(){

	lambda.Start(HandleLambdaEvent)



}