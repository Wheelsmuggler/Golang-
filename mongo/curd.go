package mongo

import (
	"go.mongodb.org/mongo-driver/bson"
	"gopkg.in/mgo.v2"
	"log"
	"strconv"
)


func FindRelevantPing(keyWords [3][]string) (items []Ping) {
	mongo, err := mgo.Dial(Url)
	if err != nil || mongo == nil {
		log.Fatal(err)
	}
	defer mongo.Close()
	client := mongo.DB(DataBase).C(Collection)
	for i := 0; i < 3; i++ {
		iter := client.Find(bson.M{
			"taglevel"+strconv.Itoa(i+1): bson.M{
				"$all": keyWords[i],
			},
		}).Limit(7-i).Iter()

		var ping Ping
		for iter.Next(&ping)  {
			flag := false
			for index := range items {
				if items[index].Id==ping.Id {
					flag =true
					break
				}
			}
			if !flag{
				items = append(items,ping)
			}
		}
		//fmt.Printf("*********************level%d*****************************\n",i+1)
		//for index := range items{
		//	//fmt.Printf("intem:\n")
		//	//fmt.Printf("txt:%v\n",items[index].OnePing)
		//	for i:= range items[index].Comments {
		//		fmt.Printf("comments[%v]:%v\n",i,items[index].Comments[i])
		//	}
		//}
	}
	return items
}
