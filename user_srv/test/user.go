package main

import (
	"context"
	"fmt"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"project/catsshop/user_srv/proto/proto"
)

var userClient proto.UserClient
var conn *grpc.ClientConn

func Init() {
	var err error
	conn, err = grpc.Dial("127.0.0.1:50051", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		panic(err)
	}

	userClient = proto.NewUserClient(conn)
}

func TestGetUserList() {
	rsp, err := userClient.GetUserList(context.Background(), &proto.PageInfo{
		Pn:    1,
		PSize: 5,
	})
	if err != nil {
		panic("getuserlist" + err.Error())
	}
	for _, user := range rsp.Data {
		fmt.Println(user.Mobile, user.NickName, user.Password)
		prsp, err := userClient.CheckPassword(context.Background(), &proto.PasswordCheckInfo{
			Password:          "admin123",
			EncryptedPassword: user.Password,
		})
		if err != nil {
			panic("checkpassword" + err.Error())
		}
		fmt.Println(prsp.Success)
	}

}

func TestGetUserByMobileAndId() {
	rsp1, err := userClient.GetUserByMobile(context.Background(), &proto.MobileRequest{
		Mobile: "13222222222",
	})
	if err != nil {
		panic("mobile" + err.Error())
	}
	fmt.Println(rsp1.NickName)
	rsp2, err := userClient.GetUserById(context.Background(), &proto.IdRequest{
		Id: 11,
	})
	if err != nil {
		panic("id" + err.Error())
	}
	fmt.Println(rsp2.NickName)
}

func TestUpdateUser() {
	rsp, err := userClient.UpdateUser(context.Background(), &proto.UpdateUserInfo{
		Id:       13,
		NickName: "wangheng",
		Gender:   "female",
	})
	if err != nil {
		panic("update" + err.Error())
	}
	fmt.Println(rsp)
}

func TestCreateUser() {
	for i := 0; i < 10; i++ {
		rsp, err := userClient.CreateUser(context.Background(), &proto.CreateUserInfo{
			NickName: fmt.Sprintf("bobby%d", i),
			Password: "admin123",
			Mobile:   fmt.Sprintf("1322222222%d", i),
		})
		if err != nil {
			panic(err)
		}
		fmt.Println(rsp.Id)
	}
}

func main() {
	Init()

	//TestGetUserList()
	//TestCreateUser()
	//TestGetUserByMobileAndId()
	TestUpdateUser()

	conn.Close()
}
