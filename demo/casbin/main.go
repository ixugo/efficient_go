package main

import (
	"fmt"

	"github.com/casbin/casbin/v2"
)

func main() {
	// pq := postgres.Open(`postgres://postgres:123456789@localhost:1556/saida?sslmode=disable`)
	// db, err := gorm.Open(pq, &gorm.Config{})
	// if err != nil {
	// 	panic(err)
	// }
	// var m runtime.MemStats

	// runtime.ReadMemStats(&m)
	// fmt.Printf("%d MB\n", m.Alloc/1024/1024)
	// a, err := gormadapter.NewAdapterByDB(db)
	// if err != nil {
	// 	panic(err)
	// }
	e, err := casbin.NewEnforcer("/Users/xugo/Documents/efficient_go/demo/casbin/casbin.conf", "/Users/xugo/Documents/efficient_go/demo/casbin/policy.csv")
	if err != nil {
		panic(err)
	}

	// runtime.ReadMemStats(&m)
	// fmt.Printf("%d MB\n", m.Alloc/1024/1024)

	// start := time.Now()
	// e.LoadPolicy()
	// if err := e.LoadFilteredPolicy(&gormadapter.Filter{
	// 	V1: []string{"dedao"},
	// }); err != nil {
	// 	panic(err)
	// }
	// fmt.Println("load policy used time:", time.Since(start))
	ok, err := e.Enforce("137df13e-e9", "dedao", "/admin/123", "DELETE")
	if err != nil {
		panic(err)
	}
	fmt.Println("dedao:", ok)

	ok, err = e.AddPolicy("1313e-e9", "dedao", "/aabc/123", "DELETE")
	if err != nil {
		panic(err)
	}
	fmt.Println(ok)
	if err := e.SavePolicy(); err != nil {
		panic(err)
	}

	s := e.GetAllNamedSubjects("p")
	fmt.Println("s:", s)

	// ok, err = e.Enforce("137df13e-e9", "dedao", "/admin/123", "DELETE")
	// if err != nil {
	// 	panic(err)
	// }
	// fmt.Println("apple", ok)

	// start = time.Now()
	// e.ClearPolicy()
	// if length := len(e.GetPolicy()); length > 0 {
	// 	fmt.Println(length)
	// 	e.ClearPolicy()
	// }
	// fmt.Println("remove policy used time:", time.Since(start))

	// e.RemovePolicy()

	// for i := 0; i < 100000; i++ {

	// 	str := uuid.NewString()
	// 	length := rand.Intn(10) + 5
	// 	_, err := e.AddPolicy(str[0:length], "apple", "/admin/*", "GET")
	// 	if err != nil {
	// 		panic(err)
	// 	}
	// 	e.AddPolicy(str[0:length], "apple", "/admin/*", "POST")
	// 	e.AddPolicy(str[0:length], "apple", "/admin/*", "PUT")
	// 	e.AddPolicy(str[0:length], "apple", "/admin/*", "DELETE")
	// 	e.AddPolicy(str[0:length], "dedao", "/admin/*", "DELETE")
	// 	e.AddPolicy(str[0:length], "dedao", "/admin/*", "PUT")
	// 	if i%1000 == 0 {
	// 		fmt.Println("add policy used time:", time.Since(start))
	// 		start = time.Now()
	// 		runtime.ReadMemStats(&m)
	// 		fmt.Printf("%d MB\n", m.Alloc/1024/1024)
	// 	}

	// }
	fmt.Println("end")

	// runtime.ReadMemStats(&m)
	// fmt.Printf("%d MB\n", m.Alloc/1024/1024)

	// ok, err = e.DeleteRole("1313e-e9")
	// if err != nil {
	// 	panic(err)
	// }
	// fmt.Println(ok)
	// e.SavePolicy()
	// e.AddRoleForUserInDomain("asdasd", "123", "dedao")
	// e.AddRoleForUserInDomain("5234", "123", "dedao")
	// e.AddRoleForUserInDomain("12", "123", "dedao")
	// e.AddRoleForUserInDomain("5234", "123", "dedao")
	// e.AddRoleForUserInDomain("5234", "52342", "dedao")

	// ok, err = e.RemoveFilteredGroupingPolicy(1, "123")
	// if err != nil {
	// 	panic(err)
	// }
	// fmt.Println(ok)

	// ok, err = e.RemoveFilteredGroupingPolicy(1, "123")
	// if err != nil {
	// 	panic(err)
	// }
	// fmt.Println(ok)
	// e.SavePolicy()
	vv := e.GetFilteredPolicy(0, "137df13e-e9", "", "", "DELETE")

	for _, v := range vv {
		fmt.Println("v:", v)
	}
}
