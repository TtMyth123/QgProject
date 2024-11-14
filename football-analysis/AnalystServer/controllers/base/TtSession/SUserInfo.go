package TtSession

import "github.com/TtMyth123/kit"

type SUserInfo struct {
	Id   int64
	SId  string
	Life int
}

func NewSUserInfo(id int64, sId string, life int) *SUserInfo {
	if sId == "" {
		sId = kit.GetGuid()
	}
	aSUserInfo := &SUserInfo{Id: id, SId: sId, Life: life}

	return aSUserInfo
}
