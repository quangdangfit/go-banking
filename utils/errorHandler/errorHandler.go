package errorHandler

import (
	"encoding/json"
	"gitlab.com/quangdangfit/gocommon/utils/logger"
	"go-banking/utils/response"
	"log"
	"net/http"
)

func PanicHandler(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		defer func() {
			error := recover()
			if error != nil {
				log.Println(error)

				resp := response.PrepareResponse(nil, "Internal server error", "")
				json.NewEncoder(w).Encode(resp)
			}
		}()
		next.ServeHTTP(w, r)
	})
}

func HandleErr(err error) {
	if err != nil {
		logger.Error(err.Error())
	}
}
