package operators

import (
	"net/http"
	"encoding/json"
	"io/ioutil"
	"fmt"
	"log"
)

type Agents struct {
	Status    json.Number `json:"status"`
	AgentData []Agent     `json:"data"`
}

type Agent struct {
	DisplayName string `json:"displayName"`
	DisplayIcon string `json:"displayIcon"`
}

func GetAgents() Agents {
		resp, err := http.Get("https://valorant-api.com/v1/agents")
		if err != nil {
			log.Fatal(err)
		}

		defer resp.Body.Close()
		body, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			log.Fatal(err)
		}

		var agents Agents
		if err := json.Unmarshal(body, &agents); err != nil {
			log.Fatal(err)
		}

		fmt.Println(agents)
		return agents
}