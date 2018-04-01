package main

import (
  "os"
  "fmt"
  "flag"
  "sync"
  "github.com/steakknife/rsapss/subtle"
  "net/http"
  "database/sql"
  "encoding/json"
  /* "github.com/brianarpie/go-web-app-playground/config" */
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
  _, client := ensureSessionCookie(w, r)
  if client.loggedIn == false {
    http.Redirect(w, r, "/", 401)
    return
  }
  if client.loggedIn == true {
    h.wrappedHandler.ServeHTTP(w, r)
    return
  }
}

func authentication(h http.Handler) authMiddleware {
  return authMiddleware{h}
}

func main() {
  /* db = config.OpenDatabase() */
  r := mux.NewRouter()
  sessionStore = make(map[string]Client)
  portPtr := flag.String("port", "3000", "The Go Server's HTTP Port")
  flag.Parse() // execute command line parsing.
  port := fmt.Sprintf(":%s", *portPtr)

  r.Handle("/", http.FileServer(http.Dir("./views/")))
  r.PathPrefix("/static/").Handler(http.StripPrefix("/static/", http.FileServer(http.Dir("./static/"))))
  r.PathPrefix("/dist/").Handler(http.StripPrefix("/dist/", http.FileServer(http.Dir("./dist/"))))

  r.Handle("/login", LoginHandler).Methods("POST")
  r.Handle("/status", StatusHandler).Methods("GET")
  r.Handle("/products", authentication(ProductsHandler)).Methods("GET")
  r.Handle("/products/{slug}/feedback", authentication(AddFeedbackHandler)).Methods("GET")

  http.ListenAndServe(port, handlers.LoggingHandler(os.Stdout, r))

  /* defer db.Close() */
}

func ensureSessionCookie(w http.ResponseWriter, r *http.Request) (*http.Cookie, Client) {
  var present bool
  var client Client
  cookie, err := r.Cookie("session")
  if err != nil {
    if err != http.ErrNoCookie {
      fmt.Fprint(w, err)
      return nil, client
    } else {
      err = nil
    }
  }
  if cookie != nil {
    storageMutex.RLock()
    client, present = sessionStore[cookie.Value]
    storageMutex.RUnlock()
  } else {
    present = false
  }

  if present == false {
    cookie = &http.Cookie{
      Name: "session",
      Value: uuid.NewV4().String(),
    }
    client = Client{false}
    storageMutex.Lock()
    sessionStore[cookie.Value] = client
    storageMutex.Unlock()
  }

  http.SetCookie(w, cookie)

  return cookie, client
}

var LoginHandler = http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
  cookie, client := ensureSessionCookie(w, r)

  err := r.ParseForm()
  if err != nil {
    fmt.Fprint(w, err)
    return
  }

  if subtle.ConstantTimeCompare([]byte(r.FormValue("password")), []byte("abc123")) == 1 {
    client.loggedIn = true
    fmt.Fprintln(w, "Thank you for logging in.")
    storageMutex.Lock()
    sessionStore[cookie.Value] = client
    storageMutex.Unlock()
  } else {
    fmt.Fprintln(w, "Wrong password.")
  }
})

var NotImplemented = http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
  w.Write([]byte("Not Implemented"))
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

