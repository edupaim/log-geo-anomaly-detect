package application

import (
	"github.com/tidwall/gjson"
	"fmt"
	"log"
	time2 "time"
	"strings"
)

type Application struct {
	configPath string
	config     *Config
	lastLogObj map[string]*LogObj
}

func NewApplication(configPath string) *Application {
	return &Application{configPath: configPath, lastLogObj: make(map[string]*LogObj)}
}

func (app *Application) RunApplication() {
	app.loadConfig()
	newLog, _ := WatchFile(app.config.FilePath)
	app.handleNewLogEntry(newLog)
}

func (app *Application) handleNewLogEntry(newLog chan string) {
	for newLine := range newLog {
		long := gjson.Get(newLine, app.config.FieldLat).Float()
		if long == 0 {
			continue
		}
		lat := gjson.Get(newLine, app.config.FieldLong).Float()
		if lat == 0 {
			continue
		}
		uid := gjson.Get(newLine, app.config.FieldUserId).String()
		if uid == "" {
			continue
		}
		time := gjson.Get(newLine, app.config.FieldTime)
		t, err := time2.Parse(app.config.FieldTimeFormat, time.String())
		if err != nil {
			continue
		}
		newLogObj := &LogObj{
			Time: t,
			Long: long,
			Lat:  lat,
			LogEntry: strings.Replace(newLine, "\n", "", -1),
		}
		if _, ok := app.lastLogObj[uid]; ok {
			app.calculateDistanceAndTimeLapse(uid, newLogObj)
		}
		app.lastLogObj[uid] = newLogObj
	}
}

func (app *Application) calculateDistanceAndTimeLapse(uid string, newEntry *LogObj) {
	lastEntry := app.lastLogObj[uid]
	distanceM := Distance(newEntry.Lat, newEntry.Long, lastEntry.Lat, lastEntry.Long)
	lapseTime := newEntry.Time.Sub(lastEntry.Time)
	distanceKm := distanceM / 1000
	velo := distanceKm / lapseTime.Hours()
	log.Println("distance:", distanceKm, "velocity:", velo, "user id:", uid)
	if velo > app.config.AcceptableDisplacement {
		app.handleAnomaly(newEntry, uid)
	}
}

func (app *Application) handleAnomaly(newLogObj *LogObj, uid string) {
	log.Println("ANOMALY between log entries!!!")
	log.Println("User ID:", uid)
	log.Println("New entry log:", newLogObj.LogEntry)
	log.Println("Last entry log:", app.lastLogObj[uid].LogEntry)
	log.Println("-!-!-!-!-!-!-!-!-!-!-!-!-!-!-!")
}

func (app *Application) loadConfig() {
	config, err := LoadConfig(app.configPath)
	if err != nil {
		log.Fatalln(err.Error())
	}
	fmt.Println("load config:", config)
	app.config = config
}
