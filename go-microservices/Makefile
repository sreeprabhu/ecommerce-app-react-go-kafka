# Check for the operating system
ifeq ($(OS),Windows_NT)
    SET_ENV=set APP_ENV="dev"
else
    SET_ENV=APP_ENV="dev"
endif

# Server Command (using nodemon)
server:
	$(SET_ENV) & nodemon --watch './**/*.go' --exec "go run main.go"
