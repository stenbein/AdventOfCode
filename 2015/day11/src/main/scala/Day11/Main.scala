package Day11

import scala.annotation.tailrec
import scala.io.Source

object Main extends App {

  def time[R](block: => R): R = {
    val t0 = System.nanoTime()
    val result = block    // call-by-name
    val t1 = System.nanoTime()
    println("Elapsed time: " + (t1 - t0)/1000000 + "ms")
    result
  }

  val input = "cqjxjnds"

  val m = scala.collection.mutable.Map[Char, Char]()
  for ((i, j) <- ('a' to 'z').zip(('b' to 'z') :+ 'a')) {m(i) = j}

  def map_chars(cl: List[Char]): List[Char] = {

    cl match {
      case c :: tail if c == 'z' => m(c) :: map_chars(tail)
      case c :: tail => m(c) :: tail
    }

  }

  //from our input rotate chars to the next input to try
  def rotate(s: String): String = {

    map_chars(s.toCharArray.toList.reverse)
      .reverse
      .mkString

  }

  //stream passwords instead of pre-compute
  def passwords(s: String): Stream[String] = {
    def recurse(s: String): Stream[String] = rotate(s) #:: recurse(rotate(s))
    recurse(s)
  }

  def increasingStraight(s: String) = (s.head+1).toChar == s.charAt(1) && s.charAt(1) == (s.last-1).toChar

  //Passwords must include one increasing straight of at least three letters,
  // like abc, bcd, cde, and so on, up to xyz. They cannot skip letters; abd doesn't count.
  // Sliding is new to me here, neat little thing, traverses the iterator in sets of n items
  def increasingThree(s: String): Boolean = {

    s.sliding(3).exists(sl => (sl.head+1).toChar == sl.charAt(1) && (sl.charAt(1)+1).toChar == sl.charAt(2))

  }

  //Passwords may not contain the letters i, o, or l, as these letters can be mistaken for other characters and are therefore confusing.
  val excluded_chars = Set('i', 'o', 'l')
  def hasExcludedChar(s: String): Boolean = {
    s.forall(c => !excluded_chars.contains(c))
  }
  //Passwords must contain at least two different, non-overlapping pairs of letters, like aa, bb, or zz.
  def hasTwoPair(s: String): Boolean = {

    val pairs = s.foldLeft(List.empty[(Int, Char)]) {
      case ((count, ch)::tail, c) if c == ch => (count+1, ch) :: tail
      case (xs, c) => (1, c) :: xs
    }.filter({case (i, _) => i > 1}) //keep the pairs

    pairs.map(_._2).distinct.length > 1

  }

  def part1() {

    val result = passwords(input)
      .filter(s => increasingThree(s))
      .filter(s => hasExcludedChar(s))
      .filter(s => hasTwoPair(s)).head

    println(result)

  }

  def part2() {

    val part1 = passwords(input)
      .filter(s => increasingThree(s))
      .filter(s => hasExcludedChar(s))
      .filter(s => hasTwoPair(s)).head

    val result = passwords(part1)
      .filter(s => increasingThree(s))
      .filter(s => hasExcludedChar(s))
      .filter(s => hasTwoPair(s)).head

    println(result)

  }

  part1()
  part2()

}