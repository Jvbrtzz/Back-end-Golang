# Back-end-Golang
 
#  **Gerenciador de Cards e Usuários - Backend em Go**

Este projeto é uma API REST desenvolvida em **Go** para o Gerenciamento de QA, sendo uma 'refatoração' para linguagem **Go**

---

##  **Tecnologias Utilizadas**

- **Go (Golang)** 
- **Gorilla Mux** - Router HTTP
- **GORM** - ORM para Go
- **MySQL** - Banco de dados relacional
- **bcrypt** - Criptografia de senhas
- **dotenv** - Gerenciamento de variáveis de ambiente

---

## 📦 **Instalação**

```bash
# 1️⃣ Clone o repositório
$ git clone https://github.com/Jvbrtzz/Back-end-golang.git
$ cd Back-end-golang

# 2️⃣ Configure o arquivo .env
$ touch .env
```

Exemplo de `.env`:

```env
USER=seu_usuario
SENHA=sua_senha
DB=nome_do_banco
HOST=localhost:3306
```

```bash
# 3️⃣ Instale as dependências
$ go mod tidy

# 4️⃣ Inicie o servidor
$ go run main.go
```

---

## 🗂️ **Estrutura do Projeto**

```
📦 Back-end-golang
├── controllers        # Lógica dos endpoints
├── database           # Conexão com o banco de dados
├── models             # Estruturas de dados (Users, Cards, Comments)
├── routes             # Definição das rotas
├── main.go            # Ponto de entrada da aplicação
└── .env               # Variáveis de ambiente
```

---

## 📋 **Endpoints Disponíveis** (**EM CONSTRUÇÂO**)

### 🔐 **Usuários**

- `GET /users` → Lista todos os usuários
- `GET /users/{id}` → Retorna um usuário específico
- `POST /users` → Registra um novo usuário (senha criptografada)

### 🗂️ **Cards**

- `GET /cards/user/{id}` → Retorna todos os cards de um usuário
- `POST /cards` → Cria um novo card

### 💬 **Comentários**

- `GET /comments/card/{id}` → Retorna todos os comentários de um card

---

## 🔐 **Segurança**

- As **senhas dos usuários** são **criptografadas** usando o `bcrypt` antes de serem salvas no banco de dados.

---

## 💡 **Contribuição**

1. **Fork** este repositório
2. Crie uma nova branch:
   ```bash
   git checkout -b feature-nova
   ```
3. Faça suas alterações e **commit**:
   ```bash
   git commit -m "Adiciona nova feature"
   ```
4. Envie para o seu repositório:
   ```bash
   git push origin feature-nova
   ```
5. Abra um **Pull Request** 🚀

---

##

