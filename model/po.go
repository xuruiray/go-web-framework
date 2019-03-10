package model

// TestPo
type TestPo struct {
	ID   int64  `db:"id" json:"id"`
	Name string `db:"name" json:"name"`
}
