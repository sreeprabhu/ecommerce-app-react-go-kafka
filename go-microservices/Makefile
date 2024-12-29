ifeq ($(OS),Windows_NT)
    $env:APP_ENV="dev";
else
    SET_ENV=APP_ENV=dev
endif

server:
	nodemon --watch './**/*.go' --signal SIGTERM --exec $(SET_ENV) 'go' run main.go