package Day10

import scala.annotation.tailrec
import scala.io.Source

object Main extends App {

  /*val filename = "resources/input"

  for (line <- Source.fromFile(filename).getLines()) {



  }
  */

  val input = "1113122113"

  /*def look(s: String): List[(Char, Int)]  = s.foldLeft(List.empty[(Char, Int)]) {
    case ((chr, count) :: tail, c) if chr == c => (chr, count +1) :: tail
    case (xs, c) => (c, 1) :: xs
  }*/

  @tailrec
  def countString(lc: List[Char], counts: Int, acc: List[(Char, Int)]): List[(Char, Int)] = {

    lc match {
      case Nil => acc
      case head::Nil => (head, counts) :: acc
      case head :: tail => {
        if (head == tail.head) {countString(tail, counts + 1, acc)}
        else {countString(tail, 1, (head, counts) :: acc)}
      }
    }

  }

  def countsToString(counts: List[(Char, Int)]): String = {
    counts.reverse.map{case (c: Char, i: Int) => s"$i$c"}.mkString
  }

  def lookAndSay(s: String) : String = {

    //split into counts
    val counts = countString(s.toCharArray.toList, 1, List.empty[(Char, Int)])
    //rejoin counts into output
    countsToString(counts)

  }



  def part1() {

    var s = input
    for (i <- 1 to 40) {

      s = lookAndSay(s)

    }
    println(s.length)

  }

  def part2() {

    var s = input
    for (i <- 1 to 50) {

      s = lookAndSay(s)

    }
    println(s.length)

  }

  part1()
  part2()

}


