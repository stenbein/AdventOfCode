import scala.io.Source

object Day3 extends App {

  val filename = "resources/input"

  def part1() {

    val town: Neighbourhood = new Neighbourhood

    var house: town.House = town.houseAt(location = (0,0))

    for (char <- Source.fromFile(filename)) {

      house = house.toNext(char)

    }

    println(town.visitedOnce())

  }

  def part2() {

    val town: Neighbourhood = new Neighbourhood

    var santaLoc: town.House = town.houseAt(location = (0,0))
    var roboLoc: town.House = town.houseAt(location = (0,0))

    var toggle: Boolean = true
    for (char <- Source.fromFile(filename)) {

      if (toggle) {
        santaLoc = santaLoc.toNext(char)
        //println("Santa loc:", santaLoc.location.toString())
        toggle = false
      } else {
        roboLoc = roboLoc.toNext(char)
        //println("Robo loc:", roboLoc.location.toString())
        toggle = true
      }

    }

    println(town.visitedOnce())

  }


  part1()
  part2()

}
