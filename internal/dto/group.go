package dto

import (
	"time"
)

type Group struct {
	ID          string    `json:"id"`
	GroupName   string    `json:"group_name"`
	GroupAdmin  string    `json:"group_admin"`
	SimplifyTxn bool      `json:"simplify_txn"`
	CreatedAt   time.Time `json:"created_at"`
}

func (g Group) SetGroupAdmin(admin string) Group {
	g.GroupAdmin = admin
	return g
}
