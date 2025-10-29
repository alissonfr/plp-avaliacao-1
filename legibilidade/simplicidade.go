package main

import (
	"fmt"
	"strings"
)

type Usuario struct {
	Nome  string
	Idade int
}

func (u *Usuario) SetNome(nome string) {
	u.Nome = nome
}

func (u *Usuario) GetNome() string {
	return u.Nome
}

func (u *Usuario) SetIdade(idade int) {
	u.Idade = idade
}

func (u *Usuario) GetIdade() int {
	return u.Idade
}

func (u *Usuario) ComemorarAniversario() {
	u.Idade++
}

func (u *Usuario) String() string {
	return fmt.Sprintf("nome: %s, idade: %d", u.Nome, u.Idade)
}

func main() {
	usuario := Usuario{}

	usuario.SetNome("mario da silva")
	usuario.SetIdade(28)
	fmt.Println(usuario)

	usuario.ComemorarAniversario()
	fmt.Println(usuario)

	if strings.HasPrefix(usuario.GetNome(), "mario") {
		fmt.Println("é mario")
	} else {
		fmt.Println("não é mario")
	}
}

