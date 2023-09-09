# Running CompiledDaemon
CompileDaemon --build="go build main.go" --command="./main"

# Running migrations, sun script at root project
go run ./migrations/migrate.go 
