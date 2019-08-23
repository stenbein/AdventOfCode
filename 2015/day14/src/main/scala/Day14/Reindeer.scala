package Day14

class Reindeer(val name: String, velocity: Int, run: Int, rest: Int) {

  var points: Int = 0

  def distanceAt(seconds: Int): Int = {

    val cycles: Int = seconds / (run + rest)
    val extra: Int = seconds % (run + rest)

    cycles * velocity * run + (
      if(extra > run) run * velocity
      else extra * velocity
    )

  }

}
