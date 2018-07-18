import scala.io.Source

object Day2 extends App {

  val filename = "resources/input"

  def makeGift(s: String): Gift = {
    val dimensions = (s.split("x").map(Integer.parseInt(_)))
    new Gift(dimensions(0), dimensions(1), dimensions(2))
  }

  def part1() {

    var total = 0

    for (line <- Source.fromFile(filename).getLines()) {
      total += makeGift(line).wrapSize()
    }

    println(total)

  }

  def part2() {

    var total = 0

    for (line <- Source.fromFile(filename).getLines()) {
      total += makeGift(line).ribbonSize()
    }

    println(total)

  }

  part1()
  part2()

}
