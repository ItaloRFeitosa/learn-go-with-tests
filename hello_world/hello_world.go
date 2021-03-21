package hello_world

const HELLO = "Hello "

func Hello(name string) string {
	return HELLO + name
}

func HelloWorld() string {
	return Hello("World")
}