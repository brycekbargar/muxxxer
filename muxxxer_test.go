package muxxxer_test

import (
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"net/http/httptest"

	"github.com/brycekbargar/muxxxer"
)

func ExampleMustRegisterRoute_routepriority() {
	muxxxer.MustRegisterRoute(muxxxer.NewRoute("/Users/",
		func(res http.ResponseWriter, req *http.Request) {
			io.WriteString(res, req.URL.String())
		}))
	muxxxer.MustRegisterRoute(muxxxer.NewRoute("/Users/Admins",
		func(res http.ResponseWriter, req *http.Request) {
			panic("Even though this route is a better match it was registered second")
		}))

	serv := httptest.NewServer(&muxxxer.Default)
	defer serv.Close()

	res, err := serv.Client().Get(fmt.Sprintf("%s/Users/Admins", serv.URL))
	if err != nil {
		fmt.Print(err)
		return
	}

	defer res.Body.Close()
	msg, err := ioutil.ReadAll(res.Body)
	if err != nil {
		fmt.Print(err)
		return
	}

	fmt.Println(res.StatusCode, string(msg))
	// Output: 200 /Users/Admins
}

func ExampleMustRegisterRoute_argumentroutes() {
	muxxxer.MustRegisterRoute(muxxxer.NewArgumentRoute("/Admin/;id",
		func(res http.ResponseWriter, req *http.Request, args *muxxxer.ArgumentBag) {
			if err := args.Parse(); err != nil {
				http.Error(res, err.Error(), http.StatusBadRequest)
				return
			}

			res.Header().Add("Content-Type", "application/json")
			j := json.NewEncoder(res)
			j.Encode(args)
		}))

	serv := httptest.NewServer(&muxxxer.Default)
	defer serv.Close()

	res, err := serv.Client().Get(fmt.Sprintf("%s/Admin/123?filter=country:us&filter=age:31", serv.URL))
	if err != nil {
		fmt.Print(err)
		return
	}

	defer res.Body.Close()
	msg, err := ioutil.ReadAll(res.Body)
	if err != nil {
		fmt.Print(err)
		return
	}

	fmt.Println(res.StatusCode, string(msg))
	// Output: 200 {"Args":{"filter":["country:us","age:31"],"id":["123"]}}
}
