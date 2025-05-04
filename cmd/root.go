package cmd

func Execute() {
	if err := configureLogger("info"); err != nil {
		panic(err)
	}

	if err := generate("", nil); err != nil {
		panic(err)
	}
}
