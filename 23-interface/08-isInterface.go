package main

import "fmt"

type Duck struct {

}

func (d Duck) quack() {
  fmt.Println("꽥~!")
}

func (d Duck) feathers() {
  fmt.Println("오리는 흰색과 회색 털을 가지고 있습니다.")
}

type Quacker interface {
  quack()
  feathers()
}

func inTheForest(q Quacker) {
  q.quack()
  q.feathers()
}

func main() {
  var donald Duck

  if v, ok := interface{}(donald).(Quacker); ok {
    fmt.Println(v, ok)
  }

}
