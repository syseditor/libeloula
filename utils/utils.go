package utils

func Check(err error) {
	if err != nil {
		panic(err)
	}
}

// Include all required procedures for safe-termination (f.e. closing all db connections, saving data, etc)
func TerminateServer() {
	Server.Close()
}
