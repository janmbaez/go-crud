# How to Run
```bash
go mod init go-crud
go get github.com/githubnemo/CompileDaemon
got get github.com/joho/godotenv
go get -u gorm.io/gorm
go get -u gorm.io/driver/postgres
CompileDaemon -commnad="./go-crud" 
```

## Database

> The Database use in this repository is ElephantSQL easy to setup

## Update the .env File using your own ElephantSQL credentials

```bash
DB_SCHEMA="host= user= password= dbname= port=5432 sslmode=disable"
```