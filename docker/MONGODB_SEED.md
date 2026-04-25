# MongoDB Seed Setup

Este projeto está configurado para inicializar automaticamente o MongoDB com dados de teste quando o container é iniciado.

## Como usar

### Iniciar MongoDB com dados de seed

```bash
# Na pasta raiz do projeto
docker-compose up -d
```

O MongoDB será iniciado e automaticamente carregará os seguintes dados:

- **genres.json** - Gêneros de filmes (9 registros)
- **movies.json** - Filmes com informações de IMDB
- **users.json** - Usuários de teste (incluindo admin)
- **rankings.json** - Sistema de classificação

### Banco de dados

- **Nome do BD**: `movie_streaming`
- **Host**: localhost
- **Porta**: 27017

### Acessar os dados

```bash
# Conectar ao MongoDB
mongosh mongodb://localhost:27017/movie_streaming

# Listar coleções
db.getCollectionNames()

# Ver documentos
db.genres.find()
db.movies.find()
db.users.find()
db.rankings.find()
```

### Usuário de teste

Email: `bobjones@hotmail.com`
Role: `ADMIN`

### Resetar dados

Para remover os dados de seed e começar do zero:

```bash
# Parar e remover o container
docker-compose down

# Remover o volume de dados
docker volume rm docker_mongodb_data

# Iniciar novamente
docker-compose up -d
```

## Arquivos de configuração

- `docker-compose.yml` - Configuração do MongoDB com volumes de seed
- `init-mongo.js` - Script de inicialização que carrega os dados JSON
- `seed/magic-stream-seed-data/` - Dados JSON para semear o banco

## Estrutura dos Arquivos de Seed

Todos os arquivos estão em formato JSON Array e são automaticamente importados para suas respectivas coleções:

- `genres.json` → coleção `genres`
- `movies.json` → coleção `movies`
- `users.json` → coleção `users`
- `rankings.json` → coleção `rankings`
