package main

import (
        "fmt"
        "github.com/aws/aws-lambda-go/lambda"
)
//AWS lamda code  
type employee struct {
		ID float64 `json:"name"`
		Value string `json:"value"`
}

type Response struct {
		Message string `json:"message"`

}
//AWS lamda handler function to get a record from db.go (dynomoDB) and send a response
func myfirst() (Response, error){
	
	record, err := getItem(2211)
    if err != nil {
        return Response {Message:"Error"}, err
    }
	return Response {
			Message: fmt.Sprintf("Employee %f : %s!", record.ID, record.Value )}, nil
		
}

func main() {
        lambda.Start(myfirst)
}