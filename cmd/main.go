package main

import (
	"database/sql"
	"fmt"
	"net/http"
	"os"
	"regexp"

	_ "github.com/mattn/go-sqlite3"

	"github.com/go-chi/chi"
	"github.com/vanyovan/test-product.git/internal/handler"
	"github.com/vanyovan/test-product.git/internal/repo"
	"github.com/vanyovan/test-product.git/internal/usecase"
)

func main() {
	db, err := sql.Open("sqlite3", getRootDirectory()+"/database.db")
	if err != nil {
		panic(err)
	}
	defer db.Close()

	walletRepo := repo.NewProductRepo(db)

	walletUsecase := usecase.NewProductService(walletRepo)

	handler := handler.NewHandler(walletUsecase)

	router := chi.NewRouter()

	router.Group(func(r chi.Router) {
		r.Method(http.MethodPost, "/api/v1/product", http.HandlerFunc(handler.HandleCreateProduct))   //create product
		r.Method(http.MethodGet, "/api/v1/product", http.HandlerFunc(handler.HandleViewProduct))      //read product
		r.Method(http.MethodDelete, "/api/v1/product", http.HandlerFunc(handler.HandleDeleteProduct)) //update product
		r.Method(http.MethodPatch, "/api/v1/product", http.HandlerFunc(handler.HandleUpdateProduct))  //update product
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
