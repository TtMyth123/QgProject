package GInstance

import "github.com/TtMyth123/QgProject/football-analysis/AnalystServer/analyst"

var (
	mWorkContainer *analyst.WorkContainer
)

func Init() {
	mWorkContainer = analyst.NewWorkContainer()
}

func GetWorkContainer() *analyst.WorkContainer {
	if mWorkContainer == nil {
		Init()
	}
	return mWorkContainer
}
