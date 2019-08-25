package day18


class Grid(val l: List[List[Int]]) {

  val transformations = (for (i <- -1 to 1; j <- -1 to 1) yield (i, j)).filter(_ != (0, 0))
  val coords = (for (i <- 0 to l.length - 1) yield (for (j <- 0 to l.length - 1) yield (i, j)).toList).toList

  def onGrid(coord: (Int, Int)): Boolean = {
    coord._1 >= 0 && coord._1 < this.l.length &&
      coord._2 >= 0 && coord._2 < this.l.length
  }

  def neighboursOf(coord: (Int, Int)): List[(Int, Int)] = {

    transformations
      .map(c => (c._1 + coord._1, c._2 + coord._2))
      .filter(this.onGrid).toList

  }

  def neighboursOn(coord: (Int, Int)): Int = {

    neighboursOf(coord).foldLeft(0) { case (b, (i, j)) => b + this.l(i)(j) }

  }

  def step(): Grid = {

    new Grid(coords.map(row =>
      row.map(c =>
        if (this.l(c._1)(c._2) == 1) {
          if (neighboursOn(c) == 2 || neighboursOn(c) == 3) {
            1
          } else {
            0
          }
        } else {
          if (neighboursOn(c) == 3) {
            1
          } else {
            0
          }
        }
      )
    ))

  }

  def stepCorners(): Grid = {

    new Grid(coords.map(row =>
      row.map(c =>
        if (c == (0,0) || c == (0, l.length-1) || c == (l.length-1, 0) || c == (l.length-1, l.length-1)) {1
        } else if (this.l(c._1)(c._2) == 1) {
          if (neighboursOn(c) == 2 || neighboursOn(c) == 3) {
            1
          } else {
            0
          }
        } else {
          if (neighboursOn(c) == 3) {
            1
          } else {
            0
          }
        }
      )
    ))

  }

  def numOn(): Int = {
    l.flatten.sum
  }

}