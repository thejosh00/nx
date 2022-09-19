package db

import (
	"database/sql"
	"fmt"
	_ "github.com/lib/pq"
	"os"
	"text/tabwriter"
)

type DbListAssetsCommand struct {
	Format string `short:"f" long:"format" default:"maven2" description:"Format of assets to query"`
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

	err = listAssets(db, format)
	if err != nil {
		return err
	}
	return nil
}

func listAssets(db *sql.DB, format string) error {
	rows, err := db.Query("select path, kind from " + format + "_asset limit 30")
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
	var kind string
	for rows.Next() {
		err = rows.Scan(&path, &kind)
		if err != nil {
			return err
		}
		fmt.Fprintf(w, "\n %s\t%s\t", path, kind)
	}
	fmt.Fprintf(w, "\n")
	err = rows.Err()
	if err != nil {
		return err
	}

	return nil
}
