package crud

import (
	"cadastro/src/banco"
	"fmt"
	"github.com/badoux/checkmail"
	"github.com/gin-gonic/gin"
	_ "github.com/jinzhu/gorm/dialects/mysql" //
	"net/http"
	"strconv"
)

//CreateUser cria um usuario
func CreateUser(c *gin.Context) {

	nome := c.Query("nome")
	cpf := c.Query("cpf")
	datadeNascimento := c.Query("data")
	telefone := c.Query("tel")
	email := c.Query("email")
	rua := c.Query("rua")
	bairro := c.Query("bairro")
	comp := c.Query("complemento")
	cid := c.Query("id")

	cidID, erro := strconv.Atoi(cid)
	if erro != nil {
		fmt.Println("ERRO AO CONVERTER")
	}

	cpff, erro := strconv.Atoi(cpf)
	if erro != nil {
		fmt.Println("erro ao converter")
	}

	tel, erro := strconv.Atoi(telefone)
	if erro != nil {
		fmt.Println("erro ao converter")
	}

	err := checkmail.ValidateFormat(email)
	if err != nil {
		fmt.Println(err)
		c.JSON(http.StatusNotAcceptable, gin.H{
			"message": "Email invalido",
		})

	} else if len(cpf) != 9 {
		c.JSON(http.StatusNotAcceptable, gin.H{
			"message": "CPF invalido",
		})

	} else {

		if erro := banco.DBClient.Create(&banco.Usuario{Nome: nome, CPF: int64(cpff), DatadeNascimento: datadeNascimento, Telefone: int64(tel),
			Email: email, Rua: rua, Bairro: bairro, Complemento: comp, Cidade: int64(cidID)}); erro != nil {
			fmt.Println(erro.Error)
		}

		var cidade banco.Cidade
		banco.DBClient.First(&cidade, cid)
		banco.DBClient.Raw("SELECT Nome, Uf FROM cidades WHERE Id = ?", cidID).Scan(&cidade)

		if erro := banco.DBClient.Create(&banco.UsuarioBusca{Nome: nome, CPF: int64(cpff), DatadeNascimento: datadeNascimento, Telefone: int64(tel),
			Email: email, Rua: rua, Bairro: bairro, Complemento: comp, Cidade: cidade.Nome}); erro != nil {
			fmt.Println(erro.Error)
		}

	}
	c.JSON(http.StatusCreated, gin.H{
		"message": "USUÁRIO CRIADO",
	})
}

//GetUsers query por todos usuarios 
func GetUsers(c *gin.Context) {

	var usuarios []banco.UsuarioBusca
	if erro := banco.DBClient.Find(&usuarios); erro.Error != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"mensagem": "usuario nao encontrado",
		})
	}
	c.JSON(http.StatusOK, usuarios)
}

//GetUserByName busca usuario por nome
func GetUserByName(c *gin.Context) {
	nome := c.Query("nome")

	var usuario banco.UsuarioBusca

	if erro := banco.DBClient.Where("nome LIKE ?", nome).Find(&usuario); erro.Error != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"mensagem": "usuario nao encontrado",
		})

	} else {
		c.JSON(http.StatusOK, usuario)

	}

}

//PutUser atualiza um usuario pela id
func PutUser(c *gin.Context) {
	id := c.Query("id")
	nome := c.Query("nome")
	cpf := c.Query("cpf")
	datadeNascimento := c.Query("data")
	telefone := c.Query("tel")
	email := c.Query("email")
	rua := c.Query("rua")
	bairro := c.Query("bairro")
	comp := c.Query("complemento")
	cid := c.Query("cid")

	cidID, erro := strconv.Atoi(cid)
	if erro != nil {
		fmt.Println("ERRO AO CONVERTER")
	}

	cpff, erro := strconv.Atoi(cpf)
	if erro != nil {
		fmt.Println("erro ao converter")
	}

	tel, erro := strconv.Atoi(telefone)
	if erro != nil {
		fmt.Println("erro ao converter")
	}

	err := checkmail.ValidateFormat(email)
	if err != nil {
		fmt.Println(err)
		c.JSON(http.StatusNotAcceptable, gin.H{
			"message": "Email invalido",
		})

	} else if len(cpf) != 9 {
		c.JSON(http.StatusNotAcceptable, gin.H{
			"message": "CPF invalido",
		})

	}

	var usuario banco.Usuario
	if erro := banco.DBClient.Where("ID LIKE ?", id).Find(&usuario); erro.Error != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"mensagem": "usuario nao encontrado",
		})
	}

	if erro := banco.DBClient.Model(&usuario).Updates(&banco.Usuario{Nome: nome, CPF: int64(cpff), DatadeNascimento: datadeNascimento, Telefone: int64(tel),
		Email: email, Rua: rua, Bairro: bairro, Complemento: comp, Cidade: int64(cidID)}); erro.Error != nil {
		c.JSON(http.StatusNonAuthoritativeInfo, gin.H{
			"message": "Erro ao atualizar usuário",
		})
	}

	var usuariob banco.UsuarioBusca
	if erro := banco.DBClient.Where("ID LIKE ?", id).Find(&usuariob); erro.Error != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"mensagem": "usuario nao encontrado",
		})
	}
	var cidade banco.Cidade
	banco.DBClient.First(&cidade, cid)
	banco.DBClient.Raw("SELECT Nome, Uf FROM cidades WHERE Id = ?", cidID).Scan(&cidade)

	fmt.Println("endpoint usuario")

	if erro := banco.DBClient.Model(&usuariob).Updates(&banco.UsuarioBusca{Nome: nome, CPF: int64(cpff), DatadeNascimento: datadeNascimento, Telefone: int64(tel),
		Email: email, Rua: rua, Bairro: bairro, Complemento: comp, Cidade: cidade.Nome}); erro.Error != nil {
		c.JSON(http.StatusNonAuthoritativeInfo, gin.H{
			"message": "Erro ao atualizar usuário",
		})
	}
	fmt.Println("endpoint usuariobusca")
	c.JSON(http.StatusOK, gin.H{
		"mensagem": "usuario atualizado",
	})

}

//DeleteUser deleta usuario pela id
func DeleteUser(c *gin.Context) {
	id := c.Query("id")

	var usuario banco.Usuario
	if erro := banco.DBClient.Where("id = ?", id).Delete(&usuario); erro.Error != nil {
		c.JSON(http.StatusNotModified, gin.H{
			"message": "Erro ao deletar usuário",
		})
	}

	var usuariob banco.UsuarioBusca
	if erro := banco.DBClient.Where("id = ?", id).Delete(&usuariob); erro.Error != nil {
		c.JSON(http.StatusNotModified, gin.H{
			"message": "Erro ao deletar usuário",
		})
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "Deletado com sucesso",
	})
}
