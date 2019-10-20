package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"strings"
)

type wordResult struct {
	Word  string   `json:"word"`
	Score int      `json:"score"`
	Tags  []string `json:"tags"`
}

func main() {
	var r []wordResult

	baseURI := "https://api.datamuse.com/words?ml="

	wordToQuery := os.Args[1]

	resp, err := http.Get(baseURI + wordToQuery)
	if err != nil {
		panic(err)
	}

	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	err = json.Unmarshal(body, &r)
	if err != nil {
		panic(err)
	}

	var w []string
	for _, result := range r {
		w = append(w, result.Word)
	}

	fmt.Println("These words are similar to " + wordToQuery + ":")
	fmt.Println(strings.Join(w, ", "))
}
