package setup

import (
	"errors"
	"fmt"
	"os"
	"strings"
)

func ValidateEnvironment(input []string) error {

	var missing []string
	checkEnv := func(env string) {
		if os.Getenv(env) == "" {
			missing = append(missing, env)
		}
	}

	for _, e := range input {
		checkEnv(e)
	}

	if len(missing) > 0 {
		m := strings.Join(missing, ", ")
		cause := fmt.Sprint("Missing Environment Variables: ", m)
		return errors.New(cause)
	}
	return nil
}
