package quake

import (
	"google.golang.org/protobuf/proto"
	"io"
	"os"
)

func Read(file string) (Quake, error) {
	var q Quake

	f, err := os.Open(file)
	if err != nil {
		return q, err
	}

	defer f.Close()

	b, err := io.ReadAll(f)
	if err != nil {
		return q, err
	}

	err = proto.Unmarshal(b, &q)

	return q, err
}
