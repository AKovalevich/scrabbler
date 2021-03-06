package textclassifier

import (
	"net/http"
	"fmt"

	"github.com/lytics/multibayes"
	"github.com/AKovalevich/scrabbler/route"
)

type Entrypoint struct {
	Name string
	Instance *multibayes.Classifier
	Routes []route.Route
}

type TrainData struct {
	Text string
	Classes []string
}

type TrainDataList []TrainData

// Create new entrypoint
func New() *Entrypoint {
	entrypoint := &Entrypoint{}
	entrypoint.Instance = multibayes.NewClassifier()
	entrypoint.Instance.MinClassSize = 0

	return entrypoint
}

func (txe *Entrypoint) RoutesList() []route.Route {
	return txe.Routes
}

// Start entrypoint
func (txe *Entrypoint) Start() {}

// Stop enptrypoint
func (txe *Entrypoint) Stop() {}

// Initialize entrypoint
func (txe *Entrypoint) Init() {
	txe.Routes = []route.Route{
		{
			Path: "/train",
			Handler: func(w http.ResponseWriter, r *http.Request) {
				fmt.Fprint(w, "Welcome!\n")
			},
		},
		{
			Path: "/test-train",
			Handler: func(w http.ResponseWriter, r *http.Request) {
				txe.TestTrain()
				fmt.Fprint(w, "Done!\n")
			},
		},
		{
			Path: "/get-result",
			Handler: func(w http.ResponseWriter, r *http.Request) {
				data := txe.GetData()
				fmt.Fprintf(w,"Posterior Probabilities: %+v\n", data)
			},
		},
	}
}

func (txe *Entrypoint) String() string {
	return txe.Name
}

// Train the classifier
func (txe *Entrypoint) Train(data []TrainData) {
	for _, document := range data {
		txe.Instance.Add(document.Text, document.Classes)
	}
}

func (txe *Entrypoint) TestTrain() {
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

func (txe *Entrypoint) GetData() map[string]float64 {
	return txe.Instance.Posterior("dog")
}
