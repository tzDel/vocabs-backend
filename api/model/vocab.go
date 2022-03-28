package model

type Vocab struct {
	UserID     int    `json:"user_id"`
	FirstTerm  string `json:"firstterm"`
	SecondTerm string `json:"secondterm"`
	Languages  string `json:"languages"`
}
