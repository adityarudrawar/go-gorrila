package structures

type OutputGorilla struct {
	Status    string `json:"status"`
	IsError   bool   `json:"isError"`
	Message   string `json:"message"`
	Id        int64  `json:"id"`
	IsHealthy bool   `json:"isHealthy"`
	Numbering []int  `json:"numbering"`
}
