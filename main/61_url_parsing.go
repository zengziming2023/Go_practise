package main

import (
	"fmt"
	"net"
	"net/url"
)

func main() {
	s := "postgres://user:pass@host.com:5432/path?k=v#f"

	u, err := url.Parse(s)
	if err != nil {
		panic(err)
	}
	fmt.Println(u.Scheme)
	fmt.Println(u.User)
	fmt.Println(u.User.Username())
	password, _ := u.User.Password()
	fmt.Println(password)

	fmt.Println(u.Host)
	host, port, err := net.SplitHostPort(u.Host)
	if err != nil {
		return
	}
	fmt.Println(host)
	fmt.Println(port)

	fmt.Println(u.Path)
	fmt.Println(u.Fragment)
	fmt.Println(u.RawQuery)
	query, err := url.ParseQuery(u.RawQuery)
	if err != nil {
		return
	}
	fmt.Println(query)
	fmt.Println(query.Get("k"))
}
