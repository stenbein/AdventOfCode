package Day7

import scala.io.Source

object Main extends App {

  val filename = "resources/input"

  val wires = scala.collection.mutable.Map[String, String]()
  val c = new Circuit(wires)

  //map idents to commands
  val ident = raw"([ \w]+) -> (\w+)".r

  //Each wire can only get a signal from one source
  for (line <- Source.fromFile(filename).getLines()) {

    line match {

      case ident(cmd, wire_id) => wires(wire_id) = line

    }

  }

  def part1() {

    println(c.get_val("a"))

  }

  def part2() {

    val b = c.get_val("a").toString()
    //reset wires, just pull down new copy of commands
    wires.empty
    for (line <- Source.fromFile(filename).getLines()) {

      line match {

        case ident(cmd, wire_id) => wires(wire_id) = line

      }

    }
    wires("b") = b
    println(c.get_val("a"))

  }

  part1()
  part2()

}
