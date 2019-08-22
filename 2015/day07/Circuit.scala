package Day7

class Circuit(val wires: scala.collection.mutable.Map[String, String]) {

  // regex patterns
  val init = raw"(\w+) -> (\w)+".r //123 -> x //init
  val not = raw"NOT (\w+) -> (\w)+".r //NOT ac -> ad
  val or = raw"(\w+) OR (\w+) -> (\w)+".r //hz OR ik -> il
  val and = raw"(\w+) AND (\w+) -> (\w)+".r //jx AND jz -> ka
  val rshift = raw"(\w+) RSHIFT (\w+) -> (\w)+".r //fo RSHIFT 5 -> fr
  val lshift = raw"(\w+) LSHIFT (\w+) -> (\w)+".r //kf LSHIFT 15 -> kj

  def get_val(w: String): Int = {

    //base case input is int
    if (w forall Character.isDigit) {
      return w.toInt
    }

    //grab what we currently have for the wire
    val wire = wires(w)

    //is a digit already, return that
    if (wire forall Character.isDigit) {
      return wire.toInt
    }

    //else eval the expression, store it, and return it
    wire match {

      case init(in_val, discard) => wires(w) = get_val(in_val).toString()
      case not(in_sig, discard) => wires(w) = (65535 - get_val(in_sig)).toString()
      case or(in_sig_1, in_sig_2, discard) => wires(w) = (get_val(in_sig_1) | get_val(in_sig_2)).toString()
      case and(in_sig_1, in_sig_2, discard) => wires(w) = (get_val(in_sig_1) & get_val(in_sig_2)).toString()
      case rshift(in_sig_1, in_sig_2, discard) => wires(w) = (get_val(in_sig_1) >> get_val(in_sig_2)).toString()
      case lshift(in_sig_1, in_sig_2, discard) => wires(w) = (get_val(in_sig_1) << get_val(in_sig_2)).toString()

    }

    wires(w).toInt

  }

}
