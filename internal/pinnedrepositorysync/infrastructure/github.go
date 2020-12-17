package infrastructure

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	pr "github.com/yehudamakarov/personal-site-proto/packages/go/pinnedRepository"
)

type GithubService struct {
	credentials string
	client      *http.Client
	apiUrl      string
}

func NewGithubService(credentials string, client *http.Client) *GithubService {
	return &GithubService{credentials: credentials, client: client, apiUrl: "https://api.github.com/graphql"}
}

func (g GithubService) FetchPinnedRepositories() ([]pr.PinnedRepository, error) {
	resp, err := g.makeRequest(`{
		  user(login: "yehudamakarov") {
			pinnedItems(types: REPOSITORY, first: 6) {
			  nodes {
				... on Repository {
				  name
				  description
				  databaseId
				  url
				  createdAt
				  updatedAt
				}
			  }
			}
		  }
		}`,
	)
	if err != nil {
		return nil, err
	}

	prs, err := getPrs(resp)
	if err != nil {
		return nil, err
	}

	return prs, nil
}

func getPrs(resp *http.Response) ([]pr.PinnedRepository, error) {
	respData, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	var prs []pr.PinnedRepository
	err = json.Unmarshal(respData, &prs)
	if err != nil {
		return nil, err
	}

	err = resp.Body.Close()
	if err != nil {
		return nil, err
	}

	return prs, nil
}

func (g GithubService) makeRequest(graphqlQuery string) (*http.Response, error) {
	requestBody, err := json.Marshal(
		map[string]string{
			"query": graphqlQuery,
		},
	)
	if err != nil {
		return nil, err
	}

	req, err := http.NewRequest("POST", g.apiUrl, bytes.NewBuffer(requestBody))
	if err != nil {
		return nil, err
	}

	req.Header.Set("Authorization", fmt.Sprintf("Bearer %s", g.credentials))

	response, err := g.client.Do(req)
	if err != nil {
		return nil, err
	}
	return response, nil
}
