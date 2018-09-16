import scala.io.Source

object Day4 extends App {

  val filename = "resources/input"

  def part1() {

    val keyWord = Source.fromFile(filename).getLines.toSeq(0).toString()
    println(md5.iterator(keyWord, "00000", 1))

  }

  def part2() {

    val keyWord = Source.fromFile(filename).getLines.toSeq(0).toString()
    println(md5.iterator(keyWord, "000000", 1))

  }

  part1()
  part2()

}
