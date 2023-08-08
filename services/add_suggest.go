package services

import (
	"suggester/models"
	"suggester/repository"
)

func AddSuggest(suggest models.Suggest, repo repository.SuggestRepoInterface) error {
	if existSuggest, err := repo.GetOne(suggest.Text); err != nil {
		return repo.Add(suggest)
	} else {
		return repo.IncrementAmount(existSuggest.Id)
	}
}
