package usecase

func (uc *Usecase) GetHealthCheck() (map[string]string, error) {
	result, err := uc.repository.GetHealthCheck()
	if err != nil {
		return nil, err
	}

	return result, nil
}
