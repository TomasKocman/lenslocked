package main

import (
	"fmt"
	"lenslocked/middleware"
	"lenslocked/models"
	"net/http"

	"lenslocked/controllers"

	"github.com/gorilla/mux"
)

const (
	host   = "localhost"
	port   = 5432
	user   = "postgres"
	dbname = "lenslocked_dev"
)

func notFound(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html")
	fmt.Fprint(w, "Ops.. page not found")
}

func main() {
	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s dbname=%s sslmode=disable", host, port, user, dbname)

	services, err := models.NewServices(psqlInfo)
	if err != nil {
		panic(err)
	}
	defer services.Close()

	services.AutoMigrate()

	staticC := controllers.NewStatic()
	usersC := controllers.NewUsers(services.User)
	galleriesC := controllers.NewGalleries(services.Gallery)

	requireUserMw := middleware.RequireUser{UserService: services.User}

	r := mux.NewRouter()

	r.Handle("/", staticC.Home).Methods("GET")
	r.Handle("/contact", staticC.Contact).Methods("GET")
	r.Handle("/faq", staticC.FAQ).Methods("GET")

	r.HandleFunc("/signup", usersC.New).Methods("GET")
	r.HandleFunc("/signup", usersC.Create).Methods("POST")

	r.Handle("/login", usersC.LoginView).Methods("GET")
	r.HandleFunc("/login", usersC.Login).Methods("POST")

	r.HandleFunc("/cookietest", usersC.CookieTest).Methods("GET")

	newGallery := requireUserMw.Apply(galleriesC.New)
	createGallery := requireUserMw.ApplyFn(galleriesC.Create)
	r.Handle("/galleries/new", newGallery).Methods("GET")
	r.HandleFunc("/galleries", createGallery).Methods("POST")
	r.HandleFunc("/galleries/{id:[0-9]+}", galleriesC.Show).Methods("GET")

	r.NotFoundHandler = http.HandlerFunc(notFound)

	fmt.Println("Starting the server on :3000...")
	http.ListenAndServe(":3000", r)
}
