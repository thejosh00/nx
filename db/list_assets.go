package db

import (
	"database/sql"
	"fmt"
	_ "github.com/lib/pq"
	"log"
	"os"
	"strconv"
	"text/tabwriter"
)

type DbListAssetsCommand struct {
	Format string `short:"f" long:"format" default:"maven2" description:"Format of assets to query"`
	Kind   string `short:"k" long:"kind" description:"Kind of assets to query"`
}

func (cmd *DbListAssetsCommand) Execute(args []string) error {

	format := cmd.Format

	db, err := sql.Open("postgres",
		"postgres://joshuahill:@localhost/development?sslmode=disable")
	if err != nil {
		return err
	}
	defer db.Close()

	err = db.Ping()
	if err != nil {
		return err
	}

	err = listAssets(db, format, cmd.Kind)
	if err != nil {
		return err
	}
	return nil
}

func listAssets(db *sql.DB, format string, kind string) error {
	where := ""
	args := []any{}
	if kind != "" {
		where += " WHERE kind = $1"
		args = append(args, kind)
	}

	var count int

	query := "select count(*) from " + format + "_asset"
	if where != "" {
		query += where
	}

	log.Println(query)
	log.Println(args)
	stmt, err := db.Prepare(query)
	if err != nil {
		return err
	}

	row := stmt.QueryRow(args...)
	row.Scan(&count)

	fmt.Println("Listing", format, "assets", "("+strconv.Itoa(count)+")")

	query = "select path, kind from " + format + "_asset"
	if where != "" {
		query += where
	}
	query += " limit 30"

	log.Println(query)
	log.Println(args)
	stmt, err = db.Prepare(query)
	if err != nil {
		return err
	}

	rows, err := stmt.Query(args...)
	if err != nil {
		return err
	}
	defer rows.Close()

	w := new(tabwriter.Writer)
	w.Init(os.Stdout, 80, 8, 0, '\t', 0)
	defer w.Flush()

	fmt.Fprintf(w, "\n %s\t%s\t", "Path", "Kind")
	fmt.Fprintf(w, "\n %s\t%s\t", "----", "----")
	var path string
	var kindValue string
	for rows.Next() {
		err = rows.Scan(&path, &kindValue)
		if err != nil {
			return err
		}
		fmt.Fprintf(w, "\n %s\t%s\t", path, kindValue)
	}
	fmt.Fprintf(w, "\n")
	err = rows.Err()
	if err != nil {
		return err
	}

	return nil
}
