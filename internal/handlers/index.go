package handlers

import (
	"fmt"
	"net/http"

	"github.com/fuadop/my_server/internal/utils"
)

func getIndex(w http.ResponseWriter, r *http.Request) {
	message := fmt.Sprint("Your remote address is: ", r.RemoteAddr)
	utils.HandleResponse(w, 200, message, r.Header)
}

// export all func as handlers
var GetIndex = http.HandlerFunc(getIndex)
