package naming

import (
	"errors"
	"log"

	"github.com/mvgmb/Middleware/rpc/util"
)

var (
	servicesTable = make(map[string][]util.AOR)
	roundRobin    = -1
)

func bind(aor *util.AOR) {
	log.Println("New service registered: " + aor.ObjectID)
	servicesTable[aor.ObjectID] = append(servicesTable[aor.ObjectID], *aor)
}

func lookup(serviceName string) (*util.AOR, error) {
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
