// +build racedebug

package libcontainer

import "time"

func raceDebug() {
	time.Sleep(time.Second)
}
