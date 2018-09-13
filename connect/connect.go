package connect

import (

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"log"
)
var connection *gorm.DB

const engine_sql string = "mysql"
const username string = "root"
const password string = "root"
const database string = "tabla1"

func InitializeDatabase(){
	connection= ConnectORM(CreateString() )
	log.Println("La conexcion con la base de datos")
}

func CloseConnection()  {
	connection.Close()
	log.Println("Esta cerrada")
}

func ConnectORM(stringConnection string)  *gorm.DB{
	connection, err := gorm.Open(engine_sql, stringConnection)
	if err != nil {
		log.Fatal(err)
		return nil
	}
	return connection
}

func CreateString()  string{
	return username + ":" + password + "@tcp(104.236.20.127:3306)/" + database
}