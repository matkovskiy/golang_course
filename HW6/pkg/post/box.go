package post

import (
	"fmt"
	"time"
)

type Box struct {
	Destination_country string
	Recipient           string
	Source_county       string
	Sender              string
	Dispatch_time       time.Time
	Weight              string
}

func (b Box) Send() {
	if (b.Destination_country == "") || (b.Source_county == "") {
		panic("BOX: Destination_country or Source_county is empty")
	}
	fmt.Printf("Sending the Box from: %s to: %s\n", b.Source_county, b.Destination_country)
}
