package grpchttpamqp

import (
	"fmt"

	"github.com/cpapidas/pegasus/netamqp"
	"github.com/cpapidas/pegasus/netgrpc"
	"github.com/cpapidas/pegasus/nethttp"
	"github.com/cpapidas/pegasus/peg"
)

func handler(channel *peg.Channel) {

	// Get the payload
	payload := channel.Receive()

	// Just create the reply message
	replyMessage := string(payload.Body) + " world :)"

	// Unmarshal the options in order to get the token from headers
	options := peg.NewOptions().Unmarshal(payload.Options)

	// Get the header (HTTP-GRPC-AMQP)
	token := options.GetHeader("Token")

	// Get url param (HTTP-GRPC-AMQP)
	order := options.GetParam("order")

	replyMessage += " token:" + token + " order:" + order

	fmt.Println("------------------------ RESPONSE ------------------------")
	fmt.Println("Reply Message:", replyMessage)
	fmt.Println("Receive Message:", replyMessage)
	fmt.Println("Receive options:", options)
	fmt.Println("----------------------------------------------------------")

	// AMQP will ignore this because we don't have any response
	// Send to client the response
	channel.Send(peg.BuildPayload([]byte(replyMessage), nil))
}

// Server initialize the server
func Server() {

	rabbitMQAddress := "amqp://guest:guest@localhost:5672/"

	// Create the servers objects.
	grpcServer := netgrpc.NewServer(nil)
	httpServer := nethttp.NewServer()
	amqpServer := netamqp.NewServer()

	// We have to start the amqp server before listeners in order to get the connection object
	amqpServer.Serve(rabbitMQAddress)

	// Create the listeners
	grpcServer.Listen(netgrpc.SetConf("/sample/put"), handler, nil)
	httpServer.Listen(nethttp.SetConf("/sample/put", nethttp.Put), handler, nil)
	amqpServer.Listen(netamqp.SetConf("/sample/put"), handler, nil)

	// We have to keep the main goroutine up so we have to create something like while(true) but more elegant
	stop := make(chan bool)

	// Start the servers
	grpcServer.Serve("localhost:9091")
	httpServer.Serve("localhost:9092")

	// Print a cool message
	fmt.Println("\nThe servers GRPC-HTTP is up and running,\n" +
		"now run [$ ./examples grpchttp client] command :) \n")

	// wait here forever
	<-stop
}
