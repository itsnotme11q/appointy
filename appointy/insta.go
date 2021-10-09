package main

import (
	"fmt"
	"time"
	// "log"
	"net/http"
	"github.com/gorilla/mux"
	"html/template"
	// "regexp"
	"go.mongodb.org/mongo-driver/bson"
    "go.mongodb.org/mongo-driver/mongo"
    "go.mongodb.org/mongo-driver/mongo/options"
)

var rnd *renderer.Render 
var db *mgo.Database

const(
	hostName	string  = "localhost:27017"
	dbname		string  = "demo_todo"
	collects	string  = "todo"
	port		string  = "19000"
	"db_auth": true 
)

type(
	User struct{
		Username		bson.ObjectId `bson:"_id,omitempty`
		Name   string `bson:"title"`
		email   string `bson:"email"`
		Password   string `bson:"password"`
	}
	Posts struct{
		ID		bson.ObjectId `bson:"_id,omitempty`
		Content   string `bson:"content"`
		Imageurl   string `bson:"imageurl"`
		createdAt time.Time `bson:"createdAt"`
	}
)

func init(){
	rnd = renderer.New()
	sess,err = mgo.Dial(hostname)
	checkerr(err)
	sess.SetMode(mgo.Monotonic,true )
	db = sess.DB(dbname)
}


func HomeHandler(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		http.Error(w, "404 not found. oh no home", http.StatusNotFound)
		return
	}
	switch r.Method {
	case "GET":		
		http.ServeFile(w, r, "home.html")
	case "POST":
		
	default:
		fmt.Fprintf(w, "Sorry, only GET and POST methods are supported.")
	}
}

func Users(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/users" {
		http.Error(w, "404 not found. oh no users", http.StatusNotFound)
		return
	}
	switch r.Method {
	case "GET":		
		http.ServeFile(w, r, "users.html")
	case "POST":
		username := r.FormValue("id")
		name := r.FormValue("name")
		email := r.FormValue("email")
		password := r.FormValue("password")
		fmt.Println(w, "userName = %s\n", username)
		fmt.Println(w, "Name = %s\n", name)
		fmt.Println(w, "email = %s\n", email)
		fmt.Println(w, "password = %s\n", password)	
		http.Redirect(w, r, "/users",http.StatusSeeOther)
	default:
		fmt.Fprintf(w, "Sorry, only GET and POST methods are supported.")
	}
}

func Login(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/login" {
		http.Error(w, "404 not found. oh no login", http.StatusNotFound)
		return
	}
	switch r.Method {
	case "GET":		
		http.ServeFile(w, r, "login.html")
	case "POST":
		username := r.FormValue("id")
		password := r.FormValue("password")
		fmt.Println(w, "userName = %s\n", username)
		fmt.Println(w, "password = %s\n", password)	
		p := &Page{Title: "bbbbb"}
		t, _ := template.ParseFiles("users_id.html")
		t.Execute(w, p)
		http.Redirect(w, r, "/users/" + username,http.StatusSeeOther)
	default:
		fmt.Fprintf(w, "Sorry, only GET and POST methods are supported.")
	}
}

func Users_id(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case "GET":		
		http.ServeFile(w, r, "users_id.html")
	case "POST":
		
	default:
		fmt.Fprintf(w, "Sorry, only GET and POST methods are supported.")
	}
}

func Posts(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/posts" {
		http.Error(w, "404 not found. oh no posts", http.StatusNotFound)
		return
	}
	switch r.Method {
	case "GET":		
		http.ServeFile(w, r, "posts.html")
	case "POST":
		id := r.FormValue("id")
		caption := r.FormValue("password")
		imageurl := r.FormValue("imageurl")
		now := time.Now() 
		fmt.Println(w, "id = %s\n", id)
		fmt.Println(w, "caption = %s\n", caption)
		fmt.Println(w, "imageurl = %s\n", imageurl)
		fmt.Println(w, "now = %s\n", now)		
		http.Redirect(w, r, "/posts",http.StatusSeeOther)
	default:
		fmt.Fprintf(w, "Sorry, only GET and POST methods are supported.")
	}
}

func Posts_id(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case "GET":		
		http.ServeFile(w, r, "posts_id.html")
	case "POST":
		
	default:
		fmt.Fprintf(w, "Sorry, only GET and POST methods are supported.")
	}
}


func Posts_users_id(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case "GET":		
		http.ServeFile(w, r, "posts_users_id.html")
	case "POST":
		
	default:
		fmt.Fprintf(w, "Sorry, only GET and POST methods are supported.")
	}
}

func main() {

	r := mux.NewRouter()
	r.HandleFunc("/", HomeHandler)
	r.HandleFunc("/users", Users)
	r.HandleFunc("/login", Login)
	r.HandleFunc("/users/{id}", Users_id)
	r.HandleFunc("/posts",Posts)
	r.HandleFunc("/posts/{id}",Posts_id)
	r.HandleFunc("/posts/users/{id}",Posts_users_id)
	http.Handle("/", r)
	http.Handle("/{id}", r)
	http.ListenAndServe(":8080", nil)

	client, err := mongo.NewClient(options.Client().ApplyURI("<MONGODB_URI>"))
    if err != nil {
        log.Fatal(err)
    }
    ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
    err = client.Connect(ctx)
    if err != nil {
        log.Fatal(err)
    }
    defer client.Disconnect(ctx)
    databases, err := client.ListDatabaseNames(ctx, bson.M{})
    if err != nil {
        log.Fatal(err)
    }
    fmt.Println(databases)
}