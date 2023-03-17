package interfaces

type RebootJob struct {
	Tag        string `json:"tag"`
	Expression string `json:"expression"`
}
