package main

func main() {

	log.SetLogger(log.AdapterFile, `{"filename":"temp/project.log","level":7,"maxlines":20000,"maxsize":1048576,"daily":true,"maxdays":10}`)

	for index := 0; index < 50000; index++ {
		//an official log.Logger with prefix ORM
		log.GetLogger("ORM").Println("this is a message of orm")

		log.Debug("my book is bought in the year of ", 2016)
		log.Info("this %s cat is %v years old", "yellow", 3)
		log.Warn("json is a type of kv like", map[string]int{"key": 2016})
		log.Error(1024, "is a very", "good game")

	}

}
