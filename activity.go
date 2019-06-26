package sample

import (
	"strconv"

	"github.com/MichaelS11/go-dht"
	"github.com/project-flogo/core/activity"
	"github.com/project-flogo/core/data/metadata"
)

func init() {
	_ = activity.Register(&Activity{}) //activity.Register(&Activity{}, New) to create instances using factory method 'New'
}

var activityMd = activity.ToMetadata(&Settings{}, &Input{}, &Output{})

//New optional factory method, should be used if one activity instance per configuration is desired
func New(ctx activity.InitContext) (activity.Activity, error) {

	s := &Settings{}
	err := metadata.MapToStruct(ctx.Settings(), s, true)
	if err != nil {
		return nil, err
	}

	ctx.Logger().Debugf("Setting: %s", s.ASetting)

	act := &Activity{} //add aSetting to instance

	return act, nil
}

// Activity is an sample Activity that can be used as a base to create a custom activity
type Activity struct {
}

// Metadata returns the activity's metadata
func (a *Activity) Metadata() *activity.Metadata {
	return activityMd
}

// Eval implements api.Activity.Eval - Logs the Message
func (a *Activity) Eval(ctx activity.Context) (done bool, err error) {

	err = dht.HostInit()
	if err != nil {
		return true, err
	}

	dht, err := dht.NewDHT("GPIO19", dht.Celsius, "")
	if err != nil {
		return true, err
	}

	humidity, temperature, err := dht.ReadRetry(100)
	if err != nil {
		return true, err
	}

	temp := strconv.FormatFloat(temperature, 'f', 6, 64)
	humy := strconv.FormatFloat(humidity, 'f', 6, 64)

	output := &Output{AnOutput: temp + " - " + humy}
	err = ctx.SetOutputObject(output)
	if err != nil {
		return true, err
	}

	return true, nil
}
