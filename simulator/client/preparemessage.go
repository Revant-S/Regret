package client

import (
	"Regret/domain"
	"net/http"
	"strconv"
)

func SetTickHeader(request *http.Request) {
	currentTick := GetTickInstance().Get()
	request.Header.Set(domain.TickHeaderString, strconv.FormatUint(currentTick, 10))
}
