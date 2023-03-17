package interfaces

type RestartJob struct {
	Unit       string `json:"unit"`
	Expression string `json:"expression"`
}
