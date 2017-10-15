package textclassifier

import (
	"github.com/lytics/multibayes"

	"github.com/AKovalevich/scrabbler/route"
)

type TextClassifierEntrypoint struct {
	Instance *multibayes.Classifier
	Routes []route.Route
}

type TrainData struct {
	Text string
	Classes []string
}

type TrainDataList []TrainData

// Create new entrypoint
func New() *TextClassifierEntrypoint {
	entrypoint := &TextClassifierEntrypoint{}
	entrypoint.Instance = multibayes.NewClassifier()
	entrypoint.Instance.MinClassSize = 0

	// Get entrypoint routes
	entrypoint.Routes = GetRoutes()

	return entrypoint
}

// Start entrypoint
func (txe *TextClassifierEntrypoint) Start() {}

// Stop enptrypoint
func (txe *TextClassifierEntrypoint) Stop() {}

// Initialize entrypoint
func (txe *TextClassifierEntrypoint) Init() {
	trainList := TrainDataList{
		{
			Text:    "My dog has fleas.",
			Classes: []string{"vet"},
		},
		{
			Text:    "My cat has ebola.",
			Classes: []string{"vet", "cdc"},
		},
		{
			Text:    "Aaron has ebola.",
			Classes: []string{"cdc"},
		},
	}

	txe.Train(trainList)
}

// Train the classifier
func (txe *TextClassifierEntrypoint) Train(data []TrainData) {
	for _, document := range data {
		txe.Instance.Add(document.Text, document.Classes)
	}
}
