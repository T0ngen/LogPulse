package sqlc

import (
	
	"database/sql"
	"fmt"

	

	_ "github.com/lib/pq"
)




func OpenPostgresConnection() (*sql.DB, error) {
    connStr := "user=root password=1234 dbname=Webloggertg sslmode=disable port=5436"

    db, err := sql.Open("postgres", connStr)
    if err != nil {
        return nil, fmt.Errorf("failed to connect to database: %v", err)
    }

    err = db.Ping()
    if err != nil {
        db.Close()
        return nil, fmt.Errorf("failed to ping database: %v", err)
    }

    return db, nil
}
