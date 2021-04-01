package estruturasdedados

import "fmt"

type Agenda struct {
	contatos []*Contato
	NomeProprietario string
}

func NewAgenda(nomeProprietario string) *Agenda {
	novaAgenda := &Agenda{NomeProprietario: nomeProprietario}
	return novaAgenda
}

//metodo
func (t *Agenda) SalvarContato(novoContato *Contato) {
	t.contatos = append(t.contatos, novoContato)
}

func (v Agenda) ExibirLista() {
	fmt.Println("A SUA LISTA DE CONTATOS DE: ",v.NomeProprietario)
	for _, item := range v.contatos {
		fmt.Println("------------------")
		item.VerContato()
	}
	fmt.Println()
}
