package main

import "fmt"

func main () {
  terrestriaPlanet := map[string]map[string]float32{
    "Mercury" : map[string]float32{
      "meanRadius" : 2439.7,
      "mass" : 3.3022E+23,
      "orbitalPeriod" : 87.969,
    },
    "Venus" : map[string]float32{
      "meanRadius" : 6051.8,
      "mass" : 4.8676E+24,
      "orbitalPeriaod" : 224.70069,
    },
    "Earth" : map[string]float32{
      "meanRadius": 6371.0,
      "mass" : 5.97219E+24,
      "prbitalPeriod" : 365.25641,
    },
    "Mars" : map[string]float32{
      "meanRadius" : 3389.5,
      "mass" : 6.4185E+23,
      "orbitalPeriaod" : 686.9600,
    },
  }

  fmt.Println(terrestriaPlanet["Mars"]["mass"])
}
