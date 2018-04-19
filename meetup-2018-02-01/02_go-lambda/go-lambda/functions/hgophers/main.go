// package main

// import (
// 	"context"
// 	"fmt"

// 	"github.com/aws/aws-lambda-go/lambda"
// )

// type MyEvent struct {
// 	Name string `json:"name"`
// }

// func HandleRequest(ctx context.Context, name MyEvent) (string, error) {
// 	return fmt.Sprintf("Hello %s! 💩", name.Name), nil
// }

// func main() {
// 	lambda.Start(HandleRequest)
// }

// echo '{"name":"Gophers"}' | apex --profile=hgophers invoke hgophers

////////

package main

import (
	"fmt"

	"github.com/aws/aws-lambda-go/lambda"
)

// type MyEvent struct {
// 	Name string `json:"name"`
// 	Age  int    `json:"age"`
// }

// type MyResponse struct {
// 	Message string `xml:"answer"`
// }

func HandleLambdaEvent(event interface{}) (string, error) {

	// return MyResponse{Message: fmt.Sprintf("%s is %d years old!", event.Name, event.Age)}, nil
	i, ok := event.(map[string]interface{})
	if !ok {
		return "", nil
	}

	return fmt.Sprintf("%s", i["name"]), nil
}

func main() {
	lambda.Start(HandleLambdaEvent)
}

// echo '{"name":"Gophers","age":10}' | apex --profile=hgophers invoke hgophers

// echo '{"name":"{{name}}","age":10}' | phony --tick 1s | apex --profile=hgophers invoke hgophers
