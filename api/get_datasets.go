package api

import (
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

var githubBaseURL = "https://api.github.com/repos/"

func GetDataset(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)

	path := r.FormValue("path")
	owner := params["owner"]
	repo := params["repo"]

	log.Printf("Searching %s in %s/%s...\n", path, owner, repo)

	resp, err := requestGithubDataset(owner, repo, path)

	if err != nil {
		http.Error(w, "Error requesting to github", http.StatusInternalServerError)
	}

	log.Printf("Found %s in %s/%s!\n", path, owner, repo)

	result, _ := searchForFiles(resp)

	response, _ := json.Marshal(result)
	w.Header().Set("Content-Type", "application/json")
	w.Write(response)
}

func requestGithubDataset(owner string, repo string, path string) ([]map[string]string, error) {
	requestURL := githubBaseURL + owner + "/" + repo + "/contents" + path
	return requestGithubDatasetWithURL(requestURL)
}

func requestGithubDatasetWithURL(requestURL string) ([]map[string]string, error) {

	resp, err := http.Get(requestURL)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	if resp.StatusCode >= 400 {
		return nil, errors.Unwrap(fmt.Errorf("Error %d requesting to github", resp.StatusCode))
	}

	var result []map[string]string
	json.NewDecoder(resp.Body).Decode(&result)

	return result, nil
}

func searchForFiles(dir []map[string]string) (map[string]interface{}, error) {
	result := map[string]interface{}{}
	for _, e := range dir {
		if e["type"] == "file" {
			result[e["name"]] = e["download_url"]
		} else {
			newDir, err := requestGithubDatasetWithURL(e["url"])
			if err != nil {
				return nil, err
			}

			result[e["name"]], err = searchForFiles(newDir)
			if err != nil {
				return nil, err
			}
		}
	}

	return result, nil
}
