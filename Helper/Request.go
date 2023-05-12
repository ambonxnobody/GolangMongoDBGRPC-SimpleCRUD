package Helper

import (
	helperPb "GolangMongoDBGRPCSimpleCRUD/Controller/Helper"
)

type ResponseOnAllData struct {
	PreviousPage int64         `json:"PreviousPage"`
	CurrentPage  int64         `json:"CurrentPage"`
	NextPage     int64         `json:"NextPage"`
	Limit        int64         `json:"Limit"`
	Offset       int64         `json:"Offset"`
	TotalPage    int64         `json:"TotalPage"`
	TotalData    int64         `json:"TotalData"`
	Data         []interface{} `json:"Data"`
}

func (r *ResponseOnAllData) ProcessResponseOnAllData(totalData int64, limit int64, page int64) (rPb helperPb.ResponseOnAllData) {
	rPb.CurrentPage = page
	rPb.Limit = limit
	rPb.Offset = (page - 1) * limit
	rPb.TotalData = totalData
	rPb.TotalPage = totalData / limit
	if totalData%limit != 0 {
		rPb.TotalPage = rPb.TotalPage + 1
	}
	rPb.PreviousPage = page - 1
	rPb.NextPage = page + 1
	if rPb.NextPage > rPb.TotalPage {
		rPb.NextPage = 0
	}

	return rPb
}
