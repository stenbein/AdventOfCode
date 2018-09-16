import scala.collection.mutable
import scala.collection.mutable.HashMap

class Neighbourhood {

  val start: House = new House((0,0), 1)
  val houses: HashMap[(Int, Int), House] = new mutable.HashMap

  houses.put(start.location, start)

  class House(val location: (Int, Int), val visits: Int) {

    def toNext(direction: Char): House = {

      val nextLoc: (Int, Int) = direction match {

        case '>' => (location._1 + 1, location._2)
        case '<' => (location._1 - 1, location._2)
        case '^' => (location._1, location._2 + 1)
        case 'v' => (location._1, location._2 - 1)

      }

      val house = houses.get(nextLoc) match {
        case Some(i) => new House(nextLoc, Some(i).value.visits + 1) //add 1 to visits for new house
        case None => new House(nextLoc, 1)
      }

      houses.update(nextLoc, house)
      house

    }

  }

  def houseAt(location: (Int, Int)): House = {
    houses.get(location) match {
      case Some(i) => Some(i).value
      case None => new House(location, 1)
    }
  }

  //def next(current: House, direction: Char): House = {
  //  current.toNext(direction)
  //}

  def visitedOnce(): Int = {
    //println(houses.keySet.seq)
    houses.size
  }

  //mistakenly read that we were attempting to find how many houses got more than one present
  /*def visitedTwice(): Int = {
    houses.count(v => v._2.visits > 1)
  }*/

}
