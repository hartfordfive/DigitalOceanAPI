package utils

import (
	"bytes"
	"errors"
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
	"time"
)

type Config struct {
	ApiToken string
	Debug    int
}

func LoadConfig(filename string, conf *Config) error {

	valid := map[string]int{
		"debug":     1,
		"api_token": 1,
	}

	buf := bytes.NewBuffer(nil)
	f, err := os.Open(filename) // Error handling elided for brevity.
	if err != nil {
		return errors.New("Invalid or missing config!")
	}

	io.Copy(buf, f) // Error handling elided for brevity.
	f.Close()
	s := string(buf.Bytes())

	for _, l := range strings.Split(strings.Trim(s, " "), "\n") {
		// Ignore line that begins with #
		if l == "" || string(l[0]) == "#" {
			continue
		}
		parts := strings.SplitN(strings.Trim(l, " "), "=", 2)
		if _, ok := valid[parts[0]]; ok {
			if parts[0] == "debug" {
				v, _ := strconv.Atoi(parts[1])
				if v < 0 || v > 1 {
					fmt.Println(DateStampAsString(), "Config ERROR: debug can only be 0 or 1")
					os.Exit(1)
				}
				conf.Debug = v
			} else if parts[0] == "api_token" {
				conf.ApiToken = parts[1]
			}
		}
	}
	return nil
}

func YmdToString() string {
	t := time.Now()
	y, m, d := t.Date()
	return strconv.Itoa(y) + fmt.Sprintf("%02d", m) + fmt.Sprintf("%02d", d)
}
func DateStampAsString() string {
	t := time.Now()
	return "[" + YmdToString() + " " + fmt.Sprintf("%02d", t.Hour()) + ":" + fmt.Sprintf("%02d", t.Minute()) + ":" + fmt.Sprintf("%02d", t.Second()) + "]"
}
