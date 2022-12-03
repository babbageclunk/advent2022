package lib

import "io/ioutil"

func Read(filename string) string {
	data, err := ioutil.ReadFile(filename)
	Check(err)
	return string(data)
}

func Check(err error) {
	if err != nil {
		panic(err)
	}
}
