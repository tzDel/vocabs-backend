package model

type Vocab struct {
	UserID     int    `json:"userId"`
	FirstTerm  string `json:"firstTerm"`
	SecondTerm string `json:"secondTerm"`
	Languages  string `json:"languages"`
}
