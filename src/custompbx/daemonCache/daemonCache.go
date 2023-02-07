package daemonCache

import "custompbx/mainStruct"

var State *mainStruct.DaemonState

func InitDaemonState() {
	State = &mainStruct.DaemonState{
		DatabaseConnection: false,
		ESLConnection:      false,
	}
}
