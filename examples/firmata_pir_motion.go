package main

import (
	"github.com/hybridgroup/gobot"
	"github.com/hybridgroup/gobot/drivers/gpio"
	"github.com/hybridgroup/gobot/platforms/firmata"
)

func main() {
	firmataAdaptor := firmata.NewAdaptor("/dev/ttyACM0")

	sensor := gpio.NewPIRMotionDriver(firmataAdaptor, "5")
	led := gpio.NewLedDriver(firmataAdaptor, "13")

	work := func() {
		sensor.On(gpio.MotionDetected, func(data interface{}) {
			led.On()
		})
		sensor.On(gpio.MotionStopped, func(data interface{}) {
			led.Off()
		})
	}

	robot := gobot.NewRobot("motionBot",
		[]gobot.Connection{firmataAdaptor},
		[]gobot.Device{sensor, led},
		work,
	)

	robot.Start()
}
