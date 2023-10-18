package structs

type Text struct {
	Source_language string `json:"source_language"`
	Target_language string `json:"target_language"`
	Text            string `json:"text"`
}

type Translated_data struct {
	Text          string `json:"text"`
	Pronunciation string `json:"pronunciation"`
}
