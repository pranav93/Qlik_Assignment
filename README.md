# Qlik_Assignment

Assignment solution for SRE

## How to run locally?

Make sure that you have docker and compose installed on your dev machine.
Run the following commands in terminal.

```bash
git clone https://github.com/pranav93/Qlik_Assignment.git
cd Qlik_Assignment
```

This setup requires the `config.env` file in `Qlik_Assignment` directory.
Create a `config.env`, like this

```bash
touch config.env
```
Add something like the following
```env
DB_USERNAME=great_user
DB_PASSWORD=super_secret_password
DB_HOST=postgres
DB_PORT=5432
DB_NAME=qlik_database
DB_SSLMODE=disable
GIN_MODE=debug
ALLOWED_ORIGIN="*"

POSTGRES_USER=great_user
POSTGRES_PASSWORD=super_secret_password
POSTGRES_DB=qlik_database
```

Now run the backend

```
docker-compose up
```

Once you curl the following, you should get a response.
```bash
curl localhost:8080/api/
```

Now, once the backend is running as expected, refer [https://github.com/pranav93/Qlik_Frontend] to setup the frontend

## What do these env fields even mean?

| Field             |                                Value                                |
| ----------------- | :-----------------------------------------------------------------: |
| DB_USERNAME       |                         Database user name                          |
| DB_PASSWORD       |                          Database password                          |
| DB_HOST           |                         Database host name                          |
| DB_PORT           |                        Database port number                         |
| DB_NAME           |                            Database name                            |
| DB_SSLMODE        |                  disable sslmode for db connection                  |
| GIN_MODE          |              `debug` (for dev) or `release` (for prod)              |
| ALLOWED_ORIGIN    | For dev setup cors is used, in that case specify the allowed origin |
| POSTGRES_USER     |                        Postgres db username                         |
| POSTGRES_PASSWORD |                        Postgres db password                         |
| POSTGRES_DB       |                          Postgres db name                           |

For prod deployment I have created an image here -> [https://hub.docker.com/repository/docker/pranav93/qlik_backend]
For K8s deployment refer to -> [https://github.com/pranav93/Qlik_Kubernetes]
