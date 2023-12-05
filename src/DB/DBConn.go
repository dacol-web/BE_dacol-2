package DB

import "database/sql"

type DBConn struct {
	db    *sql.DB
	query string
}

func (ctx DBConn) Single() *sql.Row {
	defer ctx.db.Close()
	return ctx.db.QueryRow(ctx.query)
}

func (ctx DBConn) Many() *sql.Rows {
	defer ctx.db.Close()
	rows, err := ctx.db.Query(ctx.query)
	if err != nil {
		panic(err)
	}
	return rows
}

func (ctx DBConn) Exec() sql.Result {
	defer ctx.db.Close()
	result, err := ctx.db.Exec(ctx.query)
	if err != nil {
		panic(err)
	}
	return result
}
