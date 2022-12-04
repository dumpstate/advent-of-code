import scala.io.Source

case class Range(from: Int, to: Int):

    def contains(range: Range) = range.from >= from && range.to <= to

    def overlaps(range: Range) =
        (range.to <= to && range.to >= from) ||
        (range.from >= from && range.from <= to)

def rangesFromInput(input: String) =
    Source.fromFile(input)
        .getLines
        .map(_.split(",").map(_.split("-").map(_.toInt) match
            case Array(l, r) => Range(l, r)) match
            case Array(l, r) => (l, r))
        .toList

def partI(ranges: List[(Range, Range)]) = ranges
    .filter((l, r) => l.contains(r) || r.contains(l))
    .length

def partII(ranges: List[(Range, Range)]) = ranges
    .filter((l, r) => l.overlaps(r) || r.overlaps(l))
    .length

@main def main(input: String) =
    val ranges = rangesFromInput(input)

    println(s"Part I: ${partI(ranges)}")
    println(s"Part II: ${partII(ranges)}")
