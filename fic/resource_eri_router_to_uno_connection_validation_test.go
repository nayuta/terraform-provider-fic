package fic

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	ficclient "github.com/nttcom/go-fic"
)

func TestUNOInterconnectValidation(t *testing.T) {
	validate := resourceEriRouterToUNOConnectionV1().Schema["destination_interconnect"].ValidateFunc
	for _, tc := range []struct {
		value string
		valid bool
	}{
		{"Interconnect-Tokyo-1", true},
		{"Interconnect-Tokyo-104", true},
		{"Interconnect-Osaka-1", true},
		{"Interconnect-Tokyo-999", false},
	} {
		t.Run(tc.value, func(t *testing.T) {
			_, errors := validate(tc.value, "destination_interconnect")
			if (len(errors) == 0) != tc.valid {
				t.Fatalf("validation errors: %v", errors)
			}
		})
	}
}

func TestUNOReadPreservesQoSFromAPI(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		_, _ = w.Write([]byte(`{"connection":{"id":"connection-1","source":{"routeFilter":{}},"destination":{"interconnect":"Interconnect-Tokyo-104","qosType":"guarantee","routeFilter":{}}}}`))
	}))
	defer server.Close()

	client := &ficclient.ProviderClient{
		EndpointLocator: func(ficclient.EndpointOpts) (string, error) { return server.URL + "/", nil },
	}
	d := schema.TestResourceDataRaw(t, resourceEriRouterToUNOConnectionV1().Schema, nil)
	d.SetId("connection-1")
	if err := resourceEriRouterToUNOConnectionV1Read(d, &Config{OsClient: client}); err != nil {
		t.Fatal(err)
	}
	if got := d.Get("destination_qos_type"); got != "guarantee" {
		t.Fatalf("destination_qos_type = %v, want guarantee", got)
	}
}
