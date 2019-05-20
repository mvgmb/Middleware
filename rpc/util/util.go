package util

import (
	"fmt"
	"log"
	"strconv"
	"strings"
)

// AOR is the Absolute Object Reference
type AOR struct {
	Host     string
	Port     uint16
	ObjectID string
}

func (e *AOR) String() string {
	return fmt.Sprintf("%s:%d:%s", e.Host, e.Port, e.ObjectID)
}

// StringToAOR converts a string to an AOR
func StringToAOR(aor string) *AOR {
	split := strings.Split(aor, ":")

	num, err := strconv.ParseUint(split[1], 10, 16)
	if err != nil {
		log.Fatal(err)
	}

	e := AOR{
		Host:     split[0],
		Port:     uint16(num),
		ObjectID: split[2],
	}

	return &e
}
