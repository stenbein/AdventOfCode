class Gift(val l: Int, w: Int, h: Int ) {

  def wrapSize(): Int = {
    val sides = List((l*w), (w*h), (h*l))
    return 2 * sides.sum + sides.min
  }

  def ribbonSize(): Int = {
    val sides = List(l, w, h)
    return sides.map(x => x*2).sum - (sides.max * 2) + sides.product
  }

}
