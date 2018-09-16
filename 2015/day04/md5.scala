import java.security.MessageDigest

object md5 {

  def iterator(k: String, p: String, i: Int): (String, Int) = {

    val hash: String = md5.hash(k + i)
    hash indexOf p match {
      case 0 => return (hash, i)
      case -1 => return iterator(k, p, i+1)
      case _ => {
        iterator(k, p, i+1)
      }
    }

  }


  //https://stevenwilliamalexander.wordpress.com/2012/06/11/scala-md5-hash-function-for-scala-console/
  def hash(text: String) : String = {
    java.security.MessageDigest
      .getInstance("MD5")
      .digest(text.getBytes())
      .map(0xFF & _)
      .map { "%02x".format(_) }
      .foldLeft(""){_ + _}
  }

}
