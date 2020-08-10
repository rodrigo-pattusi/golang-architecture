package main

import (
	"fmt"

	architecture "github.com/rodrigo-pattusi/golang-architecture"
	"github.com/rodrigo-pattusi/golang-architecture/storage/mongo"
	"github.com/rodrigo-pattusi/golang-architecture/storage/postgres"
)

func main() {
	dbm := mongo.Db{}
	dbp := postgres.Db{}

	p1 := architecture.Person{
		First: "Jenny",
	}

	p2 := architecture.Person{
		First: "James",
	}

	ps := architecture.NewPersonService(dbm)

	architecture.Put(dbm, 1, p1)
	architecture.Put(dbm, 2, p2)
	fmt.Println(architecture.Get(dbm, 1))
	fmt.Println(architecture.Get(dbm, 2))

	fmt.Println(ps.Get(1))
	fmt.Println(ps.Get(3))

	// or store in some other data storage
	architecture.Put(dbp, 1, p1)
	architecture.Put(dbp, 2, p2)
	fmt.Println(architecture.Get(dbp, 1))
	fmt.Println(architecture.Get(dbp, 2))
}

//package main
//
//import (
//"fmt"
//)
//
//type user struct {
//	first string
//}
//
//type mongo map[int]user
//
//func (m mongo) save(n int, u user) {
//	m[n] = u
//}
//
//func (m mongo) retrieve(n int) user {
//	return m[n]
//}
//
//type harddrive map[int]user
//
//func (hd harddrive) save(n int, u user) {
//	hd[n] = u
//}
//
//func (hd harddrive) retrieve(n int) user {
//	return hd[n]
//}
//
//type accessor interface {
//	save(n int, u user)
//	retrieve(n int) user
//}
//
//func put(a accessor, n int, u user) {
//	a.save(n, u)
//}
//
//func get(a accessor, n int) user {
//	return a.retrieve(n)
//}
//
//func main() {
//	storage := harddrive{}
//
//	u1 := user{
//		first: "James",
//	}
//
//	u2 := user{
//		first: "Jenny",
//	}
//
//	put(storage, 1, u1)
//	put(storage, 2, u2)
//
//	fmt.Println(get(storage, 1))
//	fmt.Println(get(storage, 2))
//}
