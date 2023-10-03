package post

import (
	"fmt"
	"time"
)

type Envelope struct {
	Destination_country string
	Recipient           string
	Source_county       string
	Sender              string
	Dispatch_time       time.Time
	Categoty            string
}

func (e Envelope) Send() {
	if (e.Destination_country == "") || (e.Source_county == "") {
		panic("Envelope: Destination_country or Source_county is empty")
	}

	fmt.Printf("Sending envelope from: %s to: %s\n", e.Source_county, e.Destination_country)
}
