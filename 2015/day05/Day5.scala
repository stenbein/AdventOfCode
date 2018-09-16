import scala.io.Source

object Day5 extends App {

  val filename = "resources/input"

  def hasDoubleChar(s: String): Boolean = {

    var i = 1
    while ({ i < s.length }) {
      if (s.charAt(i-1) == s.charAt(i)) {return true}
      i += 1
    }
    false

  }

  def pairFilter(s: String): Boolean = {

    var i = 0; var j = 0
    while ({ i < s.length - 3 }) {
      var subSt = s.substring(i, i+2)
      j = i + 2
      while ({ j < s.length - 1 }) {
        if (s.substring(j, j+2) == subSt) {return true}
        j += 1
      }
      i += 1
    }
    false


  }

  def hasXyX(s: String): Boolean = {
    var i = 2
    while ({ i < s.length }) {
      if ((s.charAt(i-2) == s.charAt(i)) && (s.charAt(i-1) != s.charAt(i))) {return true}
      i += 1
    }
    false
  }

  def part1() {

    val filterKeys = Set("ab", "cd", "pq", "xy")
    val sf = new SetFilter(filterKeys)

    val vf = new SetFilter(Set("a", "e", "i", "o", "u"))

    println(Source.fromFile(filename)
      .getLines()
      .filterNot(p => sf.anyIn(p))
      .filter(p => vf.countIn(p) >= 3)
      .filter(p => hasDoubleChar(p))
      .length
    )

  }

  def part2() {

    println(Source.fromFile(filename)
      .getLines()
      .filter(p => pairFilter(p))
      .filter(p => hasXyX(p))
      .length
    )

  }

  part1()
  part2()

}
