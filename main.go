package main

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/atharvam2004/rss-go/internal/database"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/cors"
	_ "github.com/lib/pq"
	//"github.com/go-chi/chi/v5/middleware"
	"github.com/joho/godotenv"
)
type apiConfig struct{
	DB *database.Queries
}
func main() {

	godotenv.Load()
	port := os.Getenv("PORT")
	if port == "" {
		log.Fatal("port not gound in env")
	}
	dbURL:= os.Getenv("DATABASE_URL")
	if dbURL == "" {
		log.Fatal("port not gound in env")
	}
	conn,err:=sql.Open("postgres",dbURL)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("shree ganesh", port)
	queries2:=database.New(conn)

	if err!=nil{
		log.Fatal("cant connect",err)
	}
	apiCfg:=apiConfig{
		DB:queries2,
	}

	router := chi.NewRouter()
	srv:=&http.Server{
		Handler:router,
		Addr: ":"+port,
	}
	router.Use(cors.Handler(cors.Options{
		// AllowedOrigins:   []string{"https://foo.com"}, // Use this to allow specific origin hosts
		AllowedOrigins:   []string{"https://*", "http://*"},
		// AllowOriginFunc:  func(r *http.Request, origin string) bool { return true },
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowedHeaders:   []string{"Accept", "Authorization", "Content-Type", "X-CSRF-Token"},
		ExposedHeaders:   []string{"Link"},
		AllowCredentials: false,
		MaxAge:           300, // Maximum value not ignored by any of major browsers
	}))
	  v1Router := chi.NewRouter()
	  v1Router.Get("/ready",handle)
	  v1Router.Get("/err",handleerr)
	  v1Router.Post("/users",apiCfg.handlerCreateUser)
	  v1Router.Get("/users",apiCfg.handlerGetUser)
	  router.Mount("/v1", v1Router)
	
	err2:=srv.ListenAndServe()
	if err!=nil{
		fmt.Println(err2)
	}else{
		fmt.Println("Server started")
	}
}
