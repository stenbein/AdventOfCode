package day18

import scala.io.Source
//import scala.util.Using

object Main extends App {

  val filename = "resources/input"

/*  val lines = Using(Source.fromFile(filename)){
    file => file.getLines().toList
  }.get*/

  val lines = Source.fromFile(filename).getLines().toList

  def part1() {

    var g = new Grid(

      lines.iterator.map(s =>
        s.toCharArray.map{
          case '#' => 1
          case _ => 0
        }.toList
      ).toList

    )

    for (i <- 1 to 100) {
      g = g.step()
    }
    println(g.numOn())

  }

  def part2() {

    var g = new Grid(

      lines.iterator.map(s =>
        s.toCharArray.map{
          case '#' => 1
          case _ => 0
        }.toList
      ).toList

    )

    for (i <- 1 to 100) {
      g = g.stepCorners()
    }
    println(g.numOn())


  }

  part1()
  part2()

}