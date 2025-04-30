package main

func main() {
	if err := ConfigureLogger("debug"); err != nil {
		panic(err)
	}

	if err := generate(""); err != nil {
		panic(err)

	}
}
