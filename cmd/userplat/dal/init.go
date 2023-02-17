package dal

import "github.com/808-not-found/tik_duck/cmd/userplat/dal/db"

// Init init dal.
func Init() {
	db.Init() // mysql init
}
