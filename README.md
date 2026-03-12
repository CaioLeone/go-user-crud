# Conheça o projeto
Este projeto é um desafio prático para iniciantes em Go, focado na construção de uma API REST. O objetivo é desenvolver um sistema CRUD (Criar, Ler, Atualizar, Deletar) de usuários com armazenamento em memória, permitindo praticar os conceitos fundamentais de HTTP, como métodos, códigos de status e manipulação de JSON na linguagem Go.

## Sobre o Projeto
O objetivo deste projeto é construir uma API RESTful em Go para gerenciar usuários. Você implementará as operações básicas de CRUD (Criar, Ler, Atualizar e Deletar) utilizando um "banco de dados" em memória, o que permitirá focar nos conceitos de HTTP e na estrutura de uma aplicação web em Go.

## Estrutura do Usuário (Schema)
```
{
  "id": "",                     // UUID, obrigatório
  "first_name": "Jane Doe",     // String, obrigatória (mín. 2, máx. 20 caracteres)
  "last_name": "Jane Doe",      // String, obrigatória (mín. 2, máx. 20 caracteres)
  "biography": "Tendo diversão" // String, obrigatória (mín. 20, máx. 450 caracteres)
}
```

## "Banco de Dados" em Memória
Como o objetivo desse projeto é praticar os conceitos basicos de Go e também HTTP, não abordamos ainda sobre persistência. Devido a isso, iremos simular um banco de dados onde todo o conteúdo enviado por nossos usuários será salvo em um Hash Map, a onde o id é a chave de cada registro.

### Exemplo da estrutura
```
type id uuid.UUID

type user struct {
	FirstName string
	LastName  string
	biography string
}

type application struct {
	data map[id]user
}
```
Você deve criar um pacote que implemente as seguintes funções para manipular este mapa:

* FindAll(): Retorna a lista completa de usuários.
* FindById(id): Retorna o usuário correspondente ao id ou nil se não existir.
* Insert(newUser): Adiciona um novo usuário e retorna o usuário recém-criado com seu id.
* Update(id, userUpdates): Atualiza um usuário existente e retorna a versão atualizada.
* Delete(id): Remove um usuário e retorna o usuário que foi deletado.

## Endpoints da API e Especificações

Sua API deverá implementar os seguintes endpoints, cada um com sua lógica de negócio específica:

Obs.: Todas as respostas de erro devem ser retornadas no formato JSON, contendo uma mensagem descritiva do erro. Ex.:

```
{
  "error": "User not found"
}
```
### ```POST /api/users``` - Criar um novo usuário
Utiliza as informações enviadas no corpo da requisição para criar um novo usuário.

* Sucesso: Se o corpo da requisição for válido, salve o usuário, retorne o status 201 e o objeto do usuário recém-criado (incluindo o id).
* Erro (Bad Request): Se faltar first_name, last_name ou biography, retorne o status 400.
* Erro (Server Error): Em caso de falha ao salvar, retorne o status 500.

### ```GET /api/users``` - Listar todos os usuários
Retorna uma lista com todos os usuários cadastrados.

* Sucesso: Retorne o status 200 e a lista de todos os usuários.
* Erro (Server Error): Se ocorrer um erro ao buscar os dados, retorne o status 500.

### ```GET /api/users/:id``` - Buscar um usuário específico
Retorna o objeto do usuário com o id especificado na URL.

* Sucesso: Se o usuário for encontrado, retorne o status 200 e o objeto do usuário.
* Erro (Not Found): Se o id não existir, retorne o status 404.
* Erro (Server Error): Em caso de falha na busca, retorne o status 500.

### ```PUT /api/users/:id``` - Atualizar um usuário
Atualiza o usuário com o id especificado usando os dados enviados no corpo da requisição.

* Sucesso: Se o usuário for encontrado e os dados forem válidos, atualize-o, retorne o status 200 e o objeto do usuário atualizado.
* Erro (Not Found): Se o id não existir, retorne o status 404.
* Erro (Bad Request): Se faltar first_name, last_name ou biography no corpo da requisição, retorne o status 400.
* Erro (Server Error): Em caso de falha na atualização, retorne o status 500.

### ```DELETE /api/users/:id``` - Deletar um usuário
Remove o usuário com o id especificado na URL.

* Sucesso: Se o usuário for encontrado, remova-o e retorne o status 200 com o objeto do usuário que foi deletado.
* Erro (Not Found): Se o id não existir, retorne o status 404.
* Erro (Server Error): Em caso de falha na remoção, retorne o status 500.

## Importante

Teste seu trabalho manualmente usando Postman, Insomnia ou outra ferramentas de Cliente HTTP