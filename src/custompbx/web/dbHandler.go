package web

import (
	"custompbx/db"
)

func InitDB(instanceId int64) {
	db.InitWebDB(instanceId)
}
