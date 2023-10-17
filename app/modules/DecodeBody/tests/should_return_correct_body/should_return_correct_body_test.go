package should_return_correct_body

import (
	"net/http"
	"testing"

	"github.com/EugeneGpil/tester"

	httpTester "github.com/EugeneGpil/httpTester"
	requestPackage "github.com/EugeneGpil/request"
)

type body struct {
	Test string
}

var testValue = "test"

func Test_should_return_correct_body(t *testing.T) {
	tester.SetTester(t)

	httpRequest := httpTester.GetRequest(httpTester.GetRequestDto{
		Method: http.MethodGet,
		Url:    "/test",
		Body: body{
			Test: testValue,
		},
	})

	var bodyVar body

	err := requestPackage.New(&httpRequest).DecodeBody(&bodyVar)

	tester.AssertNil(err)
	tester.AssertSame(bodyVar.Test, testValue)
}
