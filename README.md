# go-context

reproducing error of Golang Postgres context

## Setup

1. Clone the repo
2. Make sure docker is installed
3. Run `docker-compose up`
4. After all containers are ready, try to run `curl localhost:8080/user`. You will see the error `context cancelled` in logs. It is supposed to be print the user name and email in the logs.