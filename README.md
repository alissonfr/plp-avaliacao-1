# Trabalho de Paradigmas de Linguagem de Programação - Golang

Este trabalho apresenta implementações em **Golang** (Go) para demonstrar os critérios de avaliação de linguagens de programação: **LEGIBILIDADE** e **FACILIDADE DE ESCRITA**, com foco nas características de **SIMPLICIDADE** e **ORTOGONALIDADE**.

## Comparações Implementadas

### 1. FACILIDADE DE ESCRITA

#### Simplicidade
**Exemplo:** Lista de compras com operações de filtro e agregação
- Java: Stream API com lambda expressions
- JavaScript: Arrow functions e métodos funcionais de array
- Python: Funções map, filter, reduce
- **Go**: Range loops e funções simples (mais verboso, mas mais explícito)

#### Ortogonalidade
**Exemplo:** Transformação de listas com filter e map
- Java: Stream().filter().map()
- JavaScript: array.filter().map()
- Python: List comprehension
- **Go**: Loop for com condicionais (simples e direto, sem abstrações)

### 2. LEGIBILIDADE

#### Simplicidade
**Exemplo:** Classe Usuario com getters, setters e métodos
- Java: Classe tradicional com getters/setters explícitos
- JavaScript: Classes ES6 com métodos
- Python: Classes com @property
- **Go**: Structs com métodos (receivers), sem encapsulamento tradicional

#### Ortogonalidade
**Exemplo:** Map de usuários com operações CRUD
- Java: HashMap<K, V>
- JavaScript: Objeto com operações nativas
- Python: Dictionary
- **Go**: map[T]U nativo, mais simples para operações básicas

#### Tipagem e Sintaxe
**Exemplo:** Operações matemáticas com tipagem
- Java: Tipagem forte, estática e explícita
- JavaScript: Tipagem fraca e dinâmica
- Python: Tipagem forte e dinâmica
- **Go**: Tipagem forte, estática e inferida (sem declarações desnecessárias)

## Características da Linguagem Go

### Simplicidade em Go:
- Sintaxe minimalista e clara
- Menos abstrações que Java/JavaScript
- Mais explícito que Python para operações complexas
- Requer mais código para algumas operações

### Ortogonalidade em Go:
- Conceitos se combinam de forma natural
- Structs + Methods + Interfaces permitem OOP quando necessário
- Menos "açúcar sintático" que Python/JavaScript
- Operações básicas são diretas e previsíveis

### Tipagem em Go:
- Forte: Tipos não são convertidos implicitamente
- Estática: Verificação em tempo de compilação
- Inferida: Não precisa declarar tipos quando óbvios
- Combina segurança de tipo com simplicidade de uso

