package retrier

import "time"

type Try func() error

func DoTry(count uint, try Try) error {
	var err error

	timeout := time.Second

	for range count {
		if err = try(); err == nil {
			return nil
		}

		time.Sleep(timeout)

		timeout *= 2
	}

	return err
}

func Connect[T any](count uint, conn func() (T, error)) (T, error) {
	var (
		err   error
		value T
	)

	timeout := time.Second

	for range count {
		value, err = conn()
		if err == nil {
			return value, nil
		}

		time.Sleep(timeout)

		timeout *= 2
	}

	return value, err
}
