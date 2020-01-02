package logicApi

import (
	"boxDemo/model"
)

type (
	BoxLogic struct {
		boxModel *model.Box
	}

	BoxListRequest struct {
		Page     int `json:"page"`
		PageSize int `json:"pageSize"`
	}

	BoxListResponse struct {
		total int64
		list  []*BoxView
	}

	BoxView struct {
		BoxId      int  `json:"boxId"`
		Size       int  `json:"size"`
		Color      int  `json:"color"`
		Status     uint `json:"status"`
		IsDeleted  uint `json:"isDeleted"`
		CreateId   uint `json:"createId"`
		CreateTime uint `json:"createTime"`
		UpdateTime uint `json:"updateTime"`
	}
)

func NewBoxLogic(boxModel *model.Box) *BoxLogic {
	return &BoxLogic{boxModel: boxModel}
}

func (l *BoxLogic) GetBoxList(r *BoxListRequest) (*BoxListResponse, error) {
	boxList, total, err := l.boxModel.GetBoxList(r.Page, r.PageSize)
	if err != nil {
		return nil, err
	}
	response := &BoxListResponse{total: total, list: []*BoxView(nil)}
	for _, val := range boxList {
		response.list = append(response.list, &BoxView{
			BoxId: val.BoxId, Size: val.Size, Color: val.Color, Status: val.Status, IsDeleted: val.IsDeleted, CreateId: val.CreateId,
			CreateTime: val.CreateTime, UpdateTime: val.UpdateTime,
		})
	}
	return response, nil
}
