package main

import (
	"awesomeProject1/estruturasdedados"
)

func main() {
	thiago := &estruturasdedados.Contato{
		Nome:     "thiago",
		Telefone: "9999-9999",
		Email:    "tgazaroli@gmail.com",
	} // CTRL+SPAC ==> FILL ALL FIELDS

	// essa é uma maqneira  diferente de declarar as caracteristicas do  objeto inteiro
	//gabriel := estruturasdedados.Contato{" gabriel", "9999-8888", "gabriel@gmail.com"}

	gabriel := estruturasdedados.NewContato(" gabriel", "9999-8888", "gabriel@gmail.com")

	//	gabriel.VerContato()

	//thiago.VerContato()

	lista := estruturasdedados.NewAgenda("maria")
	lista.SalvarContato(thiago)
	lista.SalvarContato(gabriel)
	lista.ExibirLista()


	lista2 := estruturasdedados.NewAgenda("jose")
	lista2.SalvarContato(thiago)
	lista2.ExibirLista()
}
