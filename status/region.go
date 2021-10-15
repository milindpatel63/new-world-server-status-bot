package status

import (
	"errors"
	"fmt"
)

type Region int

const (
	UsEast Region = iota
	SaEast
	EuCentral
	ApSoutheast
	UsWest
)

func (r Region) String() string {
	switch r {
	case UsEast:
		return "US EAST"
	case SaEast:
		return "SA EAST"
	case EuCentral:
		return "EU CENTRAL"
	case ApSoutheast:
		return "AP SOUTHEAST"
	case UsWest:
		return "US WEST"
	default:
		return fmt.Sprintf("%d", int(r))
	}
}

func ParseRegion(region string) (Region, error) {
	if region == "USE" || region == "use" {
		return UsEast, nil
	} else if region == "USW" || region == "usw" {
		return UsWest, nil
	} else if region == "SA" || region == "sa" {
		return SaEast, nil
	} else if region == "EU" || region == "eu" {
		return EuCentral, nil
	} else if region == "AP" || region == "ap" {
		return ApSoutheast, nil
	} else {
		return -1, errors.New("could not parse region parameter: " + region)
	}
}
