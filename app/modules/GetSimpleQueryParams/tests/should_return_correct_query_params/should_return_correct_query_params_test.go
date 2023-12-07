package should_return_correct_query_params

import (
	"net/http"
	"testing"

	"github.com/EugeneGpil/tester"

	httpTester "github.com/EugeneGpil/httpTester"
	requestPackage "github.com/EugeneGpil/request"
)

func Test_should_return_correct_query_params(t *testing.T) {
	tester.SetTester(t)

	httpRequest := httpTester.GetRequest(httpTester.GetRequestDto{
		Method: http.MethodGet,
		Url: "/test",
		Query: map[string]string {
			"test_key": "test_value",
		},
	})

	queryParams := requestPackage.New(&httpRequest).GetSimpleQueryParams()

	tester.AssertSame("test_value", queryParams["test_key"])
}