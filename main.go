package main

import (
    "database/sql"
    "fmt"
    _ "github.com/go-sql-driver/mysql"
)

// Connection string for MySQL database
const (
    username = "username"
    password = "password"
    hostname = "localhost"
    port     = "3306"
    dbname   = "example_database"
)

// User struct representing data structure for user
type User struct {
    ID    int
    Name  string
    Email string
}

func main() {
    db, err := sql.Open("mysql", fmt.Sprintf("%s:%s@tcp(%s:%s)/%s", username, password, hostname, port, dbname))
    if err != nil {
        panic(err.Error())
    }
    defer db.Close()

    // Test CRUD operations
    // Create
    newUser := User{Name: "John Doe", Email: "john@example.com"}
    id, err := createUser(db, newUser)
    if err != nil {
        panic(err.Error())
    }
    fmt.Printf("New user ID: %d\n", id)

    // Read
    users, err := getUsers(db)
    if err != nil {
        panic(err.Error())
    }
    fmt.Println("All users:")
    for _, user := range users {
        fmt.Printf("ID: %d, Name: %s, Email: %s\n", user.ID, user.Name, user.Email)
    }

    // Update
    updatedUser := User{ID: id, Name: "Jane Doe", Email: "jane@example.com"}
    err = updateUser(db, updatedUser)
    if err != nil {
        panic(err.Error())
    }
    fmt.Println("User updated successfully")

    // Delete
    err = deleteUser(db, id)
    if err != nil {
        panic(err.Error())
    }
    fmt.Println("User deleted successfully")
}

// Function to create a new user
func createUser(db *sql.DB, user User) (int64, error) {
    result, err := db.Exec("INSERT INTO users(name, email) VALUES(?, ?)", user.Name, user.Email)
    if err != nil {
        return 0, err
    }
    id, err := result.LastInsertId()
    if err != nil {
        return 0, err
    }
    return id, nil
}

// Function to retrieve all users
func getUsers(db *sql.DB) ([]User, error) {
    rows, err := db.Query("SELECT id, name, email FROM users")
    if err != nil {
        return nil, err
    }
    defer rows.Close()

    var users []User
    for rows.Next() {
        var user User
        err := rows.Scan(&user.ID, &user.Name, &user.Email)
        if err != nil {
            return nil, err
        }
        users = append(users, user)
    }
    return users, nil
}

// Function to update user details
func updateUser(db *sql.DB, user User) error {
    _, err := db.Exec("UPDATE users SET name = ?, email = ? WHERE id = ?", user.Name, user.Email, user.ID)
    if err != nil {
        return err
    }
    return nil
}

// Function to delete a user
func deleteUser(db *sql.DB, id int64) error {
    _, err := db.Exec("DELETE FROM users WHERE id = ?", id)
    if err != nil {
        return err
    }
    return nil
}
package main

import (
    "database/sql"
    "fmt"
    _ "github.com/go-sql-driver/mysql"
)

// Connection string for MySQL database
const (
    username = "username"
    password = "password"
    hostname = "localhost"
    port     = "3306"
    dbname   = "example_database"
)

// User struct representing data structure for user
type User struct {
    ID    int
    Name  string
    Email string
}

func main() {
    db, err := sql.Open("mysql", fmt.Sprintf("%s:%s@tcp(%s:%s)/%s", username, password, hostname, port, dbname))
    if err != nil {
        panic(err.Error())
    }
    defer db.Close()

    // Test CRUD operations
    // Create
    newUser := User{Name: "John Doe", Email: "john@example.com"}
    id, err := createUser(db, newUser)
    if err != nil {
        panic(err.Error())
    }
    fmt.Printf("New user ID: %d\n", id)

    // Read
    users, err := getUsers(db)
    if err != nil {
        panic(err.Error())
    }
    fmt.Println("All users:")
    for _, user := range users {
        fmt.Printf("ID: %d, Name: %s, Email: %s\n", user.ID, user.Name, user.Email)
    }

    // Update
    updatedUser := User{ID: id, Name: "Jane Doe", Email: "jane@example.com"}
    err = updateUser(db, updatedUser)
    if err != nil {
        panic(err.Error())
    }
    fmt.Println("User updated successfully")

    // Delete
    err = deleteUser(db, id)
    if err != nil {
        panic(err.Error())
    }
    fmt.Println("User deleted successfully")
}

// Function to create a new user
func createUser(db *sql.DB, user User) (int64, error) {
    result, err := db.Exec("INSERT INTO users(name, email) VALUES(?, ?)", user.Name, user.Email)
    if err != nil {
        return 0, err
    }
    id, err := result.LastInsertId()
    if err != nil {
        return 0, err
    }
    return id, nil
}

// Function to retrieve all users
func getUsers(db *sql.DB) ([]User, error) {
    rows, err := db.Query("SELECT id, name, email FROM users")
    if err != nil {
        return nil, err
    }
    defer rows.Close()

    var users []User
    for rows.Next() {
        var user User
        err := rows.Scan(&user.ID, &user.Name, &user.Email)
        if err != nil {
            return nil, err
        }
        users = append(users, user)
    }
    return users, nil
}

// Function to update user details
func updateUser(db *sql.DB, user User) error {
    _, err := db.Exec("UPDATE users SET name = ?, email = ? WHERE id = ?", user.Name, user.Email, user.ID)
    if err != nil {
        return err
    }
    return nil
}

// Function to delete a user
func deleteUser(db *sql.DB, id int64) error {
    _, err := db.Exec("DELETE FROM users WHERE id = ?", id)
    if err != nil {
        return err
    }
    return nil
}
