## Organising Database Access:
Three alternatives when it comes to database accessing in go:
1. Using a global variable:
    - Initialise the sql.DB connection pool in main.go func, assign it to a global variable, and then access the global from anywhere that need to execute a database query.
    ```go
    // main.go
        models.DB, err = sql.Open("postgres", "dns")
        bks , err := models.go 
    // models.go
        var DB *sql.DB
        func getAllBooks() {
            rows, err := DB.Query("select ...")
    ```
    - Good for Small application , and keeping track of globals in my head isn't a problem.
    - Http handlers are spread across multiple packages, but all database-related code lives in one package.
    - When I don't need to mock the database for testing.
see ==> 

    - There is a possibility to use initialisation func to set up connection pool (good for test database ,all database-realted code lives in a single package )
    ```
    func InitDB(dataSourceName string) error{
        db, errsql.Open()
    }
    // main.go
        err := models.InitDB("dns")
    ```


2. Using Dependency Injection
In more complex app , there are probably additional application-level objects That I want to access such logger, template cache extra 
rather than storing all these dependencies in global variable, a neat approach is to store them in a single custom Env struct
```go
    type Env struct{
        db *sql.DB
        logger *log.Logger
        templates *templates.Template
    }
```
    - Clarity  all the dependencies for my handlers are explicitly in one place (Env struct in the main func)
    - The ability to add another connection pool to a test database in Env 
    - All my http handlers live in one package but database-realted code may spread across multiple packages
    - I don't need to mock the database
* Using the same approach I can use closure in DI
```go
        http.Handle("/books", booksIndex(env))
// Use a closure to make Env available to the handler logic.
func booksIndex(env *Env) http.HandlerFunc {
    return func(w http.ResponseWriter, r *http.Request) {
            bks, err := models.AllBooks(env.db)
            ..... }
```

3. Wrapping the connection pool
Dependency injection but wrapping the sql.DB connection pool in my own custom type.
```go
type BookModel struct {
    DB *sql.DB
    }
    func (m BookModel) All() ([]Book, error) {
        rows, err := m.DB.Query("SELECT * FROM books")
        ....}
    // main.go
    type Env struct{
        books models.BookModel
        }
        db, err := sql.open("","")
        env :=&Env{
            books: models.BookModel{DB: db}
            }
            func (env *Env) booksList(w http.ResponseWriter, r *http.Request) {
            bks, err := env.books.GetAll()
            .....}
```
    - Good for complex application
    - Ability to create mock implementation to my BookModel which I can use during testing


