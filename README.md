# 1 Install dependence
go mod tidy
# 2 Run Application
go run main.go
# 3 Build Application
go build
# 4 Build Docker
docker build -t gojango .
# 5 Run Docker
docker run --name gojango -p 8000:8000 -v /config/settings.yml:/config/settings.yml -d gojango-server