package main

func main() {
	whereErrorsGoToDie(nil)
}

func whereErrorsGoToDie(err error) {
	if nil != err {
		panic(err)
	}
}
