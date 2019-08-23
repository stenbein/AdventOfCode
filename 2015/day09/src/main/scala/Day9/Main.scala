package Day9

import scala.io.Source

object Main extends App {

  val filename = "resources/input"

  val dist_regex = raw"(\w+) to (\w+) = (\d+)".r
  val distances = scala.collection.mutable.Map[(String, String), Int]()

  for (line <- Source.fromFile(filename).getLines()) {

    line match {

      case dist_regex(from, to, distance) => {
        distances((from, to)) = distance.toInt
        distances((to, from)) = distance.toInt
      }
      case _ => throw new Exception("NO MATCH")

    }

  }

  //what is this wizardly bullshit?
  val locations = (for (d <- distances.seq) yield {d._1._1}).toList.distinct

  def get_distance(loc: List[String]): Int = {

    loc match {
      case head::Nil => 0
      case head::tail => distances((head, tail.head)) + get_distance(tail)
    }

  }

  //we can start at any point, and then move to every other point
  //traveling salesman problem so there is no magic shortcut except to check
  //and use heuristics or time saving approaches


  def part1() {

    var min_dist = get_distance(locations)
    for (p <- locations.permutations) {
      val dist = get_distance(p)
      if (dist < min_dist) {min_dist = dist}
    }

    println(min_dist)
  }

  def part2() {

    var max_dist = get_distance(locations)
    for (p <- locations.permutations) {
      val dist = get_distance(p)
      if (dist > max_dist) {max_dist = dist}
    }

    println(max_dist)
  }

  part1()
  part2()

}