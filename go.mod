module github.com/trapajim/rest

go 1.12

// +heroku goVersion go1.12
require (
	github.com/auth0-community/go-auth0 v1.0.0
	github.com/denisenkom/go-mssqldb v0.0.0-20190204142019-df6d76eb9289 // indirect
	github.com/go-sql-driver/mysql v1.4.1 // indirect
	github.com/gofrs/uuid v3.2.0+incompatible // indirect
	github.com/gorilla/mux v1.7.0
	github.com/gorilla/sessions v1.1.3
	github.com/lib/pq v1.0.0
	github.com/mattn/go-runewidth v0.0.4 // indirect
	github.com/mattn/go-sqlite3 v1.10.0 // indirect
	github.com/mitchellh/cli v1.0.0 // indirect
	github.com/olekukonko/tablewriter v0.0.1 // indirect
	github.com/pkg/errors v0.8.1
	github.com/rs/cors v1.6.0
	github.com/rubenv/sql-migrate v0.0.0-20190212093014-1007f53448d7
	github.com/spf13/viper v1.3.1
	github.com/stretchr/testify v1.3.0 // indirect
	github.com/volatiletech/inflect v0.0.0-20170731032912-e7201282ae8d // indirect
	github.com/volatiletech/null v8.0.0+incompatible
	github.com/volatiletech/sqlboiler v3.2.0+incompatible
	golang.org/x/oauth2 v0.0.0-20190226205417-e64efc72b421
	gopkg.in/gorp.v1 v1.7.2 // indirect
	gopkg.in/square/go-jose.v2 v2.3.0
)

replace github.com/trapajim/rest/api => ./api

replace github.com/trapajim/rest/handler => ./handler

replace github.com/trapajim/rest/models => ./models

replace github.com/trapajim/rest/router => ./router

replace github.com/trapajim/rest/service => ./service
