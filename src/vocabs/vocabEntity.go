package vocabs

type VocabEntity struct {
	UserID     int    `json:"userId"`
	FirstTerm  string `json:"firstTerm"`
	SecondTerm string `json:"secondTerm"`
	Languages  string `json:"languages"`
}
