class SetFilter(val a: Set[String]) {

  def anyIn(t: String): Boolean = {
    for (elem <- a) {
      if (t.contains(elem)) {return true}
    }
    false
  }

  def noneIn(t: String): Boolean = {
    for (elem <- a) {
      if (t.contains(elem)) {return false}
    }
    true
  }

  def countIn(t: String): Int = {

    var c: Int = 0
    for (elem <- a) {
      c += elem.r.findAllMatchIn(t).length
    }

    return c
  }

}
