package nocache

import (
	"custompbx/db"
	"custompbx/intermediateDB"
)

func InitDB() {
	//db.InitDirectoryDB()
	intermediateDB.InitDirectoryDB()
	intermediateDB.InitConfDB()
	db.InitDialplanDB()
}
