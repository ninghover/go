package handler

import (
	"context"

	"strconv"

	"house/model"
	pb "house/proto"
	"house/utils"
)

type House struct{}

func (e *House) GetHouseInfo(ctx context.Context, req *pb.HouseInfoReq, rsp *pb.HouseInfoRsp) error {
	houseInfos, err := model.GetUserHouse(req.Name)
	if err != nil {
		rsp.Errno = utils.RECODE_DBERR
		rsp.Errmsg = utils.RecodeText(utils.RECODE_DBERR)
		return nil
	}
	rsp.Errno = utils.RECODE_OK
	rsp.Errmsg = utils.RecodeText(utils.RECODE_OK)
	data := pb.GetData{Houses: houseInfos}
	rsp.Data = &data
	return nil
}

func (e *House) PostHouseInfo(ctx context.Context, req *pb.PostHouseReq, rsp *pb.PostHouseRsp) error {
	houseId, err := model.AddHouseInfo(req)
	if err != nil {
		rsp.Errno = utils.RECODE_DBERR
		rsp.Errmsg = utils.RecodeText(utils.RECODE_DBERR)
	} else {
		rsp.Errno = utils.RECODE_OK
		rsp.Errmsg = utils.RecodeText(utils.RECODE_OK)
		data := pb.PostHouseRsp_Data{
			HouseId: strconv.Itoa(houseId),
		}
		rsp.Data = &data
	}
	return nil
}
