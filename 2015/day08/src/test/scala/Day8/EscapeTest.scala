package Day8


import org.scalatest.FunSuite

class EscapeTest extends FunSuite {

  val s0 = "\"\""
  val s1 = "\"abc\""
  val s2 = "\"aaa\\\"aaa\""
  val s3 = "\"\\x27\""

  test("Escape init") {
    assert(Main.eval(s0).length() === 0)
    assert(Main.eval(s1).length() === 3)
    assert(Main.eval(s2).length() === 7)
    assert(Main.eval(s3).length() === 1)
  }

  test("Escape lengths") {
    assert(s0.length() === 2)
    assert(s1.length() === 5)
    assert(s2.length() === 10)
    assert(s3.length() === 6)
  }

  test("Escape totals") {
    assert((s0.length() + s1.length() + s2.length() + s3.length()) === 23)
    assert((Main.eval(s0).length() +
    Main.eval(s1).length() +
    Main.eval(s2).length() +
    Main.eval(s3).length()) === 11)
  }

  test("Escape sums") {

    val lengths = (s0.length() +
      s1.length() +
      s2.length() +
      s3.length())

    val lengths2 = (Main.eval(s0).length() +
      Main.eval(s1).length() +
      Main.eval(s2).length() +
      Main.eval(s3).length())

    assert((lengths - lengths2) === 12)


  }

}
