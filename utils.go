package messagebroker

import (
	"encoding/json"
	"errors"
)

func DecodeMessage(message interface{}, result interface{}) error {
	dataByte, ok := message.([]byte)
	if ok {
		err := json.Unmarshal(dataByte, result)
		if err != nil {
			return err
		}

		return nil
	}

	dataStr, ok := message.(string)
	if ok {
		err := json.Unmarshal([]byte(dataStr), result)
		if err == nil {
			return nil
		}

		result = dataStr
	}

	return errors.New("cannot decode message")
}
