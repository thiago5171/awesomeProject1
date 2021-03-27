package estruturasdedados

type ListaDeContatos struct {
	*contatos []Contato
}

func NewListaDeContatos() ListaDeContatos {
	return &ListaDeContatos{}
}

//metodo
func (t *ListaDeContatos) SalvarContato(novoContato Contato) {
	t.contatos = append(t.contatos, novoContato)
}
func (v ListaDeContatos) ExibirLista() {
	for _, item := range v.contatos {
		item.VerContato()
	}

}
