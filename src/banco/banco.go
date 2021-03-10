package banco

import (
	"fmt"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql" // GORM
)

//Usuario a ser usado como modelo
type Usuario struct {
	gorm.Model
	Nome             string
	CPF              int64
	DatadeNascimento string
	Telefone         int64
	Email            string
	Rua              string
	Bairro           string
	Complemento      string
	Cidade           int64
}

//UsuarioBusca é o struct para mostrar o resultados, tendo cidade como string
type UsuarioBusca struct {
	gorm.Model
	Nome             string
	CPF              int64
	DatadeNascimento string
	Telefone         int64
	Email            string
	Rua              string
	Bairro           string
	Complemento      string
	Cidade           string
}
//Cidade struct modelo de ciadade
type Cidade struct {
	gorm.Model
	Id   int
	Nome string
	Uf   string
}

//DBClient é o export do bd
var DBClient *gorm.DB

//IniciarMigracaoBD Inicia conexao com bando
func IniciarMigracaoBD() {

	stringConexao := "golang:golang@/cadastroSimples?charset=utf8&parseTime=True&loc=Local"
	db, erro := gorm.Open("mysql", stringConexao)
	if erro != nil {
		fmt.Println(erro.Error())
		panic("Falha ao conectar ao banco de dados")
	}

	DBClient = db
	fmt.Println(db)
	db.AutoMigrate(&Usuario{})
	db.AutoMigrate(&Cidade{})
	db.AutoMigrate(&UsuarioBusca{})

}
