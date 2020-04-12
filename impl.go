package bj21

import (
	"context"
	"fmt"

	"github.com/Justyer/bj21/pb"
)

type BJ21Impl struct {
}

func (l *BJ21Impl) CreateTable(ctx context.Context, req *pb.CreateTableReq) (res *pb.CreateTableRes, err error) {
	tableIdx := int32(1)

	st := &Stage{}
	Tables[tableIdx] = st
	st.SetP1(&Player{
		Number: 1,
	})

	res = &pb.CreateTableRes{Data: &pb.CreateTableRes_Data{
		TableIdx: tableIdx,
	}}
	st.ShowRepo()
	return
}

func (l *BJ21Impl) JoinTable(ctx context.Context, req *pb.JoinTableReq) (res *pb.JoinTableRes, err error) {
	st := Tables[req.TableIdx]
	st.SetP2(&Player{
		Number: 2,
	})
	res = &pb.JoinTableRes{}
	st.ShowRepo()
	return
}

func (l *BJ21Impl) StartStage(ctx context.Context, req *pb.StartStageReq) (res *pb.StartStageRes, err error) {
	st := Tables[req.TableIdx]
	st.InitStage()
	st.ShowRepo()
	res = &pb.StartStageRes{}
	return
}

func (l *BJ21Impl) HIT(ctx context.Context, req *pb.HITReq) (res *pb.HITRes, err error) {
	st := Tables[req.TableIdx]
	st.Deal(req.PNumber)
	st.ShowRepo()
	res = &pb.HITRes{}
	return
}

func (l *BJ21Impl) CalcResult(ctx context.Context, req *pb.CalcResultReq) (res *pb.CalcResultRes, err error) {
	st := Tables[req.TableIdx]
	winner := st.Result(req.TableIdx)
	fmt.Println("winner", winner)
	res = &pb.CalcResultRes{}
	return
}
