package agents

type CommonRequest struct {
	MsgId string `json:"msgId"`
	Cmd   int    `json:"cmd"`
	Data  string `json:"data"`
}
