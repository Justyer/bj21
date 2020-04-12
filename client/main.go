package main

import (
	"context"
	"fmt"
	"log"

	"github.com/Justyer/bj21/pb"
	"google.golang.org/grpc"
)

func main() {
	conn, err := grpc.Dial(":9090", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()

	client := pb.NewBJ21Client(conn)

	var p int32
	fmt.Println("login")
	fmt.Scanf("%d", &p)

	fmt.Println("1创建房间/2加入房间/3开始对局/4要牌/5计算结果")
	var choice int
	for {
		fmt.Scanf("%d", &choice)
		switch choice {
		case 1:
			req := &pb.CreateTableReq{
				Name: "XC-Jiang",
			}
			res, err := client.CreateTable(context.Background(), req)
			if err != nil {
				log.Fatalf("CreateTable: %v", err)
			}

			fmt.Println(res.String())
		case 2:
			req := &pb.JoinTableReq{
				TableIdx: 1,
			}
			res, err := client.JoinTable(context.Background(), req)
			if err != nil {
				log.Fatalf("JoinTable: %v", err)
			}

			fmt.Println(res.String())
		case 3:
			req := &pb.StartStageReq{
				TableIdx: 1,
			}
			res, err := client.StartStage(context.Background(), req)
			if err != nil {
				log.Fatalf("StartStage: %v", err)
			}

			fmt.Println(res.String())
		case 4:
			req := &pb.HITReq{
				TableIdx: 1,
				PNumber:  p,
			}
			res, err := client.HIT(context.Background(), req)
			if err != nil {
				log.Fatalf("HIT: %v", err)
			}

			fmt.Println(res.String())
		case 5:
			req := &pb.CalcResultReq{
				TableIdx: 1,
			}
			res, err := client.CalcResult(context.Background(), req)
			if err != nil {
				log.Fatalf("CalcResult: %v", err)
			}

			fmt.Println(res.String())
		}
	}
}
