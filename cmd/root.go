package cmd

func Execute() {
	if err := configureLogger("info"); err != nil {
		panic(err)
	}

	if err := Generate("", nil); err != nil {
		panic(err)
	}
}
