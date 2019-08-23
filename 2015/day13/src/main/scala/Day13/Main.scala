package Day13

import scala.io.Source


object Main extends App {

  val filename = "resources/input"
  //val filename = "resources/test_input"

  val gain = raw"(\w+) would gain (\d+) happiness units by sitting next to (\w+).".r
  val lose = raw"(\w+) would lose (\d+) happiness units by sitting next to (\w+).".r

  val m = scala.collection.mutable.Map[(String, String), Int]()

  for (line <- Source.fromFile(filename).getLines()) {

    line match {

      case gain(name, units, name2) =>
        m += (name, name2) -> units.toInt
      case lose(name, units, name2) =>
        m += (name, name2) -> -units.toInt
      case _ => println("Nothing here")

    }

  }

  def deltaHappiness(l: List[String]): Int = {

    //the final add is the first in the list sitting next to the last in the list
    l.sliding(2).foldLeft(0)((i, ls) =>
      i
        + m((ls.head, ls.last))
        + m((ls.last, ls.head))
    ) + m(l.head, l.last) + m(l.last, l.head)

  }

  def part1() {

    //println(m)
    val persons: List[String] = m.keys.map(_._1).toList.distinct

    println(persons.permutations.map(deltaHappiness).max)

  }

  def part2() {

    val persons: List[String] = m.keys.map(_._1).toList.distinct ::: List("Me")

    persons.foreach(s => {m += (s, "Me") -> 0; m += ("Me", s) -> 0})

    println(persons.permutations.map(deltaHappiness).max)

  }

  part1()
  part2()

}