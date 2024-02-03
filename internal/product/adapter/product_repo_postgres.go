package adapter

import "github.com/jackc/pgx"

type ProductRepo struct{
	*pgx.Conn
}