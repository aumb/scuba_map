package utils

import (
	"encoding/json"
	"fmt"
)

func QueryAndUnmarshal(f func() ([]byte, int64, error), r interface{}) error {
	result, _, err := f()

	if err != nil {
		fmt.Println(err)
		return err
	}

	return json.Unmarshal(result, &r)
}
