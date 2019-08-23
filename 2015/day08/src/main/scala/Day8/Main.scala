package Day8

import scala.io.Source

object Main extends App {

  val filename = "resources/input"

  /*for (line <- Source.fromFile(filename).getLines()) {

    println(line)

  }*/


  def eval(s: String): String = {

    s.slice(1, s.length() - 1)
      .replaceAll("\\\\\\\\", "\\\\")
      .replaceAll("\\\\\\\"", "\\\"")
      .replaceAll("\\\\x[a-f0-9][a-f0-9]?", "*") //just replace with placeholder for count

  }

  def escape(s: String): String = {

    "\"" + s.replaceAll("\\\\", "\\\\\\\\")
      .replaceAll("\\\"", "\\\\\\\"") + "\""

  }

  def part1() {

    val diffs = for (line <- Source.fromFile(filename).getLines()) yield {
      line.length() - eval(line).length()
    }
    println(diffs.foldLeft(0)(_+_))

  }

  def part2() {

    val diffs = for (line <- Source.fromFile(filename).getLines()) yield {
      escape(line).length() - line.length()
    }
    println(diffs.foldLeft(0)(_+_))

  }

  part1()
  part2()


}
