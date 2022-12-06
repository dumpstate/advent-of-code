import scala.io.Source

def unique(marker: String) =
    Set.from(marker.split("")).size == marker.length

def findPacket(buffer: String, size: Int) =
    buffer.sliding(size)
        .zipWithIndex
        .find((marker, _) => unique(marker))
        .map(_._2).get + size

def partI(buffer: String) = findPacket(buffer, 4)

def partII(buffer: String) = findPacket(buffer, 14)

@main def main(input: String) =
    val buffer = Source.fromFile(input).getLines().next

    println(s"Part I: ${partI(buffer)}")
    println(s"Part II: ${partII(buffer)}")
