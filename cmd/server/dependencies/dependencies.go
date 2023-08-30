package dependencies

import (
	"app/cmd/server/handlers"
	"app/internal/products/storage"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

// NewApp returns a new instance of App.
func NewApp(serverAddr string) *App {
	return &App{
		ServerAddr: serverAddr,
	}
}

// App is a struct that represents the application.
type App struct {
	// Addr is the address to listen on.
	ServerAddr string
}

func (a *App) Run() (err error) {
	// dependencies
	// -> storage
	storageProduct := storage.NewStorageProductInMemory()
	// -> controllers
	controllerProduct := handlers.NewControllerProduct(storageProduct)

	// server
	r := chi.NewRouter()
	// -> middlewares
	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)
	// -> routes / endpoints
	r.Get("/products", controllerProduct.GetProducts())

	// -> run
	err = http.ListenAndServe(a.ServerAddr, r)
	if err != nil {
		return
	}

	return
}