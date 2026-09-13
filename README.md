# Libeloula
A [Dragonfly](https://github.com/df-mc/dragonfly) library focusing on really basic admin features (for now).

> [!NOTE]
> This library is developed for educational purposes. Honestly I don't know if anything here is useful, and some stuff might not work properly!

## Features
The plugin has several features. In order for everything to work, you need to initialize the basic functions of the library:
```go
libeloula.Initialize(srv) // srv is the Server object initialized in main.go
```

### Console Command Source
This library introduces a basic console buffer which reads input (commands), executes them and prints the standard command output from the `cmd.Output`.

In order for the buffer to start working, you need to execute asynchronously the following function:
```go
go libeloula.StartConsoleBuffer()
```

### Packet Writing
For Minecraft player features not included in the code base (like Operator assignment), there is a trick to write packets from the Player Connection object (`*minecraft.Conn`). The server listeners catch the connections when being created, using the `connection.ConnectionListener` listeners. Then, we must change all pre-registered default Listeners to ours using this block of code, placed in `main.go`:
```go
for i, factory := range conf.Listeners {
	originalFactory := factory

	conf.Listeners[i] = func(c server.Config) (server.Listener, error) {
		l, err := originalFactory(c)
		if err != nil {
			return nil, err
		}

		return &connection.ConnectionListener{Listener: l}, nil
	}
}
```

Then, we can send any packets through the server's Minecraft protocol, for instance:
```go
func SendChatMessagePacket(handle *world.EntityHandle) {
	msg := &packet.Text{
		TextType:         1, //chat type
		NeedsTranslation: false,
		SourceName:       "Server",
		Message:          "Welcome to the server fellow player!",
		Parameters:       []string{},
		PlatformChatID:   "",
	}

	err = connection.GetConn(handle.UUID().String()).WritePacket(msg)
	Check(err) // checks for errors and panics :o
}
```

### Commands
This library attempts to create all the useful vanilla commands. The list is as follows:

- `/op <player:string>`: Gives operator permissions to specified player