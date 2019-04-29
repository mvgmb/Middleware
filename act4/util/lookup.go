package util

import (
	"errors"
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

var (
	servicesTable = make(map[string][]AOR)
	roundRobin    = -1
)

func Bind(aor *AOR) {
	servicesTable[aor.ObjectID] = append(servicesTable[aor.ObjectID], *aor)
}

func Lookup(serviceName string) (*AOR, error) {
	if _, ok := servicesTable[serviceName]; !ok {
		return nil, errors.New("404 - Service not found")
	}

	if roundRobin+1 == len(servicesTable[serviceName]) {
		roundRobin = 0
	} else {
		roundRobin++
	}

	return &servicesTable[serviceName][roundRobin], nil
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
