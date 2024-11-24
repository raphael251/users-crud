# users-crud

## Running the project

### Running Outside Docker Compose

First you will need the environment variables. So copy the .env.example file and rename to .env. Fill the variables with the values you want (e.g. the database address, user and password).

Now run the command `go build cmd/main.go`. After the build, just run the binary with `./main`.

### Running With Docker Compose

Using docker compose you just need to run `docker compose up -d` to run the database and the API in background, and then start using the API.
