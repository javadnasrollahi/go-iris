package socket

import (
	"log"
	"net/http"
	"os"

	"github.com/kataras/neffos"
	"github.com/kataras/neffos/gobwas"
)

var (
	upgrader = gobwas.DefaultUpgrader
	handler  = neffos.WithTimeout{
		ReadTimeout:  timeout,
		WriteTimeout: timeout,
		Namespaces: neffos.Namespaces{
			"default": neffos.Events{
				neffos.OnNamespaceConnected: func(c *neffos.NSConn, msg neffos.Message) error {
					log.Printf("[%s] connected to [%s].", c.Conn.ID(), msg.Namespace)

					if !c.Conn.IsClient() {
						c.Emit("chat", []byte("welcome to server's namespace"))
					}

					return nil
				},
				neffos.OnNamespaceDisconnect: func(c *neffos.NSConn, msg neffos.Message) error {
					log.Printf("[%s] disconnected from [%s].", c.Conn.ID(), msg.Namespace)

					if c.Conn.IsClient() {
						os.Exit(0)
					}

					return nil
				},
				neffos.OnRoomJoined: func(c *neffos.NSConn, msg neffos.Message) error {
					log.Printf("[%s] joined to room [%s].", c.Conn.ID(), msg.Room)
					return nil
				},
				neffos.OnRoomLeft: func(c *neffos.NSConn, msg neffos.Message) error {
					log.Printf("[%s] left from room [%s].", c.Conn.ID(), msg.Room)
					return nil
				},
				"chat": func(c *neffos.NSConn, msg neffos.Message) error {
					if !c.Conn.IsClient() {
						log.Printf("--server-side-- send back the message [%s:%s]", msg.Event, string(msg.Body))

						if msg.Room == "" {
							// send back the message to the client.
							// c.Emit(msg.Event, msg.Body) or
							return neffos.Reply(msg.Body)
						}

						c.Conn.Server().Broadcast(c.Conn, neffos.Message{
							Namespace: msg.Namespace,
							Event:     msg.Event,
							// Broadcast to all other members inside this room except this connection(the emmiter, client in this case).
							// If first argument was nil then to all inside this room including this connection.
							Room: msg.Room,
							Body: msg.Body,
						})
					}

					log.Printf("---------------------\n[%s] %s", c.Conn.ID(), msg.Body)
					return nil
				},
			},
		},
	}
)

func newsocket() {
	srv := neffos.New(upgrader, handler)
	srv.OnConnect = func(c *neffos.Conn) error {
		log.Printf("[%s] connected to server.", c.ID())
		// time.Sleep(3 * time.Second)
		// c.Connect(nil, namespace) // auto-connect to a specific namespace.
		// c.Write(namespace, "chat", []byte("Welcome to the server (after namespace connect)"))
		// println("client connected")
		return nil
	}
	srv.OnDisconnect = func(c *neffos.Conn) {
		log.Printf("[%s] disconnected from the server.", c.ID())
	}
	srv.OnUpgradeError = func(err error) {
		log.Printf("ERROR: %v", err)
	}

	log.Printf("Listening on: %s\nPress CTRL/CMD+C to interrupt.", endpoint)
	go http.ListenAndServe(endpoint, srv)

}
