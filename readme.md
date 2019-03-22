#rest api for wh40kstats

##dependencies

| package     | link                                        |
| ----------- | ------------------------------------------- |
| sqlboiler   | (https://github.com/volatiletech/sqlboiler) |
| sql-migrate | (https://github.com/rubenv/sql-migrate)     |
| Mux         | (https://github.com/gorilla/mux)            |
| Mux         | (https://github.com/gorilla/mux)            |
| viper       | (https://github.com/spf13/viper)            |
| auth0       | (https://auth0.com/)                        |

## migrations

first create a new migration file with

```bash
#replace name with your migration name
sql-migrate new -config ./api/config/config.yaml NAME
```

after that you can run the migration

```bash
#run migration
sql-migrate up -config ./api/config/config.yaml
#regenerate models
sqlboiler psql --wipe
```
