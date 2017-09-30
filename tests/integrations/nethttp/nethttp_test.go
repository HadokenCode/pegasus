package nethttp_test

import (
	"bitbucket.org/code_horse/pegasus/network"
	"bitbucket.org/code_horse/pegasus/network/nethttp"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("Nethttp", func() {

	var handlerGet = func(channel *network.Channel) {
		// Receive the payload
		receive := channel.Receive()

		// Unmarshal options, change them and send them back
		options := network.NewOptions().Unmarshal(receive.Options)

		replyOptions := network.NewOptions()

		replyOptions.SetHeader("Custom", options.GetHeader("Custom")+" response")

		// Create the new payload
		payload := network.BuildPayload([]byte(options.GetParam("foo")+" response"), replyOptions.Marshal())

		// Send it back
		channel.Send(payload)
	}

	var handlerPost = func(channel *network.Channel) {
		// Receive the payload
		receive := channel.Receive()

		// Unmarshal options, change them and send them back
		options := network.NewOptions().Unmarshal(receive.Options)

		replyOptions := network.NewOptions()

		replyOptions.SetHeader("Custom", options.GetHeader("Custom")+" response")
		replyOptions.SetHeader("name", options.GetParam("name")+" response")

		responseBody := string(receive.Body) + " response"

		// Create the new payload
		payload := network.BuildPayload([]byte(responseBody), replyOptions.Marshal())

		// Send it back
		channel.Send(payload)
	}

	var handlerPut = func(channel *network.Channel) {
		// Receive the payload
		receive := channel.Receive()

		// Unmarshal options, change them and send them back
		options := network.NewOptions().Unmarshal(receive.Options)

		replyOptions := network.NewOptions()

		replyOptions.SetHeader("Custom", options.GetHeader("Custom")+" response")
		replyOptions.SetHeader("name", options.GetParam("name")+" response")
		replyOptions.SetHeader("id", options.GetParam("id")+" response")

		responseBody := string(receive.Body) + " response"

		// Create the new payload
		payload := network.BuildPayload([]byte(responseBody), replyOptions.Marshal())

		// Send it back
		channel.Send(payload)
	}

	var handlerDelete = func(channel *network.Channel) {
		// Receive the payload
		receive := channel.Receive()

		// Unmarshal options, change them and send them back
		options := network.NewOptions().Unmarshal(receive.Options)

		replyOptions := network.NewOptions()

		replyOptions.SetHeader("Custom", options.GetHeader("Custom")+" response")
		replyOptions.SetHeader("name", options.GetParam("name")+" response")
		replyOptions.SetHeader("id", options.GetParam("id")+" response")

		// Create the new payload
		payload := network.BuildPayload([]byte(string(receive.Body)+" response"), replyOptions.Marshal())

		// Send it back
		channel.Send(payload)
	}

	var middleware = func(handler network.Handler, channel *network.Channel) {

		// Receive the payload
		receive := channel.Receive()

		// Unmarshal options, change them and send them back
		options := network.NewOptions().Unmarshal(receive.Options)

		options.SetHeader("Custom", options.GetHeader("Custom")+" middleware")

		// Create the new payload
		payload := network.BuildPayload(nil, options.Marshal())

		// Send it back
		channel.Send(payload)

		handler(channel)
	}

	server := nethttp.NewServer(nil)

	server.Listen(nethttp.SetPath("/http", nethttp.Get), handlerGet, nil)
	server.Listen(nethttp.SetPath("/http", nethttp.Post), handlerPost, nil)
	server.Listen(nethttp.SetPath("/http/{id}", nethttp.Put), handlerPut, nil)
	server.Listen(nethttp.SetPath("/http/{id}", nethttp.Delete), handlerDelete, nil)

	server.Listen(nethttp.SetPath("/http/middleware", nethttp.Get), handlerGet, middleware)

	server.Serve("localhost:7000")

	Describe("HTTP Server", func() {

		Context("Exchange message via HTTP", func() {

			It("Should not be nil", func() {
				Expect(server).ToNot(BeNil())
			})

		})

		Context("Send a GET request", func() {
			// Create a payload
			options := network.NewOptions()

			options.SetHeader("Custom", "header-value")

			payload := network.BuildPayload(nil, options.Marshal())

			// Send the payload
			response, err := nethttp.NewClient().
				Send(nethttp.SetPath("http://localhost:7000/http?foo=bar", nethttp.Get), payload)

			It("Should not throw an error", func() {
				Expect(err).To(BeNil())
			})

			It("The response should have the following values", func() {
				Expect(response.Body).To(Equal([]byte("bar response")))
				options := network.NewOptions().Unmarshal(response.Options)
				Expect(options.GetHeader("Custom")).To(Equal("header-value response"))
			})
		})

		Context("Send a POST request", func() {
			// Create a payload
			options := network.NewOptions()

			options.SetHeader("Custom", "header-value")

			payload := network.BuildPayload([]byte("foo"), options.Marshal())

			// Send the payload
			response, err := nethttp.NewClient().
				Send(nethttp.SetPath("http://localhost:7000/http?name=christos", nethttp.Post), payload)

			replyOptions := network.NewOptions().Unmarshal(response.Options)

			It("Should not throw an error", func() {
				if err != nil {
					panic(err)
				}
				Expect(err).To(BeNil())
			})

			It("The response should have the following values", func() {
				Expect(response.Body).To(Equal([]byte("foo response")))
				Expect(replyOptions.GetHeader("Custom")).To(Equal("header-value response"))
			})

			It("Should returns the param name", func() {
				Expect(replyOptions.GetHeader("Name")).To(Equal("christos response"))
			})
		})

		Context("Send a PUT request", func() {
			// Create a payload
			options := network.NewOptions()

			options.SetHeader("Custom", "header-value")

			payload := network.BuildPayload([]byte("foo"), options.Marshal())

			// Send the payload
			response, err := nethttp.NewClient().
				Send(nethttp.SetPath("http://localhost:7000/http/44?name=christos", nethttp.Put), payload)

			replyOptions := network.NewOptions().Unmarshal(response.Options)

			It("Should not throw an error", func() {
				if err != nil {
					panic(err)
				}
				Expect(err).To(BeNil())
			})

			It("The response should have the following values", func() {
				Expect(response.Body).To(Equal([]byte("foo response")))
				Expect(replyOptions.GetHeader("Custom")).To(Equal("header-value response"))
			})

			It("Should returns the param name", func() {
				Expect(replyOptions.GetHeader("Name")).To(Equal("christos response"))
			})

			It("Should return the path param", func() {
				Expect(replyOptions.GetHeader("Id")).To(Equal("44 response"))
			})
		})

		Context("Send a DELETE request", func() {
			// Create a payload
			options := network.NewOptions()

			options.SetHeader("Custom", "header-value")

			payload := network.BuildPayload([]byte("foo"), options.Marshal())

			// Send the payload
			response, err := nethttp.NewClient().
				Send(nethttp.SetPath("http://localhost:7000/http/44?name=christos", nethttp.Delete), payload)

			replyOptions := network.NewOptions().Unmarshal(response.Options)

			It("Should not throw an error", func() {
				if err != nil {
					panic(err)
				}
				Expect(err).To(BeNil())
			})

			It("The response should have the following values", func() {
				Expect(response.Body).To(Equal([]byte("foo response")))
				Expect(replyOptions.GetHeader("Custom")).To(Equal("header-value response"))
			})

			It("Should returns the param name", func() {
				Expect(replyOptions.GetHeader("Name")).To(Equal("christos response"))
			})

			It("Should return the path param", func() {
				Expect(replyOptions.GetHeader("Id")).To(Equal("44 response"))
			})

		})

		Context("Send a GET middleware request", func() {
			// Create a payload
			options := network.NewOptions()

			options.SetHeader("Custom", "header-value")

			payload := network.BuildPayload(nil, options.Marshal())

			// Send the payload
			response, err := nethttp.NewClient().
				Send(nethttp.SetPath("http://localhost:7000/http/middleware?foo=bar", nethttp.Get), payload)

			It("Should not throw an error", func() {
				Expect(err).To(BeNil())
			})

			It("The response should have the following values", func() {
				Expect(response.Body).To(Equal([]byte("bar response")))
				options := network.NewOptions().Unmarshal(response.Options)
				Expect(options.GetHeader("Custom")).To(Equal("header-value middleware response"))
			})
		})

	})

})