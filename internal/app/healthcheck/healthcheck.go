package healthcheck

type (
	HealthcheckRepository interface {
		Ping() error
	}

	HealthcheckService struct {
		repo HealthcheckRepository
	}
)

func (hs *HealthcheckService) PingDb() error {
	return hs.repo.Ping()
}

func NewHealthcheckService(repo HealthcheckRepository) *HealthcheckService {
	return &HealthcheckService{
		repo: repo,
	}
}
