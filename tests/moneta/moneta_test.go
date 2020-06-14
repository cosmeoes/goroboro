package tests

import (
	"fmt"
	"os"
	"strconv"
	"strings"
	"testing"

	"github.com/cosmeoes/goroboro/config"
	"github.com/cosmeoes/goroboro/moneta"
	yaml "gopkg.in/yaml.v2"
)

type User struct {
    Id int `moneta:"id"`
    Name string `moneta:"nombre"` 
    Password string `moneta:"password"`
    Username string `moneta:"username"`
    Ap string `moneta:"ap"`
    Am string `moneta:"am"`
    Email string `moneta:"email"`
}

func (User) TableName() string {
    return "users"
}

var insertedIds []int64;

func TestInsert(t *testing.T) {
    dbConfig := config.GetDBConfig()
    f, err := os.Open("/home/cosme/Documents/projects/goroboro/framework/tests/database.yml")
    fmt.Println("lmao 1", dbConfig)
    if err != nil {
        //TODO: handle not file error
        fmt.Println(err)
        return
    }
    decoder := yaml.NewDecoder(f)
    err = decoder.Decode(dbConfig)

    newUser := User{Name: "Cosme", Password: "123", Email: "comeoes@gmail.com"}
    newUser2 := User{Name: "Damian", Password: "324", Email: "damian050697@gmail.com"}
    result, err := moneta.Model(newUser).Save()

    if err != nil {
        t.Error("Save method returned err: {}", err)
        return
    }
    id, err := result.LastInsertId()
    insertedIds = append(insertedIds, id)

    result, err = moneta.Model(newUser2).Save()
    if err != nil {
        t.Error("Save method returned err: {}", err)
        return 
    }

    id, err = result.LastInsertId()
    insertedIds = append(insertedIds, id)

    rows := moneta.Find(User{}).Get()
    if len(rows) != 2 {
        t.Error("There should be 2 users, only {} found", len(rows))
    }
}

func TestColumn(t *testing.T) {
    result := moneta.Find(User{}).Get("nombre")

    user := result[0].(User)
    if user.Password != "" {
        t.Error("When getting a column other columns should be empty")
    }
    fmt.Println(user.Password)
}
func TestWorks(t *testing.T) {
    result := moneta.Find(User{}).First()

    _, ok := result.(User)
    if !ok {
        t.Error("Result not of type user")
    }
}

func TestUpdate(t *testing.T) {
    idString := ""
    for _, id := range insertedIds {
       idString += strconv.FormatInt(id, 10) + ", "
    }

    idString = strings.Trim(idString, ", ")
    users := moneta.Find(User{}).Where("id", "in", "(" + idString + ")").Get()

    user := users[0].(User)
    newEmail := "updated@gmail.com"
    user.Email = newEmail

    result, err := moneta.Model(user).Save()
    if err != nil {
        fmt.Println("lmao", err)
        return
    }

    id, _ := result.LastInsertId()
    fmt.Println("Affected", id)
    userUpdated := moneta.Find(User{}).Where("email", "=", "'" + newEmail + "'").First().(User)
    
    if user.Id != userUpdated.Id {
        fmt.Println(user.Id, userUpdated.Id)
        t.Error("User wasnt updated");
        t.Fail()
    }

}

func TestDelete(t *testing.T) {
    idString := ""
    for _, id := range insertedIds {
       idString += strconv.FormatInt(id, 10) + ", "
    }

    idString = strings.Trim(idString, ", ")
    users := moneta.Find(User{}).Where("id", "in", "(" + idString + ")").Get()

    for _, user := range users {
        result, err := moneta.Model(user.(User)).Delete()
        if err != nil {
            t.Error("Delete returned an error", err)
            return
        }

        affected, err := result.RowsAffected()
        fmt.Println("Rows deleted", affected)
    }
}

func TestMysqlGrammar(t *testing.T) {
}
