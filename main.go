package main

import (
  "os"
  "fmt"
  "time"
  "net/http"
  "database/sql"
  "encoding/json"
  "github.com/brianarpie/go-web-app-playground/config"
  "github.com/brianarpie/go-web-app-playground/models"

  "github.com/gorilla/mux"
  "github.com/gorilla/handlers"
  "github.com/satori/go.uuid"
)

var db *sql.DB

// for now hard code seed the model instead of creating a database wrapper
var products = []models.Product{
  models.Product{Id: 1, Name: "Best Cheese Knife Ever", Slug: "best-cheese-knife-ever", Description: "A++" },
  models.Product{Id: 2, Name: "Rusty Nail File", Slug: "rusty-nail-file", Description: "Hepatitis Free" },
  models.Product{Id: 3, Name: "Cantelope Scooper", Slug: "cantelope-scooper-1", Description: "Melon Balls Galore" },
}

type sessionRequest struct { 
  Email string `json:"email"`
  Password string `json:"password"`
}

type sessionResponse struct {
  Ok string `json:"ok"`
}

var sessionStore map[string]Client
var storageMutex sync.RWMutex

type Client struct {
  loggedIn bool
}

type authMiddleware struct {
  wrappedHandler http.Handler
}

func (h authMiddleware) ServeHTTP(w http.ResponseWriter, r *http.Request) {
}

func authentication(h http.Handler) authMiddleware {
  return authMiddleware{h}
}

func main() {
  db = config.OpenDatabase()
  r := mux.NewRouter()
  sessionStore = make(map[string]Client)

  r.Handle("/", http.FileServer(http.Dir("./views/")))
  // TODO: mess around with status files
  r.PathPrefix("/static/").Handler(http.StripPrefix("/static/", http.FileServer(http.Dir("./static/"))))

  r.Handle("/login", LoginHandler).Methods("POST")
  r.Handle("/status", StatusHandler).Methods("GET")
  r.Handle("/products", authMiddleware(ProductsHandler)).Methods("GET")
  r.Handle("/products/{slug}/feedback", authMiddleware(AddFeedbackHandler)).Methods("GET")
  r.Handle("/get-token", GetTokenHandler).Methods("GET")

  http.ListenAndServe(":3000", handlers.LoggingHandler(os.Stdout, r))

  defer db.Close()
}


var NotImplemented = http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
  w.Write([]byte("Not Implemented"))
})

var authMiddleware(h http.Handler) http.Handler {
}

var LoginHandler = http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
})

var StatusHandler = http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
  w.Write([]byte("API is up and running"))
})

var ProductsHandler = http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
  payload, _ := json.Marshal(products)
  w.Header().Set("Content-Type", "application/json")
  w.Write([]byte(payload))
})

var AddFeedbackHandler = http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
  var product models.Product
  vars := mux.Vars(r)
  slug := vars["slug"]

  for _, p := range products {
    if p.Slug == slug {
      product = p
    }
  }

  w.Header().Set("Content-Type", "application/json")
  if product.Slug != "" {
    payload, _ := json.Marshal(product)
    w.Write([]byte(payload))
  } else {
    w.Write([]byte("Product Not Found"))
  }
})

// NOTE: educational purposes only
// var mySigningKey = []byte("secret")
//
// var GetTokenHandler = http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
//   token := jwt.New(jwt.SigningMethodHS256)
//   claims := token.Claims.(jwt.MapClaims)
//   claims["admin"] = true
//   claims["name"] = "Brian Arpie"
//   claims["exp"] = time.Now().Add(time.Hour * 24).Unix()
//   tokenString, _ := token.SignedString(mySigningKey)
//   w.Write([]byte(tokenString))
// })

// var jwtMiddleware = jwtmiddleware.New(jwtmiddleware.Options{
//   ValidationKeyGetter: func(token *jwt.Token) (interface{}, error) {
//     return mySigningKey, nil
//   },
//   SigningMethod: jwt.SigningMethodHS256,
// })

// func sessionsHandler(w http.ResponseWriter, r *http.Request) {
//   switch r.Method {
//   case "GET":

//   default:
//     // no-op - error message
//   }
// }

