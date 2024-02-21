package main

import (
	"fmt"
	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
	"github.com/go-chi/jwtauth"
	"github.com/obrunogonzaga/pos-go-expert/07-APIS/07-APIS/configs"
	_ "github.com/obrunogonzaga/pos-go-expert/07-APIS/07-APIS/docs"
	"github.com/obrunogonzaga/pos-go-expert/07-APIS/07-APIS/internal/entity"
	"github.com/obrunogonzaga/pos-go-expert/07-APIS/07-APIS/internal/infra/database"
	"github.com/obrunogonzaga/pos-go-expert/07-APIS/07-APIS/internal/infra/webserver/handlers"
	httpSwagger "github.com/swaggo/http-swagger"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"log"
	"net/http"
)

// @title 			Go Expert API de Produtos
// @version 		1.0
// @description 	Product API with authentication
// @termsOfService 	http://swagger.io/terms/

// @contact.name 	Bruno Gonzaga
// @contact.url 	http://www.brunogonzaga.dev
// @contact.email 	brunog86@gmail.com

// @license.name 	MIT
// @license.url 	http://mit.com

// @host			localhost:8000
// @BasePath /

// @securityDefinitions.apiKey ApiKeyAuth
// @in header
// @name Authorization
func main() {
	cfg, err := configs.LoadConfig(".")
	if err != nil {
		panic(err)
	}
	fmt.Println(cfg.DBDriver)

	db, err := gorm.Open(sqlite.Open("test.db"), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}
	db.AutoMigrate(&entity.Product{}, &entity.User{})

	productDB := database.NewProduct(db)
	productHandler := handlers.NewProductHandler(productDB)

	userDB := database.NewUser(db)
	userHandler := handlers.NewUserHandler(userDB, cfg.TokenAuth, cfg.JWTExpiresIn)

	r := chi.NewRouter()
	r.Use(middleware.Logger)
	r.Use(LogRequest)
	r.Use(middleware.Recoverer)

	r.Route("/products", func(r chi.Router) {
		r.Use(jwtauth.Verifier(cfg.TokenAuth))
		r.Use(jwtauth.Authenticator)
		r.Post("/", productHandler.CreateProduct)
		r.Get("/", productHandler.GetProducts)
		r.Get("/{id}", productHandler.GetProduct)
		r.Put("/{id}", productHandler.UpdateProduct)
		r.Delete("/{id}", productHandler.DeleteProduct)
	})

	r.Post("/users", userHandler.Create)
	r.Post("/users/generateToken", userHandler.GetJwt)

	r.Get("/docs/*", httpSwagger.Handler(httpSwagger.URL("http://localhost:8000/docs/doc.json")))

	http.ListenAndServe(":8000", r)
}

// LogRequest Criando um middleware
func LogRequest(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// r.Context().Value("user") poderia pegar valores do contexto
		log.Println(r.Method, r.URL)
		next.ServeHTTP(w, r)
	})
}
