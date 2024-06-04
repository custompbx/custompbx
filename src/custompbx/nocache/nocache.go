package nocache

import (
	"custompbx/db"
	"custompbx/intermediateDB"
)

func InitDB() {
	intermediateDB.InitDirectoryDB()
	intermediateDB.InitConfDB()
	db.InitDialplanDB()
}

func InitServicesDB() {
	intermediateDB.InitServicesDB()
}
