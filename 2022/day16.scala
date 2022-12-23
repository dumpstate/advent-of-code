import scala.collection.mutable
import scala.io.Source

type Layout = Map[String, (Int, List[String])]

def layoutFromInput(input: String) = Source.fromFile(input).getLines
    .map(line =>
        val split = line.split("; ")
        val valve = split(0).substring(6, 8)
        val flowRate = split(0).substring(23).toInt
        val valves = split(1).substring(22).trim.split(", ")
        (valve -> (flowRate, valves.toList))).toMap

def distance(layout: Layout, from: String, to: String): Int =
    val visited = mutable.HashSet[String]()
    var queue = mutable.Queue[String](from)
    var dist = 0

    while queue.nonEmpty do
        val nextQueue = mutable.Queue[String]()

        for valve <- queue do
            if !visited.contains(valve) then
                visited.add(valve)

                if valve == to then
                    return dist
                else
                    nextQueue.enqueueAll(layout(valve)._2)

        queue = nextQueue
        dist += 1

    throw new Exception("path does not exist")

def next(layout: Layout, nonEmptyValves: Layout, current: String, timeLeft: Int): (String, Int) =
    nonEmptyValves.flatMap { case (valve, (flowRate, _)) =>
        val remNonEmpty = nonEmptyValves.removed(valve)
        val remTime = timeLeft - distance(layout, current, valve) - 1
        if remTime <= 0 then None
        else Some((valve, remTime * flowRate + next(layout, remNonEmpty, valve, remTime)._2))
    }.maxByOption(_._2).getOrElse((current, 0))

def nonEmpty(layout: Layout) = layout.filter { case (_, (flowRate, _)) => flowRate > 0 }

def iterPairs[T](set: Set[T]) =
    val items = set.toVector

    for (
        a <- (0 until items.length).iterator;
        b <- (a + 1 until items.length).iterator
    ) yield (items(a), items(b))

def score(layout: Layout, set: Set[String]) = iterPairs(set).map(p => distance(layout, p._1, p._2)).sum
def score(layout: Layout, l: Set[String], r: Set[String]): Int = score(layout, l) + score(layout, r)

def split(layout: Layout) =
    val valves = nonEmpty(layout).keySet
    valves.subsets
        .map(l => List(l, valves.removedAll(l)))
        .map(st => (st, score(layout, st(0), st(1))))
        .toVector.sortBy(_._2).head._1

def maxFlow(layout: Layout, nonEmpty: Layout, start: String, timeLeft: Int) = next(layout, nonEmpty, start, timeLeft)._2

def pick(layout: Layout, keys: Set[String]) = layout.removedAll(layout.keySet -- keys)

def partI(layout: Layout) = maxFlow(layout, nonEmpty(layout), "AA", 30)
def partII(layout: Layout) = split(layout)
    .map(pick(layout, _))
    .map(maxFlow(layout, _, "AA", 26))
    .sum

@main def main(input: String) =
    val layout = layoutFromInput(input)

    println(s"Part I: ${partI(layout)}")
    println(s"Part II: ${partII(layout)}")
