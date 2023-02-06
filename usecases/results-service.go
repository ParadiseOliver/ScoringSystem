package usecases

import "github.com/ParadiseOliver/ScoringSystem/entity"

type ResultsRepository interface {
	AllResultsByEventId(id string) ([]entity.Result, error)
	ResultByResultId(id string) (*entity.Result, error)
	ResultsByAthleteId(id string) ([]entity.Result, error)
}

type resultsService struct {
	repo ResultsRepository
}

func NewResultsService(repo ResultsRepository) *resultsService {
	return &resultsService{
		repo: repo,
	}
}

func (service resultsService) AllResultsByEventId(id string) ([]entity.Result, error) {
	return service.repo.AllResultsByEventId(id)
}

func (service resultsService) ResultByResultId(id string) (*entity.Result, error) {
	return service.repo.ResultByResultId(id)
}

func (service resultsService) ResultsByAthleteId(id string) ([]entity.Result, error) {
	return service.repo.ResultsByAthleteId(id)
}
