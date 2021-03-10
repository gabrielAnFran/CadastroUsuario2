package main

import (
	"cadastro/src/banco"
	"cadastro/src/crud"
	"log"

	"github.com/gin-gonic/gin"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

func main() {
	r := gin.Default()
	banco.IniciarMigracaoBD()

	r.POST("/usuarios", crud.CreateUser)
	r.GET("/usuarios", crud.GetUsers)
	r.GET("/usuarios/nome", crud.GetUserByName)
	r.PUT("/usuarios/id", crud.PutUser)
	r.DELETE("/usuarios/delete", crud.DeleteUser)

	if erro := r.Run(":8081"); erro != nil {
		log.Fatal(erro.Error())
	}

}
