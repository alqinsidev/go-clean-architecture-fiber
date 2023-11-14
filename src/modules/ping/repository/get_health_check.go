package repository

func (r *Repository) GetHealthCheck() (map[string]string, error) {
	status := map[string]string{
		"status": "service available",
	}
	return status, nil
}
