package main

import (
	"flag"
	"fmt"
	"time"

	"github.com/araddon/dateparse"
	"github.com/scylladb/termtables"
)

var examples = []string{
	"1660169597559109041",
	"1660169607569643098",
	"1660169617560952864",
	"1660169627564631025",
	"1660169654097770605",
	"1660169657572251142",
	"1660169667576722857",
	"1660169687470964379",
	"1660169687571762654",
	"1660169697572025601",
	"1660169707575154843",
	"1660169717567763188",
	"1660169727569159420",
	"1660169737929715064",
	"1660169747562399727",
	"1660169757559727926",
	"1660169767560283225",
}

var (
	timezone = ""
)

func main() {
	flag.StringVar(&timezone, "timezone", "UTC", "Timezone aka `America/Los_Angeles` formatted time-zone")
	flag.Parse()

	if timezone != "" {
		// NOTE:  This is very, very important to understand
		// time-parsing in go
		loc, err := time.LoadLocation(timezone)
		if err != nil {
			panic(err.Error())
		}
		time.Local = loc
	}

	table := termtables.CreateTable()

	table.AddHeaders("Input", "Parsed, and Output as %v")

	var prevSecond float64
	for i, dateExample := range examples {
		t, err := dateparse.ParseLocal(dateExample)
		if err != nil {
			panic(err.Error())
		}

		if i > 0 {
			prevTimeStamp, _ := dateparse.ParseLocal(examples[i-1])
			Duration := t.Sub(prevTimeStamp)
			prevSecond += Duration.Seconds()
			//table.AddRow(dateExample, fmt.Sprintf("%v", prevSecond))
			fmt.Println(prevSecond)
		} else {
			//table.AddRow(dateExample, fmt.Sprintf("%v", 0))
			prevSecond = 0
			fmt.Println(prevSecond)
		}
		//table.AddRow(dateExample, fmt.Sprintf("%v", t))
	}
	//fmt.Println(table.Render())
}
