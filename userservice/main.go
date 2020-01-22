package main

import (
	"fmt"

	"github.com/DataInfosec/faceappapi/userservice/utils/router"
)

// @title Swagger Example API
// @version 1.0
// @description This is a mini store application.
// @termsOfService http://swagger.io/terms/

// @contact.name API Support
// @contact.url http://www.swagger.io/support
// @contact.email support@swagger.io

// @license.name Apache 2.0
// @license.url http://www.apache.org/licenses/LICENSE-2.0.html

// @securityDefinitions.basic BasicAuth

// @securityDefinitions.apikey ApiKeyAuth
// @in header
// @name Authorization

func main() {
	// ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	// client, err := mongo.Connect(ctx, options.Client().ApplyURI("mongodb://localhost:27017"))
	// if err != nil {
	// 	log.Fatal(err)
	// }
	fmt.Println("auth starting application")
	// bucket = connection.Connection()
	router.Router()
	// fmt.Println("Starting Auth services")
	// resp, err1 := http.Get("https://tanda102.herokuapp.com/account/login?username=dirisujesse&password=qwerty&status=admin")
	// if err1 != nil {
	// 	log.Fatalln(err1)
	// }
	// log.Println("error in request 11 :: ", err1)
	// log.Println("resp in request 222 :: ", resp.Status)

	// body, err := ioutil.ReadAll(resp.Body)
	// if err != nil {
	// 	log.Fatalln(err)
	// }
	// log.Println("error in request :: ", err)
	// log.Println("request body :: ", string(body))

	// conn, err := grpc.Dial(":50052", grpc.WithInsecure())
	// if err != nil {
	// 	panic(err)
	// }
	// client := proto.NewClockServiceClient(conn)

	// req := &proto.ClockinRequest{Id: "eerrte", Email: "string@gmail.com"}
	// fmt.Println("req body :: ", req)
	// response, err := client.Clockin(context.Background(), req)
	// if err == nil {
	// 	// ctx.JSON(http.StatusOK, gin.H{
	// 	fmt.Print("response :: ", response)
	// 	// })
	// } else {
	// 	fmt.Print("error :: ", response)
	// }
}
