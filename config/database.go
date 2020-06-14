package config

import (
	"fmt"
	"strings"
	"sync"
)

type Database struct {
    Driver string `yaml:"driver"`
    Username string `yaml:"username"`
    Password string `yaml:"password"`
    Database string `yaml:"database"`
    Host string `yaml:"host"`
}

var instance *Database
var once sync.Once

func GetDBConfig() *Database {
    once.Do(func() {
        instance = &Database{}
    })
    return instance;
}

func (d *Database) ToDataSource() string {
    if (d.Host != "" && !strings.HasPrefix(d.Host, "tcp")) {
        d.Host = "tcp(" + d.Host + ")"
    }
    return fmt.Sprintf("%v:%v@%v/%v", d.Username, d.Password, d.Host, d.Database)
}
