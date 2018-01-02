package api

import (
	"log"
	"net/http"
	"os"

	gw "github.com/EwanValentine/shippy-user-api/proto/user"
	"github.com/grpc-ecosystem/grpc-gateway/runtime"
	"golang.org/x/net/context"
	"google.golang.org/grpc"
)

// Init a rest API
func main() {
	ctx := context.Background()
	ctx, cancel := context.WithCancel(ctx)
	defer cancel()

	// New server multiplexer
	mux := runtime.NewServeMux()
	opts := []grpc.DialOption{grpc.WithInsecure()}

	// Our gRPC host address
	conn := os.Getenv("SERVICE_ADDRESS")
	apiAddress := os.Getenv("API_ADDRESS")

	log.Printf("Connecting to gRPC server on: %s\n", conn)
	log.Printf("Starting API on: %s\n", apiAddress)

	// Register the handler to an endpoint
	err := gw.RegisterUserServiceHandlerFromEndpoint(ctx, mux, conn, opts)
	if err != nil {
		log.Fatal(err)
	}

	// Return a server instance
	http.ListenAndServe(apiAddress, mux)
}
