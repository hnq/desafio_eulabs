API de Gerenciamento de Produtos em GoLang
Este é o backend da API de gerenciamento de produtos, implementada em GoLang utilizando o framework Echo e MySQL para persistência de dados.

Configuração do Ambiente
1. Instalação de Dependências
Certifique-se de ter GoLang (v1.16 ou superior) instalado. Em seguida, instale as dependências do projeto:
go mod tidy

2. Configuração do Banco de Dados
Configure um banco de dados MySQL.
Edite o arquivo config/database.go com suas credenciais do MySQL.

3. Execução da Aplicação
Para iniciar o servidor da API:
go run main.go
O servidor estará acessível em http://localhost:8080.

Migrações de Banco de Dados
1. Criação de Migrações
Para criar as tabelas iniciais do banco de dados:

# Certifique-se de ter instalado a ferramenta goose
go get -u github.com/pressly/goose/cmd/goose

# Criação da migração inicial
goose -dir db/migrations mysql "user:password@tcp(127.0.0.1:3306)/productdb" create create_initial_tables sql

2. Aplicação de Migrações
Para aplicar as migrações ao banco de dados:
goose -dir db/migrations mysql "user:password@tcp(127.0.0.1:3306)/productdb" up

Testes
1. Execução dos Testes
Para executar os testes unitários:
go test ./...

2. Cobertura de Testes
Para verificar a cobertura de testes:
go test -coverprofile=coverage.out ./...
go tool cover -html=coverage.out

Documentação da API
Após iniciar o servidor, a documentação da API estará disponível em:
http://localhost:8080/swagger/index.html
Utilize esta interface interativa para explorar e testar os endpoints da API.

Configuração do Arquivo .env
Criação do Arquivo .env

Crie um arquivo .env na raiz do seu projeto. Este arquivo não deve ser versionado (adicionar ao .gitignore).

Definição das Variáveis de Ambiente

No arquivo .env, defina as variáveis de ambiente necessárias para configurar o seu projeto. Aqui está um exemplo básico com as variáveis usadas no projeto de exemplo:

# Configurações do Banco de Dados MySQL
DB_USER=seu_usuario_mysql
DB_PASSWORD=sua_senha_mysql
DB_HOST=127.0.0.1
DB_PORT=3306
DB_NAME=productdb

# Configurações do Servidor
SERVER_PORT=8080