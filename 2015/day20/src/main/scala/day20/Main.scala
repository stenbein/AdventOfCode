package day20

import scala.math.sqrt
import scala.util.control.Breaks._

object Main extends App {

  val input: Int = 29000000

  def part1() {

    var houses: Array[Int] = (1 to input).map(_ => 0).toArray
    for (elf <- 1 to input) {

      for (i <- elf to input by elf) {
        houses(i-1) += (elf * 10)
      }
      if (houses(elf - 1) >= input) {
        println(elf)
        return
      }

    }

  }

  def part2() {

    var houses: Array[Int] = (1 to input).map(_ => 0).toArray
    for (elf <- 1 to input) {
      var count: Int = 0
      for (i <- elf to input by elf) {
        if (count < 50) houses(i-1) += (elf * 11)
        count += 1
      }
      if (houses(elf - 1) >= input) {
        println(elf)
        return
      }

    }

  }

  part1()
  part2()

}
