package main

import "fmt"

func main () {
  solarSystem := make(map[string]float32)
  solarSystem["Mercury"] = 87.969
  solarSystem["venus"] = 224.70069
  solarSystem["Earth"] = 365.25641
  solarSystem["Mars"] = 686.9600
  solarSystem["Jupiter"] = 4333.2867
  solarSystem["Saturn"] = 10756.1995
  solarSystem["Uranus"] = 30707.4896
  solarSystem["Neptune"] = 60223.3528

  fmt.Println(solarSystem["Pluto"])
}
