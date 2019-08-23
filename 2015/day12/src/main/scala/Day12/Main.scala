package Day12

import scala.io.Source
import scala.util.parsing.json.JSON

object Main extends App {

  val filename = "resources/input"

  def traverse(input: Any): Int = {

    input match {
      case m: Map[Any, Any] => m.foldLeft(0) {
        case (i, (_, z)) => i + traverse(z)
      }
      case l: List[Any] => l.foldLeft(0)((i, li) => i + traverse(li))
      case i: Double => i.toInt
      case _ => 0
    }

  }

  def traverse_without_red(input: Any): Int = {

    input match {
      case m: Map[Any, Any] => {

        if (m.forall(p => p._2 != "red")) {
          m.foldLeft(0) { case (i, (_, z)) => i + traverse_without_red(z) }
        } else {
          0
        }

      }
      case l: List[Any] => l.foldLeft(0)((i, li) => i + traverse_without_red(li))
      case i: Double => i.toInt
      case _ => 0
    }

  }

    def part1() {

      for (line <- Source.fromFile(filename).getLines()) {

        val json = JSON.parseFull(line)

        json match {
          case Some(x) => println(traverse(x))
          case None => println("Nothing to parse")
        }

      }

    }

    def part2() {

      for (line <- Source.fromFile(filename).getLines()) {

        val json = JSON.parseFull(line)

        json match {
          case Some(x) => println(traverse_without_red(x))
          case None => println("Nothing to parse")
        }

      }

    }

    part1()
    part2()

  }