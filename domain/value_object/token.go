package value_object

type Token struct {
	Uid       string `json:"uid"`
	Timestamp int64  `json:"timestamp"`
	Token     string `json:"token"`
	Client    string `json:"client"`
	Version   string `json:"version"`
	Language  string `json:"language"`
}
