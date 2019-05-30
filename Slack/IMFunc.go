package Slack

import (
	"encoding/json"
	"github.com/DMEvanCT/CommService/Auth"
	"github.com/spf13/viper"
	"log"
	"net/http"
)

type SlackData struct {
	SlackMessage string `json:"SlackMessage"  bson:"SlackMessage"`
	Service string `json:"Service"  bson:"Service"`
}

type Message struct {
	Sent string
	Channel string
	SlackData


}


var service string;

func Communication(w http.ResponseWriter, r *http.Request) {
	var slackData  SlackData


	w.Header().Set("Content-Type", "application/json")

	_ = json.NewDecoder(r.Body).Decode(&slackData)
	akey := r.Header.Get("x-auth-token")
	username := r.Header.Get("x-auth-user")
	//auth := Auth.AuthenticatedUser("c4a3acd4-2ef6-4a5c-b97a-0aa5578503cf", "Clarity")
	auth := Auth.AuthorizeAuthenticate(akey, username, slackData.Service)
	if !auth  {
		w.WriteHeader(http.StatusForbidden)

		authenticated := Auth.Authenticated{
			Authenticated: "False",
		}
		json.NewEncoder(w).Encode(authenticated)

		return


	}


	// Viper configuration this reads from the config file
	viper.AddConfigPath("/etc/commservice/")
	viper.SetConfigName("comconfig")

	err := viper.ReadInConfig()
	if err != nil {
		log.Fatal("Not able to read in config", err)
	}

	// set the token and channel name to the token and channel name inside of the slack block of the yaml
	token := viper.GetString("slack.token")
	channelName := viper.GetString("slack.channel")
	// Selecting service to
	service = slackData.Service
	switch service {
	case "slack":
		SlackMessage(slackData.SlackMessage, token, channelName)
		// Returns to the message and sent to users
		message := Message{
			"Yes",
			channelName,
			SlackData{
				slackData.SlackMessage,
				slackData.Service,
			},
			}





		json.NewEncoder(w).Encode(message)




	default:
		log.Println("Service was requested but does not yet exist." + service)
	}

}
