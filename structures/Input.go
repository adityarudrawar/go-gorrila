package structures

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
