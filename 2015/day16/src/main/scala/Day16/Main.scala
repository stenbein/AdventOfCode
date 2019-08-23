package Day16

import scala.io.Source
import scala.util.Using

object Main extends App {

  val filename = "resources/input"

  val parse = raw"Sue (\d+): (\w+): (\d+), (\w+): (\d+), (\w+): (\d+)".r

  val sues: List[Sue] = Using(Source.fromFile(filename)) {
    file => file.getLines().toList
  }.get.map(s => s match {
    case parse(id, prop1, val1, prop2, val2, prop3, val3) =>
      new Sue(id.toInt, (prop1, val1.toInt) :: (prop2, val2.toInt) :: (prop3, val3.toInt) :: Nil)
  })

  class Sue(val id: Int, val l: List[(String, Int)]) {
    def has(prop: (String, Int)): Boolean = {
      l.foreach(p =>
        if (p._1 == prop._1) {
          if (p._2 == prop._2) return true
          return false
        }
      )
      return true
    }

    def hasGreater(prop: (String, Int)): Boolean = {
      l.foreach(p =>
        if (p._1 == prop._1) {
          if (p._2 > prop._2) return true
          return false
        }
      )
      return true
    }

    def hasLess(prop: (String, Int)): Boolean = {
      l.foreach(p =>
        if (p._1 == prop._1) {
          if (p._2 < prop._2) return true
          return false
        }
      )
      return true
    }
  }

  def part1() {

    sues
      .filter(s => s.has(("children", 3)))
      .filter(s => s.has(("cats", 7)))
      .filter(s => s.has(("samoyeds", 2)))
      .filter(s => s.has(("pomeranians", 3)))
      .filter(s => s.has(("akitas", 0)))
      .filter(s => s.has(("vizslas", 0)))
      .filter(s => s.has(("goldfish", 5)))
      .filter(s => s.has(("trees", 3)))
      .filter(s => s.has(("cars", 2)))
      .filter(s => s.has(("perfumes", 1)))
      .foreach(s => println(s.id))

  }

  def part2() {

    sues
      .filter(s => s.has(("children", 3)))
      .filter(s => s.hasGreater(("cats", 7)))
      .filter(s => s.has(("samoyeds", 2)))
      .filter(s => s.hasLess(("pomeranians", 3)))
      .filter(s => s.has(("akitas", 0)))
      .filter(s => s.has(("vizslas", 0)))
      .filter(s => s.hasLess(("goldfish", 5)))
      .filter(s => s.hasGreater(("trees", 3)))
      .filter(s => s.has(("cars", 2)))
      .filter(s => s.has(("perfumes", 1)))
      .foreach(s => println(s.id))

  }

  part1()
  part2()

}
