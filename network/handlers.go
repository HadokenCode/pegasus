package network

// Handler is a type of function which has a parameter a network.Channel. The handler was used at network.Client::Send &
// network.Server::Listen function in order to define a handler function which will process the request. Also handler
// is defined at network.Middleware.
//
// Channel is the only and main way to transport data from Handlers to receivers. It contains a payload struct field
// with two nested fields payload.Options and payload.Body.
//
// Channel exists only in network.Handler, network.Middleware and network.Client::Listen files, as parameter for those
// functions
//
// The payload.Options can be used in order to transport some options via options.Headers or some parameters via
// options.Params. The options.Body is similar with http body and is used in order to transfer raw content.
//
// An example for usage could be:
//	func handler(channel *network.Channel) {
//		payload := channel.Receive() // Receive the payload here
//		options := network.BuildOptions(payload.Options) // Convert the received options to options struct
// 		payload.Body = payload.Body + []byte(" sub-fix") // Change the body
//		chanel.send(network.BuildPayload(payload.Body, options.Marshal())) // return the payload
//	}
type Handler func(chanel *Channel)

// Middleware is a type of function which is executed before network.handler function. It has two parameters the
// network.Handler and the network.channel. It is used only at network.Server::Listen function. Usually the middleware
// type of function would be nil.
//
// Handler parameter is used as handler function which will handler a specif request.
//
// Channel parameter is the only and main way to transport data from Handlers to receivers. It contains a payload struct
// field with two nested fields payload.Options and payload.Body.
//
// The Middleware is responsible to call or not the handler function. Also can edit the network.Channel data that
// handler will get from channel parameter.
//
// An example for usage could be:
//	func middleware(handler Handler, chanel *Channel) {
//		payload := channel.Receive() // Receive the payload here
//		options := network.BuildOptions(payload.Options) // Convert the received options to options struct
// 		options.SetHeader("Auth-Token", "@#$fsf2hkj42@#@#") // Add an extra header
// 		payload.Body = payload.Body + []byte(" sub-fix") // Change the body
//		chanel.send(network.BuildPayload(payload.Body, options.Marshal())) // Add a new payload
// 		handler(chanel) // Call the handler with new channel
//	}
type Middleware func(handler Handler, channel *Channel)