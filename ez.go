package ez

import "ez/model"

// ---- ---- ---- ----

// Ez ORM 2018.4.18

// ---- ---- ---- ----
var  GetClient = func() (*DB) {
	return &DB{}
}


// ---- ---- ---- ----
type DB struct {
}
func (d *DB) Create (v interface{}) {
	model.SetupEzORMContext(v)
}
