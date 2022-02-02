package data_transfer_object

import (
	"goodwe_monitor/domain/value_object"
)

type LoginResponse struct {
	HasError   bool               `json:"hasError"`
	Code       int                `json:"code"`
	Msg        string             `json:"msg"`
	Data       value_object.Token `json:"data"`
	Components struct {
		Para         interface{} `json:"para"`
		LangVer      int         `json:"langVer"`
		TimeSpan     int         `json:"timeSpan"`
		Api          string      `json:"api"`
		MsgSocketAdr string      `json:"msgSocketAdr"`
	} `json:"components"`
	Api string `json:"api"`
}
