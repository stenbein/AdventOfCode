package day6

trait Grid {val height: Int; val width: Int}

case class OnOffGrid(val height: Int, val width: Int) extends Grid {

  //boolean array of arrays
  val g = Array.fill(height)(Array.fill(width)(false))

  def toggleOn(first: (Int,Int), second: (Int,Int)) {

    for ( i <- first._1 to second._1) {
      var row = g(i)
      for ( j <- first._2 to second._2) {
        row(j) = true
      }
    }

  }

  def toggleOff(first: (Int,Int), second: (Int,Int)) {

    for ( i <- first._1 to second._1) {
      var row = g(i)
      for ( j <- first._2 to second._2) {
        row(j) = false
      }
    }

  }

  def toggle(first: (Int,Int), second: (Int,Int)) {

    for ( i <- first._1 to second._1) {
      val row = g(i)
      for ( j <- first._2 to second._2) {
        row(j) = !row(j)
      }
    }

  }

  def numbOn(): Int = {
    g.map( p => p.count(q => q == true)).sum
  }

}

case class BrightnessGrid(val height: Int, val width: Int) extends Grid {

  //boolean array of arrays
  val g = Array.fill(height)(Array.fill(width)(0))

  def toggleOn(first: (Int,Int), second: (Int,Int)) {

    for ( i <- first._1 to second._1) {
      var row = g(i)
      for ( j <- first._2 to second._2) {
        row(j) += 1
      }
    }

  }

  def toggleOff(first: (Int,Int), second: (Int,Int)) {

    for ( i <- first._1 to second._1) {
      var row = g(i)
      for ( j <- first._2 to second._2) {
        if (row(j) > 0) {row(j) -= 1}
      }
    }

  }

  def toggle(first: (Int,Int), second: (Int,Int)) {

    for ( i <- first._1 to second._1) {
      val row = g(i)
      for ( j <- first._2 to second._2) {
        row(j) += 2
      }
    }

  }

  def numbOn(): Int = {
    g.map( p => p.sum ).sum
  }

}