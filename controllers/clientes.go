package controllers

import (
	"context"
	"io"
	"log"
	"net/http"
	"strconv"

	"users/grpc/client/models"
	"users/grpc/client/pb"

	"github.com/gin-gonic/gin"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func GetAllUsers(ctext *gin.Context) {
	var users []models.Users
	conn, err := grpc.Dial("localhost:9000", grpc.WithTransportCredentials(insecure.NewCredentials()))

	if err != nil {
		log.Fatal(err)
	}

	client := pb.NewUserServerClient(conn)
	req := &pb.Empty{}

	res, err := client.GetAllUsers(context.Background(), req)
	if err != nil {
		log.Fatal(err)
	}

	for {
		row, err := res.Recv()
		if err == io.EOF {
			break
		}
		if err != nil {
			log.Fatalf("%v.GetAllUsers(_) = _, %v", client, err)
		}
		users = append(users, models.Users{Id: int(row.Id), Username: row.Username, Passwd: row.Passwd, Email: row.Email})
	}
	ctext.JSON(200, users)
}

func GetUserById(ctext *gin.Context) {
	var user models.Users
	conn, err := grpc.Dial("localhost:9000", grpc.WithTransportCredentials(insecure.NewCredentials()))
	id := ctext.Params.ByName("id")
	idInt64, _ := strconv.ParseInt(id, 10, 64)

	if err != nil {
		log.Fatal(err)
	}

	client := pb.NewUserServerClient(conn)
	req := &pb.Id{Id: idInt64}

	res, err := client.GetUserById(context.Background(), req)
	if err != nil {
		log.Fatal(err)
	}
	user = models.Users{Id: int(res.Id), Username: res.Username, Passwd: res.Passwd, Email: res.Email}

	if user.Id == 0 {
		ctext.JSON(http.StatusNotFound, gin.H{
			"Message": "User not Found"})
		return
	}
	ctext.JSON(http.StatusOK, user)
}

func InsertUser(ctext *gin.Context) {
	var user models.Users

	if error := ctext.ShouldBindJSON(&user); error != nil {
		ctext.JSON(http.StatusBadGateway, gin.H{
			"Message error: ": error.Error()})
		return
	}

	if err := models.ValidaDadosClientes(&user); err != nil {
		ctext.JSON(http.StatusBadGateway, gin.H{
			"Message error: ": err.Error()})
		return
	}

	conn, err := grpc.Dial("localhost:9000", grpc.WithTransportCredentials(insecure.NewCredentials()))
	id := ctext.Params.ByName("id")
	idInt64, _ := strconv.ParseInt(id, 10, 64)

	if err != nil {
		log.Fatal(err)
	}

	client := pb.NewUserServerClient(conn)
	req := &pb.Users{Id: idInt64, Username: user.Username, Passwd: user.Passwd, Email: user.Email}

	res, err := client.InsertUser(context.Background(), req)
	if err != nil {
		log.Fatal(err)
	}
	user = models.Users{Id: int(res.Id), Username: res.Username, Passwd: res.Passwd, Email: res.Email}

	if user.Id == 0 {
		ctext.JSON(http.StatusNotFound, gin.H{
			"Message": "User not inserted"})
		return
	}
	ctext.JSON(http.StatusOK, user)
}

func UpdateUser(ctext *gin.Context) {
	var user models.Users

	if error := ctext.ShouldBindJSON(&user); error != nil {
		ctext.JSON(http.StatusBadGateway, gin.H{
			"Message error: ": error.Error()})
		return
	}

	if err := models.ValidaDadosClientes(&user); err != nil {
		ctext.JSON(http.StatusBadGateway, gin.H{
			"Message error: ": err.Error()})
		return
	}

	conn, err := grpc.Dial("localhost:9000", grpc.WithTransportCredentials(insecure.NewCredentials()))
	id := ctext.Params.ByName("id")
	idInt64, _ := strconv.ParseInt(id, 10, 64)

	if err != nil {
		log.Fatal(err)
	}

	client := pb.NewUserServerClient(conn)
	req := &pb.UpdateRequest{Id: idInt64, User: &pb.Users{Id: idInt64, Username: user.Username, Passwd: user.Passwd, Email: user.Email}}

	res, err := client.UpdateUser(context.Background(), req)
	if err != nil {
		log.Fatal(err)
	}
	user = models.Users{Id: int(res.Id), Username: res.Username, Passwd: res.Passwd, Email: res.Email}

	if user.Id == 0 {
		ctext.JSON(http.StatusNotFound, gin.H{
			"Message": "User not inserted"})
		return
	}
	ctext.JSON(http.StatusOK, user)
}

func DeleteUser(ctext *gin.Context) {
	conn, err := grpc.Dial("localhost:9000", grpc.WithTransportCredentials(insecure.NewCredentials()))
	id := ctext.Params.ByName("id")
	idInt64, _ := strconv.ParseInt(id, 10, 64)

	if err != nil {
		log.Fatal(err)
	}

	client := pb.NewUserServerClient(conn)
	req := &pb.Id{Id: idInt64}

	res, err := client.DeleteUser(context.Background(), req)
	if err != nil {
		log.Fatal(err)
	}

	ctext.JSON(http.StatusOK, gin.H{
		"Message": res.Message})
}
