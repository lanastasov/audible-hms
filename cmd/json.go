package audiblejson

func Chapters() {
	// Open our jsonFile
	jsonFile, err := os.Open("chapters.json")
	// if we os.Open returns an error then handle it
	if err != nil {
	    fmt.Println(err)
	}
	fmt.Println("Successfully Opened chapters.json")
	// defer the closing of our jsonFile so that we can parse it later on
	defer jsonFile.Close()

}
