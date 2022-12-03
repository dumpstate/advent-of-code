import scala.io.Source

def rucksacksFromInput(input: String) =
    Source.fromFile(input)
        .getLines
        .map(l => (
            l.substring(0, l.length / 2),
            l.substring(l.length / 2, l.length)
        ))
        .map((f, s) => (f.split("").toSet, s.split("").toSet))
        .toList

def encode(c: Char) =
    if c.toInt >= 97 then
        c.toInt - 96
    else
        c.toInt - 64 + 26

def partI(rucksacks: List[(Set[String], Set[String])]) =
    rucksacks
        .map((f, s) => f intersect s)
        .flatMap(_.map(_.charAt(0)))
        .map(encode)
        .sum

def partII(rucksacks: List[(Set[String], Set[String])]) =
    val grouped = rucksacks
        .map((f, s) => f union s)

    var groups = List.empty[Char]
    var taken = Set.empty[Int]

    for
        i <- 0 to grouped.length - 1
        j <- i + 1 to grouped.length - 1
        k <- j + 1 to grouped.length - 1
        if !taken.contains(i)
        if !taken.contains(j)
        if !taken.contains(k)
    do
        val intersection = grouped(i) intersect grouped(j) intersect grouped(k)

        if intersection.size == 1 then
            taken ++= Seq(i, j, k)
            groups = intersection.head.head :: groups

    groups.map(encode).sum

@main def main(input: String) =
    val rucksacks = rucksacksFromInput(input)

    println(s"Part I: ${partI(rucksacks)}")
    println(s"Part II: ${partII(rucksacks)}")
