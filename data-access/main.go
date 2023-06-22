package main
import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/go-sql-driver/mysql"
)

func main() {
	db, err := sql.Open("mysql", "root:root@tcp(127.0.0.1:3306)/reccording")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	// Check if the connection was successful
	if err := db.Ping(); err != nil {
		log.Fatal(err)
	}

	// Perform database operations
	albums := albumsByArtist(db, "Gerry Mulligan")
	fmt.Println(albums)
}

func albumsByArtist(db *sql.DB, artist string) []string {
	rows, err := db.Query("SELECT artist FROM album WHERE artist = ?", artist)
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()

	var albums []string
	for rows.Next() {
		var album string
		if err := rows.Scan(&album); err != nil {
			log.Fatal(err)
		}
		
		fmt.Print(album);
		albums = append(albums, album)
	}

	if err := rows.Err(); err != nil {
		log.Fatal(err)
	}

	return albums
}
