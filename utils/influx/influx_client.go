package influx

import (
	"fmt"
	"log"
	"time"
	"github.com/influxdata/influxdb/client/v2"
	"github.com/krise3k/armada-stats/utils"
)

var influxClient client.Client


func GetInfluxClient() *client.Client {
	if influxClient == nil {
		initInfluxClient()
	}

	return &influxClient
}

func initInfluxClient() {
	host, _ := utils.Config.String("influx_host")
	port, _ := utils.Config.String("influx_port")
	db, _ := utils.Config.String("influx_database")
	user, _ := utils.Config.String("influx_user")
	password, _ := utils.Config.String("influx_password")

	addr := fmt.Sprintf("http://%s:%s", host, port)
	c, err := client.NewHTTPClient(client.HTTPConfig{
		Addr: addr,
		Username: user,
		Password: password,
	})
	if err != nil {
		log.Panic("Cannot connect to influx: %v", err)
	}

	defer c.Close()

	dbCreateQuery := fmt.Sprintf("CREATE DATABASE %s", db)
	q := client.NewQuery(dbCreateQuery, "", "")
	if response, err := c.Query(q); err == nil && response.Error() == nil {
		log.Println(response.Results)
	}

	influxClient = c
}

func CreateBatchPoints() client.BatchPoints {
	db, _ := utils.Config.String("influx_database")
	bp, err := client.NewBatchPoints(client.BatchPointsConfig{
		Database:  db,
		Precision: "s",
	})

	if err != nil {
		log.Panic("Error creating batch : %v", err)
	}
	return bp
}

func CreatePoint(name string, tags map[string]string, fields map[string]interface{}) *client.Point {
	pt, err := client.NewPoint(name, tags, fields, time.Now())
	if err != nil {
		log.Printf("Error sending to influx %v", err)
	}

	return pt
}

func Save(points client.BatchPoints) {
	influxClient := *GetInfluxClient()
	err := influxClient.Write(points)
	if err != nil {
		log.Printf("Error sending to influx %v", err)
	}
}