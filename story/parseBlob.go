package story

import "encoding/json"

func ParseJsonBlob(blob []byte) (Story, error) {
	story := make(Story)
	if err := json.Unmarshal(blob, &story); err != nil {
		return nil, err
	}
	return story, nil
}
