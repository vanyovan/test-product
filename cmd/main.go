package main

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"
	"os"
	"regexp"

	_ "github.com/mattn/go-sqlite3"

	"github.com/go-chi/chi"
	"github.com/vanyovan/test-product.git/internal/handler"
	"github.com/vanyovan/test-product.git/internal/repo"
	"github.com/vanyovan/test-product.git/internal/usecase"
)

func initializeDatabase(db *sql.DB) {
	schema := `
    CREATE TABLE mst_product (
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		name TEXT,
		description TEXT,
		price REAL,
		variety TEXT,
		rating REAL,
		stock INTEGER
	);
    `
	_, err := db.Exec(schema)
	if err != nil {
		log.Fatalf("Failed to initialize database schema: %v", err)
	}
}

func main() {
	db, err := sql.Open("sqlite3", getRootDirectory()+"/database.db")

	if err != nil {
		panic(err)
	}
	defer db.Close()

	initializeDatabase(db)

	productRepo := repo.NewProductRepo(db)

	productUsecase := usecase.NewProductService(productRepo)

	handler := handler.NewHandler(productUsecase)

	router := chi.NewRouter()

	router.Group(func(r chi.Router) {
		r.Method(http.MethodPost, "/api/v1/product", http.HandlerFunc(handler.HandleCreateProduct))        //create product
		r.Method(http.MethodGet, "/api/v1/product", http.HandlerFunc(handler.HandleViewProduct))           //read product
		r.Method(http.MethodDelete, "/api/v1/product/{id}", http.HandlerFunc(handler.HandleDeleteProduct)) //delete product
		r.Method(http.MethodPatch, "/api/v1/product/{id}", http.HandlerFunc(handler.HandleUpdateProduct))  //update product
	})

	server := &http.Server{
		Addr:    ":8080",
		Handler: router,
	}

	fmt.Println("server listening on", server.Addr)
	server.ListenAndServe()
}

func getRootDirectory() string {
	projectName := regexp.MustCompile(`^(.*test-product)`)
	currentWorkDirectory, _ := os.Getwd()
	rootPath := projectName.Find([]byte(currentWorkDirectory))

	return string(rootPath)
}
