package main

import (
     "github.com/fsnotify/fsnotify"
	 "github.com/garyburd/redigo/redis"
)

func main() {
    fmt.Println(stringUtil.Reverse("@#oG ,olleH"));
	fmt.Println(math.Exp2(10))
	println("Beego version:", beego.VERSION);

	// querystring/query
	type Options struct {
	  Query   string `url:"q"`
	  ShowAll bool   `url:"all"`
	  Page    int    `url:"page"`
	}

	opt := Options{ "foo", true, 2 }
	v, _ := query.Values(opt)
	fmt.Print(v.Encode()) // will output: "q=foo&all=true&page=2"
	
	/*port, err := freeport.GetFreePort()
	if err != nil {
		println(err)
	} else {
		println(port)
	}*/
}



