package main

import (
	productPb "GolangMongoDBGRPCSimpleCRUD/Controller"
	"GolangMongoDBGRPCSimpleCRUD/Mongo"
	"GolangMongoDBGRPCSimpleCRUD/Service"
	"flag"
	"github.com/joho/godotenv"
	"go.mongodb.org/mongo-driver/mongo"
	"google.golang.org/grpc"
	"log"
	"net"
)

var environment = flag.String("env", "", "Environment")

func init() {
	flag.Parse()
	if *environment == "" {
		log.Fatalf("Environment is required")
	}
}

func main() {
	env, errReadingEnv := godotenv.Read(".env." + *environment)

	netListen, errRunninggRPC := net.Listen("tcp", env["APP_PORT"])

	defer func(netListen net.Listener) {
		err := netListen.Close()
		if err != nil {
			log.Fatalf("Failed to close net.Listen: %v", err)
		}
	}(netListen)

	mongoInstance := Mongo.NewMongoDBConfig(env)
	mongoConnection, errConnectingMongo := mongoInstance.OpenMongoDatabaseConnection()

	defer func(mongo Mongo.MongoDBConfig, client *mongo.Client) {
		errClosingDatabase := mongo.CloseMongoDatabaseConnection(client)
		if errClosingDatabase != nil {
			log.Fatalf("Failed to close database: %v", errClosingDatabase)
		}
	}(mongoInstance, mongoConnection)

	grpcServer := grpc.NewServer()

	productService := Service.ProductService{DB: mongoConnection}
	productPb.RegisterProductControllerServer(grpcServer, &productService)

	if errReadingEnv != nil {
		log.Fatalf("Failed to read env: %v", errReadingEnv)
	}
	if errRunninggRPC != nil {
		log.Fatalf("Failed to listen on port %s: %v", env["APP_PORT"], errRunninggRPC)
	}
	if errConnectingMongo != nil {
		log.Fatalf("Failed to connect to mongo: %v", errConnectingMongo)
	}

	log.Println("gRPC server running on port", netListen.Addr())

	if err := grpcServer.Serve(netListen); err != nil {
		log.Fatalf("Failed to serve gRPC server over port %s: %v", env["APP_PORT"], err)
	}
}
