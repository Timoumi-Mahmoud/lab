package main

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"
	"os"

	"berserk.com/m/models"
	"github.com/gorilla/websocket"
	_ "github.com/lib/pq"
)

type Env struct {
	books interface {
		AllBooks() ([]models.Book, error)
	}

	infoLog  *log.Logger
	errorLog *log.Logger
}

var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
}

func main() {
	var err error
	fmt.Println("rolling")
	fmt.Println("drivers: ", sql.Drivers())

	db, err := sql.Open("postgres", "postgresql://root:password@localhost:5433/books?sslmode=disable")
	if err != nil {
		log.Fatal("Unable to establish connection with database: ", err)
	}

	infoLog := log.New(os.Stdout, "INFO\t", log.Ldate|log.Ltime|log.Lshortfile)
	errorLog := log.New(os.Stdout, "ERROR\t", log.Ldate|log.Ltime|log.Lshortfile)

	// Initalise Env with models.BookModel instance (which in turn warps the connection pool).
	env := &Env{
		books:    models.BookModel{DB: db},
		infoLog:  infoLog,
		errorLog: errorLog,
	}

	http.HandleFunc("/books", env.booksList)
	http.HandleFunc("/echo", env.echo)

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, "websockets.html")
	})
	http.ListenAndServe(":6969", nil)
}

func (env *Env) echo(w http.ResponseWriter, r *http.Request) {
	/*	conn, err := upgrader.Upgrade(w, r, nil)
		if err != nil {
			env.errorLog.Println(err)
		}
		env.infoLog.Println("connection.............", conn)

		for {
			msgType, msg, err := conn.ReadMessage()
			if err != nil {
				env.errorLog.Println(err)
			}
			fmt.Printf("%s sent: %s\n", conn.RemoteAddr(), string(msg))
			if err = conn.WriteMessage(msgType, msg); err != nil {
				return
			}

		}

		defer conn.Close()
		fmt.Fprintf(w, "Echo ...........;")
	*/
	upgrader.CheckOrigin = func(r *http.Request) bool { return true }

	// upgrade this connection to a WebSocket
	// connection
	ws, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Println(err)
	}

	log.Println("Client Connected")
	err = ws.WriteMessage(1, []byte("Hi Client!"))
	if err != nil {
		log.Println(err)
	}
	// listen indefinitely for new messages coming
	// through on our WebSocket connection
	reader(ws)

}
func reader(conn *websocket.Conn) {
	for {
		// read in a message
		messageType, p, err := conn.ReadMessage()
		if err != nil {
			log.Println(err)
			return
		}
		// print out that message for clarity
		fmt.Println(string(p))

		if err := conn.WriteMessage(messageType, p); err != nil {
			log.Println(err)
			return
		}

	}
}

func (env *Env) booksList(w http.ResponseWriter, r *http.Request) {
	bks, err := env.books.AllBooks()
	if err != nil {
		log.Print(err)
		http.Error(w, http.StatusText(500), 500)
		return
	}
	for _, bk := range bks {
		fmt.Fprintf(w, "%s, %s, %s, £%.2f\n", bk.Isbn, bk.Title, bk.Author, bk.Price)
	}

}
