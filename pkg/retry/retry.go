package retry

type RetryableFunc func() error

func Do(fn RetryableFunc, maxAttemps int) error {
	var err error
	for i := 0; i < maxAttemps; i++ {
		err = fn()
		if err == nil {
			break
		}
	}

	return err
}
