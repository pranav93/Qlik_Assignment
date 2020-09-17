# Export env file

```bash
set -o allexport                                                 ─╯
source config.env
set +o allexport
```

# Build docker container
```bash
docker build -t qlik_backend .
```

# Tag a docker container
```bash
docker image tag qlik_backend:latest pranav93/qlik_backend:v1
```

docker exec -it qlik_assignment_postgres_1 /bin/bash