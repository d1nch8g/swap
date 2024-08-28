// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.27.0

package database

import (
	"github.com/jackc/pgx/v5/pgtype"
)

type Admin struct {
	ID        int64
	Email     string
	Passwhash string
}

type Exchanger struct {
	ID      int64
	Inmin   pgtype.Numeric
	Inmax   pgtype.Numeric
	Reserve pgtype.Numeric
	Rate    pgtype.Numeric
	Change  string
}

type Order struct {
	ID      int64
	Give    string
	Receive string
	UserID  int64
}

type OrderChat struct {
	ID      int64
	UserID  int64
	OrderID int64
}

type User struct {
	ID       int64
	Email    string
	Card     string
	Verified bool
}
