import day18.Grid
import org.scalatest.FunSuite

import scala.io.Source


class GridTest extends FunSuite {

  val filename = "resources/testInput"
  val lines = Source.fromFile(filename).getLines().toList

  var g = new Grid(

    lines.iterator.map(s =>
      s.toCharArray.map{
        case '#' => 1
        case _ => 0
      }.toList
    ).toList

  )

  test("Test Grid construction") {

    assert(g.numOn() == 15)

  }

  test("Test Grid.neighboursOf") {

    assert(g.neighboursOf((0,0)).length == 3)
    assert(g.neighboursOf((0,1)).length == 5)
    assert(g.neighboursOf((1,0)).length == 5)
    assert(g.neighboursOf((1,1)).length == 8)
    assert(g.neighboursOf((0,5)).length == 3)
    assert(g.neighboursOf((5,5)).length == 3)
    assert(g.neighboursOf((5,0)).length == 3)

  }

  test("Test Grid.neighboursOn") {

    assert(g.neighboursOn((0,0)) == 1)
    assert(g.neighboursOn((0,1)) == 0)
    assert(g.neighboursOn((1,0)) == 2)
    assert(g.neighboursOn((1,1)) == 2)
    assert(g.neighboursOn((0,5)) == 1)
    assert(g.neighboursOn((5,5)) == 1)
    assert(g.neighboursOn((5,0)) == 2)

  }

  test("Test Grid single step") {

    val g1 = g.step()
    assert(g1.numOn() == 11)
    val g2 = g1.step()
    assert(g2.numOn() == 8)
    val g3 = g2.step()
    assert(g3.numOn() == 4)
    val g4 = g3.step()
    assert(g4.numOn() == 4)

  }



}
