package handlers

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/alexeavru/keks-events/graph"
)

func TestRouter(t *testing.T) {
	srv := handler.NewDefaultServer(graph.NewExecutableSchema(graph.Config{
		Resolvers: &graph.Resolver{},
	}))

	r := Router(srv)
	ts := httptest.NewServer(r)
	defer ts.Close()

	res, err := http.Get(ts.URL + "/readyz")
	if err != nil {
		t.Fatal(err)
	}
	if res.StatusCode != http.StatusOK {
		t.Errorf("Status code for /readyz is wrong. Have: %d, want: %d.", res.StatusCode, http.StatusOK)
	}

	res, err = http.Get(ts.URL + "/healthz")
	if err != nil {
		t.Fatal(err)
	}
	if res.StatusCode != http.StatusOK {
		t.Errorf("Status code for /healthz is wrong. Have: %d, want: %d.", res.StatusCode, http.StatusOK)
	}

	res, err = http.Get(ts.URL + "/login")
	if err != nil {
		t.Fatal(err)
	}
	if res.StatusCode != http.StatusForbidden {
		t.Errorf("Status code for /login is wrong. Have: %d, want: %d.", res.StatusCode, http.StatusForbidden)
	}

	res, err = http.Get(ts.URL + "/query")
	if err != nil {
		t.Fatal(err)
	}
	if res.StatusCode != http.StatusForbidden {
		t.Errorf("Status code for /query is wrong. Have: %d, want: %d.", res.StatusCode, http.StatusForbidden)
	}

	res, err = http.Get(ts.URL + "/")
	if err != nil {
		t.Fatal(err)
	}
	if res.StatusCode != http.StatusForbidden {
		t.Errorf("Status code for / is wrong. Have: %d, want: %d.", res.StatusCode, http.StatusForbidden)
	}

	// b, err := io.ReadAll(res.Body)
	// if err != nil {
	// 	panic(err)
	// }
	// fmt.Printf("%s", b)

}
