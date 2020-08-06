module github.com/fenger/golibdemos

go 1.14

require (
	github.com/FZambia/sentinel v1.1.0
	github.com/fenger/gomoon v0.0.1
	github.com/go-sql-driver/mysql v1.5.0
	github.com/gomodule/redigo v1.8.2
	github.com/google/uuid v1.1.1
	github.com/spf13/cobra v1.0.0
)

replace github.com/fenger/gomoon => ../gomoon
