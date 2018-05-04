package main

import (
	_ "expvar"
	"fmt"
	"time"

	"github.com/bcicen/grmon/agent"
)

func main() {
	// http.ListenAndServe(":1234", nil)
	grmon.Start()
	for {
		time.Sleep(time.Second)
		fmt.Println("1-=-=")
	}
}

// package main
//
// import (
// 	"fmt"
// 	"log"
//
// 	"github.com/takama/daemon"
// )
//
// func main() {
// 	service, err := daemon.New("name", "description")
// 	if err != nil {
// 		log.Fatal("Error: ", err)
// 	}
// 	// service.Remove()
// 	status, err := service.Remove()
// 	if err != nil {
// 		log.Fatal(status, "\nError: ", err)
// 	}
// 	fmt.Println(status)
// }
