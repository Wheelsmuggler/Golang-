package mongo

var DataBase = "zhihucomment"

var Collection = "pingandcom"

var Url = "mongodb://localhost:27017"

type Ping struct {
	Id 			string
	OnePing		string
	Comments 	[4]string
	Tags 		[3][]string
}