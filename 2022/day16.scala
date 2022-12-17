import scala.collection.mutable
import scala.io.Source

def valvesFromInput(input: String) = Source.fromFile(input).getLines
    .map(line =>
        val split = line.split("; ")
        val valve = split(0).substring(6, 8)
        val flowRate = split(0).substring(23).toInt
        val valves = split(1).substring(22).trim.split(", ")
        (valve, flowRate, valves.toList)).toList

class Layout(val flowRates: Map[String, Int], val links: Map[String, List[String]]):
    def flow(valve: String) = flowRates(valve)
    def neighbours(valve: String) = links(valve)

object Layout:
    def from(valves: List[(String, Int, List[String])]) =
        new Layout(valves.map(v => (v._1, v._2)).toMap, valves.map(v => (v._1, v._3)).toMap)

def memoized[T, R](fn: (Layout, T) => R) =
    val cache = mutable.HashMap[T, R]()
    (layout: Layout, state: T) =>
        val res = cache.getOrElse(state, fn(layout, state))
        cache.put(state, res)
        res

def maxFlow(layout: Layout, state: (String, Set[String], Int)): Int = state match
    case (_, _, timeLeft) if timeLeft <= 0 => 0
    case (current, open, _) if open.contains(current) => 0
    case (current, open, timeLeft) if layout.flow(current) == 0 =>
        layout.neighbours(current)
            .map(n => memoMaxFlow(layout, (n, open, timeLeft - 1))).max
    case (current, open, timeLeft) =>
        val ps = (timeLeft - 1) * layout.flow(current)
        layout.neighbours(current)
            .map(n => Math.max(
                ps + memoMaxFlow(layout, (n, open + current, timeLeft - 2)),
                memoMaxFlow(layout, (n, open, timeLeft - 1)))).max

val memoMaxFlow = memoized(maxFlow)

def partI(layout: Layout, time: Int) = memoMaxFlow(layout, ("AA", Set(), time))

@main def main(input: String) =
    val layout = Layout.from(valvesFromInput(input))

    println(f"Part I: ${partI(layout, 30)}")
