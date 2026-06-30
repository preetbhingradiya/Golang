package intermediate

import (
	"log"
	"os"
)

func EnvVariable() {

	err := os.Setenv("MY_ENV_VAR", "Hello, World!")
	if err != nil {
		panic(err)
	}

	value := os.Getenv("MY_ENV_VAR")
	infoLogger.Println("The value of MY_ENV_VAR is:", value)

	jwt := os.Getenv("JWT_SECRET")
	if jwt == "" {
		errorLogger.Println("JWT_SECRET is not set")
	}

	infoLogger.Println("Environment variable MY_ENV_VAR set successfully")
	log.SetPrefix("INFO : ")
	log.Println("This is an info message")

}

// Logger staructure
var (
	infoLogger  = log.New(os.Stdout, "INFO : ", log.Ldate|log.Ltime|log.Lshortfile)
	errorLogger = log.New(os.Stderr, "ERROR : ", log.Ldate|log.Ltime|log.Lshortfile)
	warnLogger  = log.New(os.Stdout, "WARN : ", log.Ldate|log.Ltime|log.Lshortfile)
)
