package server

import (
	"clases/internal/config"
	"clases/internal/db"
	"clases/internal/handler"
	"clases/internal/repository"
	"clases/internal/service"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

// NewServerChi is a function that returns a new instance of ServerChi
func NewServerChi(cfg *config.Config) *ServerChi {
	// default values
	defaultConfig := &config.Config{
		ServerAddress: ":8080",
	}
	if cfg != nil {
		if cfg.ServerAddress != "" {
			defaultConfig.ServerAddress = cfg.ServerAddress
		}
		if cfg.DBHost != "" {
			defaultConfig.DBHost = cfg.DBHost
		}
		if cfg.DBPort != "" {
			defaultConfig.DBPort = cfg.DBPort
		}
		if cfg.DBUser != "" {
			defaultConfig.DBUser = cfg.DBUser
		}
		if cfg.DBPassword != "" {
			defaultConfig.DBPassword = cfg.DBPassword
		}
		if cfg.DBName != "" {
			defaultConfig.DBName = cfg.DBName
		}
	}

	return &ServerChi{
		serverAddress: defaultConfig.ServerAddress,
		config:        defaultConfig,
	}
}

// ServerChi is a struct that implements the Application interface
type ServerChi struct {
	// serverAddress is the address where the server will be listening
	serverAddress string
	config        *config.Config
}

// Run is a method that runs the application
func (a *ServerChi) Run() (err error) {
	cfg := config.LoadConfig()

	database := db.ConnectDB(cfg)
	defer database.Close()
	// - repository
	pr := repository.NewProductRepository(database)
	if err != nil {
		return
	}
	// - service
	ps := service.NewProductService(pr)
	// - handler
	ph := handler.NewProductHandler(ps)

	// router
	r := chi.NewRouter()

	// - middlewares
	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)

	// - endpoints

	// Health check endpoint
	r.Group(func(r chi.Router) {
		r.Get("/ping", func(w http.ResponseWriter, r *http.Request) {
			w.Write([]byte("pong"))
		})
	})

	r.Route("/products", func(r chi.Router) {

		// Public endpoints
		r.Group(func(r chi.Router) {
			r.Get("/", ph.GetAllProducts())
			r.Get("/{id}", ph.GetProductByID)
			r.Post("/", ph.CreateProduct)
			r.Put("/{id}", ph.Update)
			r.Delete("/{id}", ph.DeleteProduct)

		})

	})

	// run server
	err = http.ListenAndServe(a.serverAddress, r)
	return
}
