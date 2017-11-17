package main

import (
	"github.com/GeoNet/kit/shake"
	"io/ioutil"
	"strings"
	"io"
	"encoding/json"
	"github.com/pkg/errors"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/s3"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/session"
	"fmt"
)

// config is for parsing the JSON configuration file.
// The JSON is generated using https://github.com/GeoNet/delta/tree/master/tools/impact
//   ./impact --channels "[BH]N[ZNE12XY]" > impact.json
type config struct {
	Gain      float64
	Q         float64
	Rate      float64
	Longitude float64
	Latitude  float64

	n,s,l,c string

	srcname string // key in the JSON e.g., NZ_WVZ_20_BNE
}

type app struct {
	streams map[string]*shake.Stream
	config  []config
	features   map[string]*locationFeature
}

func (a *app) init() error {
	var s3Client *s3.S3

	s, err := session.NewSession(&aws.Config{
		Credentials: credentials.AnonymousCredentials,
		Region:      aws.String("ap-southeast-2"),
	})
	if err != nil {
		return err
	}
	s3Client = s3.New(s)

	conf, err := s3Client.GetObject(&s3.GetObjectInput{
		Bucket: aws.String("geonet-meta"),
		Key:    aws.String("config/impact.json"),
	})
	if err != nil {
		return err
	}
	defer conf.Body.Close()

	return a.loadConfig(conf.Body)
}

func (a *app) loadConfig(r io.Reader) error {
	var err error

	b, err := ioutil.ReadAll(r)
	if err != nil {
		return err
	}

	var m = make(map[string]config)

	err = json.Unmarshal(b, &m)
	if err != nil {
		return err
	}

	a.features = make(map[string]*locationFeature)

	for k, v := range m {
		v.srcname = k

		p := strings.Split(k, "_")

		if len(p) != 4 {
			return errors.Errorf("expected 4 parts to srcname: %s", k)
		}

		v.n = p[0]
		v.s = p[1]
		v.l = p[2]
		v.c = p[3]

		a.config = append(a.config, v)

		// store features at network.station.location
		a.features[fmt.Sprintf("%s.%s.%s", p[0], p[1], p[2])] = &locationFeature{
			Type: "Feature",
			Geometry: point{
				Type: "Point",
				Coordinates: [2]float64{
					v.Longitude,
					v.Latitude,
				},
			},
			Properties: location{
				Network:  p[0],
				Station:  p[1],
				Location: p[2],
			},
		}
	}

	a.streams = make(map[string]*shake.Stream)

	for _, v := range a.config {
		a.streams[v.srcname] = &shake.Stream{
			Rate:       v.Rate,
			HighPass:   shake.NewHighPass(v.Gain, v.Q),
			Integrator: shake.NewIntegrator(1.0, 1.0/v.Rate, v.Q),
		}
	}

	return nil
}
