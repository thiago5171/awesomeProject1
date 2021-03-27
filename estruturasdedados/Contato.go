package estruturasdedados

import "fmt"
//estrutura/classes

type Contato struct {
	Nome     string
	Telefone string
	Email    string
}

// metodo
func (c Contato) VerContato() {

	fmt.Println("Nome->", c.Nome)
	fmt.Println("Telefone->", c.Telefone)
	fmt.Println("Email->", c.Email)
}

// metodo
func (this Contato) Equal(contat Contato) bool {
	if this.Email == contat.Email {
		return true
	}

	return false

	//esse metodo tambem pode ser usado como substituto do if ao retornar uma comparação
	//return this.Email == contat.Email
}

// metodo
func NewContato(nome, telefone, email string)  Contato {
	return Contato{nome, telefone, email}

}
