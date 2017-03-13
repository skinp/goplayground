package main

import (
	"crypto/tls"
	"flag"
	"fmt"
	"log"
	"net"
	"time"

	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

func defaultTLSDial(addr *mgo.ServerAddr) (net.Conn, error) {
	return tls.Dial("tcp", addr.String(), nil)
}

var url = flag.String("url", "", "url to iterate on")

type Person struct {
	Name string `bson:"name"`
}

func main() {
	flag.Parse()
	if *url == "" {
		log.Fatal("no url provided")
	}

	info, err := mgo.ParseURL(*url)
	info.DialServer = defaultTLSDial
	info.Timeout = 10 * time.Second

	session, err := mgo.DialWithInfo(info)
	if err != nil {
		log.Fatalf("can't get mgo session: %s", err)
	}

	collection := session.DB(info.Database).C("person")
	//err = collection.Insert(bson.M{"name": "patrick"})
	//if err != nil {
	//	log.Fatalf("can't insert in person: %s", err)
	//}

	query := collection.Find(bson.M{})
	iter := query.Iter()

	result := Person{}
	docs := 0
	for iter.Next(&result) {
		docs++
	}
	if iter.Err() != nil {
		log.Fatalf("error during iter: %s", iter.Err())
	}

	fmt.Println("we succeeded with ", docs, " documents")
}
