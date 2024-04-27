package model

type ApiStruct struct {
	Id        int    `json:"id"`
	Name      string `json:"name"`
	Age       int    `json:"age"`
	TimeStamp string `json:"timestamp"`
	//IsActive bool   `json:"is_acitve"`
	//website  string `json:"website"`
}
