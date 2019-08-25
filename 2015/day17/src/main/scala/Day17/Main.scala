package Day17

object Main extends App {

  def countExact(l: List[Int], amount: Int, limit: Int): Int = {

    if (limit == 0) return 0
    l match {
      case Nil => 0
      case x :: xs if x == amount => 1 + countExact(xs, amount, limit)
      case x :: xs if x < amount =>
        countExact(xs, amount - x, limit - 1) + countExact(xs, amount, limit)
      case x :: xs => countExact(xs, amount, limit)
    }
  }

  val containers = List(33, 14, 18, 20, 45
    , 35, 16, 35, 1, 13, 18, 13, 50, 44
    , 48, 6, 24, 41, 30, 42)

  def part1() {

    println(countExact(containers, 150, containers.length))

  }

  def part2() {

    def counts(a: Int): LazyList[(Int, Int)] = (a, countExact(containers, 150, a)) #:: counts(a+1)

    println(counts(0).filter(p => p._2 > 0).head._2)

  }

  part1()
  part2()

}
