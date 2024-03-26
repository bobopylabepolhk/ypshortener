package healthcheck

type (
	HealthcheckRepository interface {
		Ping() error
	}

	HealthcheckService struct {
		repo HealthcheckRepository
	}
)

func (hs *HealthcheckService) PingDB() error {
	return hs.repo.Ping()
}

func NewHealthcheckService(repo HealthcheckRepository) *HealthcheckService {
	return &HealthcheckService{
		repo: repo,
	}
}
