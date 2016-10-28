package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"time"
)

func main() {
	log.SetFlags(log.Lmicroseconds | log.Lshortfile)

	/*
		db := postgres.Wait(nil, nil)

		m := postgres.NewMigrations()
		m.Add(1, "CREATE SEQUENCE hits")
		if err := m.Migrate(db); err != nil {
			log.Fatal(err)
		}
	*/

	port := os.Getenv("PORT")

	http.HandleFunc("/", func(w http.ResponseWriter, req *http.Request) {
		/*
			var count int
			if err := db.QueryRow("SELECT nextval('hits')").Scan(&count); err != nil {
				w.WriteHeader(500)
				w.Write([]byte(err.Error()))
				return
			}
		*/
		//返回现在时间
		tNow := time.Now()
		//时间转化为string，layout必须为 "2006-01-02 15:04:05"
		timeNow := tNow.Format("2006-01-02 15:04:05")
		fmt.Fprintf(w, "tNow(time format):", tNow)
		fmt.Fprintf(w, "tNow(string format):", timeNow)
		//fmt.Fprintf(w, "Hello from Flynn on port %s from container %s\nHits = %d\n", port, os.Getenv("HOSTNAME"), count)
	})
	http.HandleFunc("/17bb00aed475e027492608694e6b4e1e.txt", func(w http.ResponseWriter, req *http.Request) { fmt.Fprint(w, "17bb00aed475e027492608694e6b4e1e") })
	fmt.Println("hitcounter listening on port", port)
	log.Fatal(http.ListenAndServe(":"+port, nil))
}
