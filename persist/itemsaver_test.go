package persist

import (
	"crawler/model"
	"testing"
)

func TestItemSaver(t *testing.T) {
	save(model.Profile{
		Name: "name",
		Car:  "car",
	})
}
