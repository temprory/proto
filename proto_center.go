package proto

import (
	"fmt"
	"strconv"
)

var (
	SVR_TYPE_CENTER = "center"
	SVR_TYPE_PROXY  = "proxy"
	SVR_TYPE_PLAZA  = "plaza"
	SVR_TYPE_GAME   = "game"

	SvrTypes = map[string]bool{
		SVR_TYPE_CENTER: true,
		SVR_TYPE_PROXY:  true,
		SVR_TYPE_PLAZA:  true,
		SVR_TYPE_GAME:   true,
	}
)

var (
	GAME_TYPE_BAC   = "bac"   //百家乐
	GAME_TYPE_DDZ   = "ddz"   //斗地主
	GAME_TYPE_NN    = "nn"    //牛牛
	GAME_TYPE_NN100 = "nn100" //百人牛牛
	GAME_TYPE_ZJH   = "zjh"   //炸金花

	GameTypes = map[string]bool{
		GAME_TYPE_BAC:   true,
		GAME_TYPE_DDZ:   true,
		GAME_TYPE_NN:    true,
		GAME_TYPE_NN100: true,
		GAME_TYPE_ZJH:   true,
	}
)

var (
	Empty *struct{} = nil

	CMD_CENTER_UPDATE_INFO_REQ = uint32(1)
	CMD_CENTER_UPDATE_INFO_RSP = uint32(2)

	CMD_CENTER_UPDATE_SERVER_LIST_NOTIFY = uint32(3)
)

type RoomInfo struct {
	GameType   string      `json:"gameType"`
	MaxPlayers int         `json:"maxPlayers"`
	OnlineNum  int         `json:"onlineNum"`
	Odds       interface{} `json:"odds"`
}

type SvrInfo struct {
	Type  string      `json:"svrType"`
	IP    string      `json:"ip"`
	Port  string      `json:"port"`
	Rooms []*RoomInfo `json:"rooms"`
}

type CenterUpdateInfoReq struct {
	ID     string  `json:"id"`
	Passwd string  `json:"passwd"`
	Info   SvrInfo `json:"info"`
}

type CenterUpdateInfoRsp struct {
	Code int    `json:"code"`
	Msg  string `json:"msg"`
}

func (s *SvrInfo) Valid() (int, error) {
	if !SvrTypes[s.Type] {
		return ERR_CENTER_INVALID_SVR_TYPE, fmt.Errorf("%v '%v'", ERR_CENTER_INVALID_SVR_TYPE_TXT, s.Type)
	}

	if s.Port == "" {
		return ERR_CENTER_INVALID_SVR_PORT, fmt.Errorf("%v '%v'", ERR_CENTER_INVALID_SVR_PORT_TXT, s.Port)
	}
	if _, err := strconv.ParseUint(s.Port, 10, 0); err != nil {
		return ERR_CENTER_INVALID_SVR_PORT, fmt.Errorf("%v '%v'", ERR_CENTER_INVALID_SVR_PORT_TXT, s.Port)
	}

	for _, room := range s.Rooms {
		if room != nil {
			if !GameTypes[room.GameType] {
				return ERR_CENTER_INVALID_GAME_TYPE, fmt.Errorf("%v '%v'", ERR_CENTER_INVALID_GAME_TYPE_TXT, room.GameType)
			}
		}
	}

	return 0, nil
}
