package rctest

import (
	"github.com/mleuth/restclient"
	"net/http"
)

func Status200() restclient.Result {
	return Status(http.StatusOK)
}

func Status204() restclient.Result {
	return Status(http.StatusNoContent)
}

func Status400() restclient.Result {
	return Status(http.StatusBadRequest)
}

func Status404() restclient.Result {
	return Status(http.StatusNotFound)
}

func Status(code int) restclient.Result {
	return restclient.Result{StatusCode: code}
}
