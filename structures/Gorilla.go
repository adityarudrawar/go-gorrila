package structures

type OutputGorilla struct {
	Status     string   `json:"status"`
	IsError    bool     `json:"isError"`
	Message    string   `json:"message"`
	Id         int64    `json:"id"`
	IsHealthy  bool     `json:"isHealthy"`
	Numbering  []int    `json:"numbering"`
	Name       string   `json:"name"`
	HasFriends bool     `json:"hasFriends"`
	HasLivedIn []string `json:"hasLivedIn"`
	Medicals   Medical  `json:"medicals"`
}

type Medical struct {
	BloodPressure int64  `json:"bloodPressure"`
	HasSugar      bool   `json:"hasSugar"`
	Resp          string `json:"resp"`
}

type InputGorilla struct {
	Id         int64    `json:"id"`
	Name       string   `json:"name"`
	HasFriends bool     `json:"hasFriends"`
	HasLivedIn []string `json:"hasLivedIn"`
	Medicals   Medical  `json:"medicals"`
}
