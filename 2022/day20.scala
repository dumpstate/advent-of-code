import scala.io.Source

def rm[T](v: Vector[T], ix: Int) = v.slice(0, ix) ++ v.slice(ix + 1, v.length)

def put(v: Vector[(Long, Int)], n: Long, ix: Int) =
    val nix = ix match
        case _ if ix > 0 => ix % v.length
        case _ if ix == 0 => v.length
        case _ if ix < 0 => (v.length + ix) % v.length
    v.slice(0, nix) ++ Vector((n, ix)) ++ v.slice(nix, v.length)

def mix(ns: Vector[Long], loops: Int = 1) =
    val data = ns.zipWithIndex
    var res = data

    for
        _ <- 1 to loops
        (n, oIx) <- data
    do
        val ix = res.indexOf((n, oIx))
        val without = rm(res, ix)
        res = put(without, n, (ix + n).toInt)

    res

def coords(ns: Vector[(Long, Int)]) =
    val off = ns.zipWithIndex.find(_._1._1 == 0).get._2
    List(1000, 2000, 3000).map(n => (n + off) % ns.length).map(ns(_)._1).sum

def partI(ns: Vector[Long]) = coords(mix(ns))

def partII(ns: Vector[Long]) = coords(mix(ns.map(_ * 811589153), 10))

@main def main(input: String) =
    val ns = Source.fromFile(input).getLines.map(_.toLong).toVector

    println(s"Part I: ${partI(ns)}")
    println(s"Part II: ${partII(ns)}")
