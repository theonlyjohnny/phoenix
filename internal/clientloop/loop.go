package clientloop

import "time"

const (
	loopInterval = time.Second * 5
)

func Start() {

	ticker := time.NewTicker(loopInterval)

	for range ticker.C {
		loop()
	}
}

func loop() {
	//NOOP
}
