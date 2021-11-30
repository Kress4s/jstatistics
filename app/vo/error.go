package vo

type Error struct {
	Msg  string   `json:"msg"`
	Args []string `json:"args"`
	Code int      `json:"code"`
}
