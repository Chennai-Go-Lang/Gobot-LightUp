package main

import (
	"fmt"
	"strconv"

	"gobot.io/x/gobot"
	"gobot.io/x/gobot/api"
	"gobot.io/x/gobot/drivers/gpio"
	"gobot.io/x/gobot/platforms/raspi"
)

func main() {
	master := gobot.NewMaster()
	a := api.NewAPI(master)
	a.Debug()
	a.Start()

	r := raspi.NewAdaptor()
	robot := gobot.NewRobot("Oom",
		[]gobot.Connection{r},
		func() {},
	)
	leds := buildLedArray(r, robot)

	master.AddRobot(robot)
	master.AddCommand("SetCount", func(params map[string]interface{}) interface{} {
		count, _ := strconv.ParseUint(params["count"].(string), 10, 64)
		lightUp(leds, uint32(count))
		return fmt.Sprintf("Lights the leds as per count")
	})
	master.Start()
}

func lightUp(leds []*gpio.LedDriver, count uint32) {
	maxCount := uint32(100000000)

	for _, led := range leds {
		if count >= maxCount {
			fmt.Println("Turning on LED")
			led.On()
		} else {
			led.Off()
		}
		maxCount = maxCount / 10
	}
}

func buildLedArray(r *raspi.Adaptor, robot *gobot.Robot) (leds []*gpio.LedDriver) {
	ledPins := []string{"7", "11", "13", "15", "19", "21", "23", "29"}
	for _, pin := range ledPins {
		led := gpio.NewLedDriver(r, pin)
		robot.AddDevice(led)
		led.On()
		leds = append(leds, led)
	}
	return leds
}
