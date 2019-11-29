package agents

type Agent struct {
	Id           int  `json:"id"`
	Name         string `json:"name"`
	CmdUrl       string `json:"cmdUrl"`
	Host         string `json:"host"`
	HostName     string `json:"hostName"`
	Port         int  `json:"port"`
	Status       int  `json:"status"`
	ServFor      int  `json:"servFor"`
	LastPingTime int64  `json:"lastPingTime"`
	Available    bool   `json:"available"`
}

type AgentType int

const (
	Passenger  AgentType = 0
	FastDriver AgentType = 1
)

func (agentType AgentType) String() string {
	switch agentType {
	case Passenger:
		return "PASSENGER"
	case FastDriver:
		return "FASTDRIVER"
	default:
		return "UNKNOWN"
	}
}
