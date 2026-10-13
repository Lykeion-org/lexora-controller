package models

type AnalyzeTextRequestInternal struct {
	JobId string `json:"job_id"`
	RequesterId string `json:"requester_id"`
	Language string `json:"language"`
	Text string `json:"text"`
}

type AnalyzeTextResponseInternal struct {
	RequesterId string `json:"requester_id"`
	JobId string `json:"job_id"`
	OriginalText string `json:"original_text"`
	Analysis TextAnalysis `json:"analysis"`
}

type TextAnalysis struct {
	Objects 	[]WordAnalysis `json:"objects"`
	Actions		[]WordAnalysis `json:"actions"`
	Modifiers	[]WordAnalysis `json:"modifiers"`
}

type WordAnalysis struct {
	Text string `json:"text"`
	Lemma string `json:"lemma"`
	Pos string `json:"pos"`
	Attributes WordAnalysisAttributes `json:"attributes"`
}

type WordAnalysisAttributes struct {
	Gender int `json:"gender"` 
    Number int `json:"number"`
    Case int `json:"case"`
    Mood int  `json:"mood"` 
    Tense int `json:"tense"` 
    Aspect int `json:"aspect"` 
    Person int `json:"person"`
    Diminutive bool `json:"diminutive"` 
    DirectObject bool `json:"direct_object"` 
    IndirectObject bool `json:"indirect_object"` 
    Valency int `json:"valency"`
    Reflexive bool `json:"reflexive"` 
    Unique bool `json:"unique"` 
}