import scala.io.Source
import scala.math.Ordering

def elvesFromInput(input: String): List[Int] =
    Source.fromFile(input)
        .getLines
        .mkString(",")
        .split(",,")
        .map(_.split(",").map(_.toInt).sum)
        .toList

def partI(elves: List[Int]) = elves.max

def partII(elves: List[Int]) = elves
    .sorted(Ordering.Int.reverse)
    .take(3)
    .sum

@main def main(input: String) =
    val elves = elvesFromInput(input)

    println(s"Part I: ${partI(elves)}")
    println(s"Part II: ${partII(elves)}")
