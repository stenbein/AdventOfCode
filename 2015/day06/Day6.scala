package day6
import scala.io.Source

object Day6 extends App {

  val filename = "resources/input"

  //699635 to high

  //542387 to low

  def part1() {

    val g = new OnOffGrid(1000,1000)

    //turn off 232,962 through 893,979
    val off = raw"turn off (\d+),(\d+) through (\d+),(\d+)".r
    //turn on 874,567 through 943,684
    val on = raw"turn on (\d+),(\d+) through (\d+),(\d+)".r
    //toggle 911,840 through 990,932
    val tog = raw"toggle (\d+),(\d+) through (\d+),(\d+)".r
    for (line <- Source.fromFile(filename).getLines()) {

      line match {

        case off(one1, one2, two1, two2) => g.toggleOff((one1.toInt,one2.toInt), (two1.toInt,two2.toInt))
        case on(one1, one2, two1, two2) => g.toggleOn((one1.toInt,one2.toInt), (two1.toInt,two2.toInt))
        case tog(one1, one2, two1, two2) => g.toggle((one1.toInt,one2.toInt), (two1.toInt,two2.toInt))

      }

    }

    println(g.numbOn())

  }

  //we don't need to muck about with all that grid nonsense here.
  //Just sum the values as we go...
  def part2() {

    val g = new BrightnessGrid(1000,1000)

    //turn off 232,962 through 893,979
    val off = raw"turn off (\d+),(\d+) through (\d+),(\d+)".r
    //turn on 874,567 through 943,684
    val on = raw"turn on (\d+),(\d+) through (\d+),(\d+)".r
    //toggle 911,840 through 990,932
    val tog = raw"toggle (\d+),(\d+) through (\d+),(\d+)".r
    for (line <- Source.fromFile(filename).getLines()) {

      line match {

        case off(one1, one2, two1, two2) => g.toggleOff((one1.toInt,one2.toInt), (two1.toInt,two2.toInt))
        case on(one1, one2, two1, two2) => g.toggleOn((one1.toInt,one2.toInt), (two1.toInt,two2.toInt))
        case tog(one1, one2, two1, two2) => g.toggle((one1.toInt,one2.toInt), (two1.toInt,two2.toInt))

      }

    }

    println(g.numbOn())

  }

  part1()
  part2()

}
