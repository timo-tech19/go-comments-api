package main

import (
	"fmt"

	"github.com/timo-tech19/go-comments-api/internal/comment"
	"github.com/timo-tech19/go-comments-api/internal/db"
	transportHttp "github.com/timo-tech19/go-comments-api/internal/transport/http"
)

// Run - responsible for startup and instantiation
// of our go application
func Run() error {
	fmt.Println("Starting app...")

	db, err := db.NewDatabase()

	if err != nil {
		fmt.Println("Failed to connect to database")
		return err
	}

	if err := db.MigrateDB(); err != nil {
		fmt.Println("Failed to migrate database")
		return err
	}

	cmtService := comment.NewService(db)

	httpHandler := transportHttp.NewHandler(cmtService)
	if err := httpHandler.Serve(); err != nil {
		return err
	}

	fmt.Println("Database connection and ping successful")
	return nil
}

func main() {
	fmt.Println("Go Comments REST API")

	// in go you can declare a var in a control statement
	if err := Run(); err != nil {
		fmt.Println(err)
	}
}
