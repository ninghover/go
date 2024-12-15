package handler

import (
	"context"

	"getArea/model"
	pb "getArea/proto"
	"getArea/utils"
)

type GetArea struct{}

func (e *GetArea) GetArea(ctx context.Context, req *pb.AreaReq, rsp *pb.AreaRsp) error {
	areas, err := model.GetArea()
	if err != nil {
		rsp.Errno = utils.RECODE_DBERR
		rsp.Errmsg = utils.RecodeText(utils.RECODE_DBERR)
		return nil
	}
	for _, v := range areas {
		var area pb.AreaInfo
		area.Aid = uint32(v.Id)
		area.Aname = v.Name
		rsp.Data = append(rsp.Data, &area)
	}
	rsp.Errno = utils.RECODE_OK
	rsp.Errmsg = utils.RecodeText(utils.RECODE_OK)
	return nil
}
