package story

import "encoding/json"

func ParseJsonBlob(blob []byte) (map[string]StoryArcObj, error) {
	rootStoryMap := make(map[string]StoryArcObj)
	if err := json.Unmarshal(blob, &rootStoryMap); err != nil {
		return nil, err
	}
	return rootStoryMap, nil
}
