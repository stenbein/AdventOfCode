package Day14

import scala.io.Source
import scala.util.Using


object Main extends App {

  val filename = "resources/input"
  val race_duration = 2503
  val parse = raw"(\w+) can fly (\d+) km/s for (\d+) seconds, but then must rest for (\d+) seconds.".r

  val reindeer = Using(Source.fromFile(filename)){
    file => file.getLines().toList.map({
      case parse(name, velocity, duration, rest) => new Reindeer(name, velocity.toInt, duration.toInt, rest.toInt)
    })
  }.get


  def part1() {

    println(reindeer.map(r => r.distanceAt(race_duration)).max)

  }

  def part2() {

    //If there are multiple reindeer tied for the lead, they each get one point.

    var currentMax: Int = 0
    for (i <- 1 to race_duration) {
      currentMax = reindeer.map(r => r.distanceAt(i)).max
      reindeer.filter(r => r.distanceAt(i) == currentMax).foreach(_.points += 1)
    }
    println(reindeer.map(r => r.points).max)

  }

  part1()
  part2()

}