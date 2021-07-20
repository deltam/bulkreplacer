package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"strings"
)

func main() {
	var keywordJSON string
	flag.StringVar(&keywordJSON, "keyword", "", "keyword list json")
	var formatOld string
	flag.StringVar(&formatOld, "old", "_[%d]", `Replace Format(ex. "_array[%d]"`)
	var formatNew string
	flag.StringVar(&formatNew, "new", `"%s"`, `Replace Format(ex. '"%s"'`)

	flag.Parse()

	bs, err := ioutil.ReadFile(keywordJSON)
	if err != nil {
		log.Fatalf("ioutil.ReadFile(%s) failed: %v\n", keywordJSON, err)
	}

	var keywords []string
	if err := json.Unmarshal(bs, &keywords); err != nil {
		log.Fatalf("json.Unmarshal() failed: %v", err)
	}

	reps := make([]string, len(keywords)*2)
	for i, k := range keywords {
		reps[i*2] = fmt.Sprintf(formatOld, i)
		reps[i*2+1] = fmt.Sprintf(formatNew, k)
	}

	replacer := strings.NewReplacer(reps...)

	buf, err := ioutil.ReadAll(os.Stdin)
	if err != nil {
		log.Fatalf("ReadAll() failed: %v", err)
	}

	if _, err := replacer.WriteString(os.Stdout, string(buf)); err != nil {
		log.Fatalf("replacer.WriteString() failed: %v", err)
	}
}
