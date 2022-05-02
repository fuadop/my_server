package handlers

import (
	"log"
	"net/http"

	"github.com/fuadop/my_server/internal/utils"
	"github.com/fuadop/sendchamp"
)

func SendVoice(c *sendchamp.Client) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		qs := r.URL.Query()
		res, err := c.NewVoice().Send([]string{"2348153207998"}, qs.Get("message"), sendchamp.VoiceTypeOutgoing, 2)
		if err != nil {
			utils.HandleResponse(w, 500, err.Error(), err)
		}

		log.Println(res.Data.ID)
		utils.HandleResponse(w, 201, res.Message, res)
	})
}
