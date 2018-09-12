package defs

type UserCredential struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

type VideoInfo struct {
	Id string
	AuthorId int
	Name string
	DisplayCtime string
}
type Comment struct {
	Id string
	VideoId string
	Author string
	Content string
}

type SimpleSession struct {
	Username string //login name
	TTL int64
}
//response
type SignedUp struct {
	Success bool `json:"success"`
	SessionId string `json:"session_id"`
}