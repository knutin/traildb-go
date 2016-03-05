package main

import "fmt"

func main() {
	tdb, err := Open("output.30day.29of30.tdb")
	if err != nil {
		panic(err.Error())
	}
	fmt.Println(tdb)
	fmt.Println(tdb.Version())

	var total int
	for i := 0; i < tdb.numTrails; i++ {
		trail, err := NewTrail(tdb, i)
		if err != nil {
			panic(err.Error())
		}
		// fmt.Println(trail)
		for {
			evt := trail.NextEvent()
			if evt == nil {
				trail.Close()
				break
			}
			total++
			// evt.Print()
		}
	}
	fmt.Println(total)
	tdb.Close()
}