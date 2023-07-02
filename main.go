package main

import (
	"crypto/tls"
	"encoding/base64"
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/twmb/murmur3"
)

func main() {
	var domain string
	flag.StringVar(&domain, "d", "", "Domain name (e.g., example.com)")
	flag.Parse()

	if domain == "" {
		flag.PrintDefaults()
		log.Fatal("Domain not set")
	}

	url := "http://" + domain + "/favicon.ico"
	fmt.Println(getFaviconHash(url))
}

func getFaviconHash(url string) int32 {
	http.DefaultTransport.(*http.Transport).TLSClientConfig = &tls.Config{InsecureSkipVerify: true}

	final := ""
	fix := 76
	s := make([]string, 0)

	f, err := http.Get(url)
	if err != nil {
		log.Fatal("Error fetching favicon.ico:", err)
	}
	defer f.Body.Close()

	content, err := ioutil.ReadAll(f.Body)
	if err != nil {
		log.Fatal("Error reading favicon.ico data:", err)
	}
	str := base64.StdEncoding.EncodeToString(content)

	// slice up string
	for i := 0; i*fix+fix < len(str); i++ {
		it := str[i*fix : i*fix+fix]
		s = append(s, it)
	}

	// find last piece of string
	findlen := len(s) * fix
	last := str[findlen:] + "\n"

	// put it all together
	for _, s := range s {
		final = final + s + "\n"
	}
	str = final + last

	// do murmurhash3 stuff
	mm3 := murmur3.Sum32([]byte(str))

	// convert uint32 to int32
	return int32(mm3)
}
