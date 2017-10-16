package profanity

import (
	"net/http"
	"fmt"

	"github.com/BurntSushi/toml"
	"github.com/lytics/multibayes"
	"github.com/AKovalevich/scrabbler/route"
	log "github.com/AKovalevich/scrabbler/log/logrus"
)

type Entrypoint struct {
	Name string
	Instance *multibayes.Classifier
	Routes []route.Route
}

type TrainData struct {
	Text string 		`toml:"text"`
	Classes []string 	`toml:"classes"`
}

type TrainDataList struct {
	TrainData []TrainData
}

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
			Path: "/profanity/train",
			Handler: func(w http.ResponseWriter, r *http.Request) {
				fmt.Fprint(w, "Welcome!\n")
			},
		},
		{
			Path: "/profanity/test-train",
			Handler: func(w http.ResponseWriter, r *http.Request) {
				txe.TestTrain()
				fmt.Fprint(w, "Done!\n")
			},
		},
		{
			Path: "/profanity/get-result",
			Handler: func(w http.ResponseWriter, r *http.Request) {
				text := r.URL.Query().Get("text")
				data := txe.GetData(text)
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
	trainList := TrainDataList{}

	if _, err := toml.DecodeFile("entrypoint/profanity/train.ru.data.toml", &trainList); err != nil {
		log.Do.Error(err)
	}

	txe.Train(trainList.TrainData)
}

func (txe *Entrypoint) GetData(text string) map[string]float64 {
	return txe.Instance.Posterior(text)
}
