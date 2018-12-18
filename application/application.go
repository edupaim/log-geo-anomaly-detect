package application

import (
	"github.com/tidwall/gjson"
	"fmt"
	"log"
	time2 "time"
)

type Application struct {
	configPath string
	config     *Config
	lastLogObj *LogObj
}

func NewApplication(configPath string) *Application {
	return &Application{configPath: configPath}
}

func (app *Application) RunApplication() {
	app.loadConfig()
	newLog, _ := WatchFile(app.config.FilePath)
	app.handleNewLogEntry(newLog)
}

func (app *Application) handleNewLogEntry(newLog chan string) {
	for newLine := range newLog {
		Long := gjson.Get(newLine, app.config.FieldLat)
		Lat := gjson.Get(newLine, app.config.FieldLong)
		time := gjson.Get(newLine, app.config.FieldTime)
		t, err := time2.Parse(app.config.FieldTimeFormat, time.String())
		if err != nil {
			log.Println("time parse error!", err.Error())
			continue
		}
		newLogObj := &LogObj{
			Time: t,
			Long: Long.Float(),
			Lat:  Lat.Float(),
		}
		log.Println("newLogObj:", newLogObj)
		if app.lastLogObj != nil {
			log.Println("lastLogObj:", app.lastLogObj)
			distanceM := Distance(newLogObj.Lat, newLogObj.Long, app.lastLogObj.Lat, app.lastLogObj.Long)
			lapseTime := newLogObj.Time.Sub(app.lastLogObj.Time)
			distanceKm := distanceM/1000
			velo := distanceKm/lapseTime.Hours()
			log.Println("distance:", distanceKm, "velocity:", velo)
		}
		app.lastLogObj = newLogObj
	}
}

func (app *Application) loadConfig() {
	config, err := LoadConfig(app.configPath)
	if err != nil {
		log.Fatalln(err.Error())
	}
	fmt.Println("load config:", config)
	app.config = config
}
