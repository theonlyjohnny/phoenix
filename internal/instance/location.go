package instance

import "fmt"

//A Location represents an Instances deployment location
type Location struct {
	Region string `json:"region"`
	Zone   string `json:"zone"`
}

func (l Location) String() string {
	return fmt.Sprintf("%s:%s", l.Region, l.Zone)
}
