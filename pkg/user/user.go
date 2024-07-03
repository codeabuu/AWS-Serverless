package user

import (
	"github.com/aws/aws-sdk-go/service/dynamodb"
	"github.com/aws/aws-sdk-go/service/dynamodb/dynamodbiface"
)

var ()

type User struct {
	Email     string `json:"email"`
	FirstName string `json:"firstName"`
	LastName  string `json:"lastName"`
}

func FetchUser(email, tableName string, dynaClient dynamodbiface.DynamoDBAPI) (*User){
	input := &dynamodb.GetItemInput{
		Key: map[string]*dynamodb.AttributeValue{
			
		}
	}
}

func FetchUsers() {

}

func CreateUser() {

}

func UpdateUser() {

}

func DeleteUser() error {

}