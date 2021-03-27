package estruturasdedados

import "fmt"

type ListaDeContatos struct {
	contatos []*Contato
}

func NewListaDeContatos() *ListaDeContatos {
	return &ListaDeContatos{}
}

//metodo
func (t *ListaDeContatos) SalvarContato(novoContato *Contato) {
	t.contatos = append(t.contatos, novoContato)
}

func (v ListaDeContatos) ExibirLista() {
	fmt.Println("A SUA LISTA DE CONTATOS: ")
	for _, item := range v.contatos {
		fmt.Println("------------------")
		item.VerContato()
	}
	fmt.Println()
}
