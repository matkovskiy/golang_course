package external_api

import (
	"encoding/json"
	"fmt"
	"hw10/internal/structs"

	gtranslate "github.com/gilang-as/google-translate"
)

func Translate(source_language, target_language, text string) (output string) {
	value := gtranslate.Translate{
		Text: text,
		From: source_language,
		To:   target_language,
	}
	var Asnwer_data structs.Translated_data
	translated, err := gtranslate.Translator(value)
	if err != nil {
		panic(err)
	} else {
		prettyJSON, err := json.MarshalIndent(translated, "", "\t")
		err = json.Unmarshal(prettyJSON, &Asnwer_data)
		if err != nil {
			panic(err)
		}
		fmt.Printf("Answer=%s", Asnwer_data.Text)
		return Asnwer_data.Text
	}

}
