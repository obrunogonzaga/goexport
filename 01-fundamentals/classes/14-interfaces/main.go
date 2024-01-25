package main

import "fmt"

type Endereco struct {
	Logradouro string
	Numero     int
	Cidade     string
	Estado     string
}

type Pessoa interface {
	Desativar()
}

type Empresa struct {
	Nome  string
}

func (e Empresa) Desativar() {
	fmt.Printf("Empresa %s desativada com sucesso!\n", e.Nome)
}

type Cliente struct {
	Nome     string
	Idade    int
	Ativo    bool
	Address  Endereco // criando variavel do tipo Endereco
	Endereco          // composição de structs Endereco (como se fosse herança)
}

func (c Cliente) GetNome() string {
	return c.Nome
}

func (c Cliente) Desativar() {
	c.Ativo = false
	fmt.Printf("Cliente %s desativado com sucesso!\n", c.Nome)
}

func Desativacao(pessoa Pessoa) {
	pessoa.Desativar()
}

func main() {
	bruno := Cliente{
		Nome:  "Bruno",
		Idade: 37,
		Ativo: true,
	}
	bruno.Ativo = false

	bruno.Cidade = "Curitiba"
	bruno.Estado = "Paraná"
	bruno.Numero = 380
	bruno.Endereco.Logradouro = "Rua Tabajaras"

	bruno.Address.Cidade = "Sao Paulo"
	bruno.Address.Estado = "Sao Paulo"
	bruno.Address.Numero = 123
	bruno.Address.Logradouro = "Rua das flores"

	fmt.Printf("Nome: %s, Idade: %d, Ativo: %t\n", bruno.GetNome(), bruno.Idade, bruno.Ativo)
	fmt.Printf("Cidade: %s, Estado: %s, Número: %d, Logradouro: %s\n", bruno.Cidade, bruno.Estado, bruno.Numero, bruno.Endereco.Logradouro)
	fmt.Printf("Cidade: %s, Estado: %s, Número: %d, Logradouro: %s\n", bruno.Address.Cidade, bruno.Address.Estado, bruno.Address.Numero, bruno.Address.Logradouro)

	minhaEmpresa := Empresa{}

	Desativacao(bruno)
	Desativacao(minhaEmpresa)

}
