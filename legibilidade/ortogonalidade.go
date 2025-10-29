package main

import (
	"fmt"
	"hash/fnv"
	"strconv"
)

type Usuario struct {
	Nome  string
	Idade int
}

func NewUsuario(nome string, idade int) Usuario {
	return Usuario{Nome: nome, Idade: idade}
}

func (u *Usuario) GetNome() string {
	return u.Nome
}

func (u *Usuario) GetIdade() int {
	return u.Idade
}

func (u *Usuario) SetNome(nome string) {
	u.Nome = nome
}

func (u *Usuario) SetIdade(idade int) {
	u.Idade = idade
}

func (u *Usuario) ComemorarAniversario() {
	u.Idade++
}

func (u *Usuario) String() string {
	return fmt.Sprintf("nome: %s, idade: %d", u.Nome, u.Idade)
}

func criarUsuario(nome string, idade int) Usuario {
	return NewUsuario(nome, idade)
}

func criarId(usuario Usuario) uint32 {
	h := fnv.New32a()
	h.Write([]byte(usuario.GetNome() + "." + strconv.Itoa(usuario.GetIdade())))
	return h.Sum32()
}

func main() {
	// map -> Dicionario
	// TAD -> TIPO ABSTRATO DE DADOS
	usuarios := make(map[uint32]Usuario)

	mario := criarUsuario("mario da silva", 28)
	id1 := criarId(mario)
	usuarios[id1] = mario

	maria := criarUsuario("maria da silva", 32)
	id2 := criarId(maria)
	usuarios[id2] = maria

	jose := criarUsuario("josé da silva", 26)
	id3 := criarId(jose)
	usuarios[id3] = jose

	fmt.Println(usuarios)

	usuario := usuarios[id2]
	fmt.Println(usuario)

	delete(usuarios, id1)
	fmt.Println(usuarios)
}

