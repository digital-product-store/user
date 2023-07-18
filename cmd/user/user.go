//go:generate oapi-codegen -package gen -generate types,server,spec -o ../../pkg/server/gen/specs.gen.go ../../specs/openapi3.yaml

package main

import (
	"io/ioutil"
	"os"
	"userservice/pkg/config"
	"userservice/pkg/server"
	"userservice/pkg/storage"
	"userservice/pkg/token"

	"go.uber.org/zap"
)

func main() {
	logger, _ := zap.NewProduction()
	defer logger.Sync()

	conf, err := config.LoadConfig()
	if err != nil {
		logger.Error("error on during configuration", zap.Error(err))
		os.Exit(-1)
	}

	jwtGen, err := createJWTGenerator(conf)
	if err != nil {
		logger.Error("error during initialization of jwt generator", zap.Error(err))
		os.Exit(-1)
	}

	roDB := createROMemStorage()
	handler := server.NewHandler(logger, roDB, jwtGen)
	srvr := server.NewServer(&handler, conf)

	srvr.Listen()
}

func createROMemStorage() storage.ROMemStorage {
	users := []storage.User{
		{
			Uuid:     "cc1b9a28-302c-4097-88c7-be1dff561325",
			Email:    "user1@test.local",
			Username: "user1",
			Fullname: "User One",
			Password: "pass1",
			Roles:    []string{"user"},
		},
		{
			Uuid:     "4d8e4532-fbb5-4e77-a69d-95cde4ab3d44",
			Email:    "user2@test.local",
			Username: "user2",
			Fullname: "User Two",
			Password: "pass2",
			Roles:    []string{"user"},
		},
		{
			Uuid:     "0cda5e37-b995-48dc-8051-ec089ed0bd86",
			Email:    "admin@test.local",
			Username: "admin",
			Fullname: "Ad Min",
			Password: "admin",
			Roles:    []string{"user", "admin"},
		},
	}

	return storage.NewROMemStorage(users)
}

func createJWTGenerator(conf *config.Config) (*token.JWTGenerator, error) {
	privKey, err := ioutil.ReadFile(conf.JWTPrivateKey)
	if err != nil {
		return nil, err
	}

	jwtGen, err := token.NewJWTGenerator(privKey, conf.JWTIssuer, conf.JWTSubject, conf.JWTTTL)
	return jwtGen, err
}
