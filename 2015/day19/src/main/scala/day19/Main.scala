package day19

import scala.io.Source
//import scala.collection.concurrent.TrieMap

object Main extends App {

  val filename = "resources/input"

  val base = "ORnPBPMgArCaCaCaSiThCaCaSiThCaCaPBSiRnFArRnFArCaCaSiThCaCaSiThCaCaCaCaCaCaSiRnFYFArSiRnMgArCaSiRnPTiTiBFYPBFArSiRnCaSiRnTiRnFArSiAlArPTiBPTiRnCaSiAlArCaPTiTiBPMgYFArPTiRnFArSiRnCaCaFArRnCaFArCaSiRnSiRnMgArFYCaSiRnMgArCaCaSiThPRnFArPBCaSiRnMgArCaCaSiThCaSiRnTiMgArFArSiThSiThCaCaSiRnMgArCaCaSiRnFArTiBPTiRnCaSiAlArCaPTiRnFArPBPBCaCaSiThCaPBSiThPRnFArSiThCaSiThCaSiThCaPTiBSiRnFYFArCaCaPRnFArPBCaCaPBSiRnTiRnFArCaPRnFArSiRnCaCaCaSiThCaRnCaFArYCaSiRnFArBCaCaCaSiThFArPBFArCaSiRnFArRnCaCaCaFArSiRnFArTiRnPMgArF"
  val lines = Source.fromFile(filename).getLines().toList

  val parse = raw"(\w+) => (\w+)".r




  def part1() {

    val replacements = scala.collection.mutable.Map[String, List[String]]()

    lines.map{case parse(w1, w2) => (w1, w2)}.foreach(
      t => if (replacements.contains(t._1)) {
        replacements(t._1) = t._2 :: replacements(t._1)
      } else {
        replacements(t._1) = t._2 :: Nil
      }
    )

    val width = replacements.keys.map(_.length).max
    var strings = Set[String]()

    val chars = base.toCharArray
    for (w <- 1 to width) {
      for (i <- 0 to base.length - w) {

        val test = chars.slice(i, i+w).mkString
        if (replacements.contains(test)) {
          replacements(test).foreach( r =>
            strings += chars.patch(i, r.toSeq, w).mkString
          )
        }

      }
    }

    println(strings.size)

  }

  def part2() {

    val replacements = scala.collection.mutable.Map[String, List[String]]()

    lines.map{case parse(w1, w2) => (w2, w1)}.foreach(
      t => if (replacements.contains(t._1)) {
        replacements(t._1) = t._2 :: replacements(t._1)
      } else {
        replacements(t._1) = t._2 :: Nil
      }
    )

    val width = replacements.keys.map(_.length).max
    var strings = Set[String](base)

    var count: Int = 0

    while (!strings.contains("e")) {

      count += 1
      var temp = Set[String]()

      for (s <- strings) {
        val chars = s.toCharArray
        for (w <- 1 to width) {
          for (i <- 0 to chars.length - w) {

            val test = chars.slice(i, i + w).mkString
            if (replacements.contains(test)) {
              replacements(test).foreach(r =>
                temp += chars.patch(i, r.toSeq, w).mkString
              )
            }

          }
        }

      }

      strings = temp take 5 //wtf? I'm not 100% sure why this worked. I believe this is simulating a depth first search
      val m = strings.map(_.length).min
      println(count, strings.size, m)

    }

  }



  part1()
  part2()

}
