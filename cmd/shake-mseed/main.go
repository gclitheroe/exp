// shake-mseed calculates PGA and PGV from mSEED data.
// It uses the current shaking config which may not be valid for historic data.
package main

import (
	"net/http"
	"log"
	"github.com/GeoNet/kit/mseed"
	"io"
	"fmt"
	"strings"
	"encoding/json"
	"bytes"
	"io/ioutil"
)

// the record length of the miniSEED records.  Constant for all GNS miniSEED files.
const recordLength int = 512
const fdsn = "https://service.geonet.org.nz/fdsnws/dataselect/1/query?network=%s&station=%s&location=%s&channel=%s&starttime=2016-11-13T11:00:00.000&endtime=2016-11-13T12:00:00.000"

func main() {
	// load configuration
	var a app
	err := a.init()
	if err != nil {
		log.Fatal(err)
	}

	client := http.Client{}

	msr := mseed.NewMSRecord()
	defer mseed.FreeMSRecord(msr)

	record := make([]byte, recordLength)

	for _, v := range a.config {

		log.Printf("requesting data for %s.%s.%s.%s", v.n, v.s, v.l, v.c)
		r, err := client.Get(fmt.Sprintf(fdsn, v.n, v.s, v.l, v.c))
		if err != nil {
			log.Fatal(err)
		}

		switch r.StatusCode {
		case http.StatusOK:
		case http.StatusNoContent:
			log.Printf("no data for %s.%s.%s.%s", v.n, v.s, v.l, v.c)
		default:
			log.Printf("non 200 or 204 response from service: %d", r.StatusCode)
		}

		//	read the mSEED from the response 512 bytes at a time and calculate max PGA and PGV
	loop:
		for {
			_, err = io.ReadFull(r.Body, record)
			switch {
			case err == io.EOF:
				break loop
			case err != nil:
				log.Print("unexpected EOF - data may be incomplete")
				break loop
			}

			err = msr.Unpack(record, recordLength, 1, 0)
			if err != nil {
				log.Fatal(err)
			}

			src := msr.SrcName(0)

			stream, ok := a.streams[src]
			if !ok {
				continue loop
			}

			if stream.Rate != float64(msr.Samprate()) {
				continue loop
			}

			d, err := msr.DataSamples()
			if err != nil {
				log.Printf("error reading data packet for %s: %s", src, err)
				continue loop
			}

			if stream.HaveGap(msr.Starttime()) {
				stream.Reset()
				stream.Condition(d)
			}

			key := fmt.Sprintf("%s.%s.%s", msr.Network(), msr.Station(), msr.Location())

			pga, pgv := stream.Peaks(d)

			if strings.HasSuffix(msr.Channel(), "Z") {
				if a.features[key].Properties.PGAV < pga {
					a.features[key].Properties.PGAV = pga
				}
				if a.features[key].Properties.PGVV < pgv {
					a.features[key].Properties.PGVV = pgv
				}
			} else {
				if a.features[key].Properties.PGAH < pga {
					a.features[key].Properties.PGAH = pga
				}
				if a.features[key].Properties.PGVH < pgv {
					a.features[key].Properties.PGVH = pgv
				}
			}

			stream.Last = msr.Endtime()
		}
		r.Body.Close()
	}

	// output GeoJSON
	f := locationFeatures{Type: "FeatureCollection"}

	for _, v := range a.features {
		// skip any locations with no data
		if v.Properties.PGAV == 0 && v.Properties.PGAH == 0 {
			continue
		}

		f.Features = append(f.Features, *v)
	}

	b, err := json.Marshal(&f)

	var out bytes.Buffer
	json.Indent(&out, b, "", "  ")

	err = ioutil.WriteFile("out.json", out.Bytes(), 0644)
	if err != nil {
		log.Fatal(err)
	}
}
