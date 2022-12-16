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

def maxFlow(layout: Layout, cache: mutable.HashMap[(String, Set[String], Int), Int])(
    current: String,
    open: Set[String],
    timeLeft: Int,
): Int =
    val key = (current, open, timeLeft)
    val res = cache.getOrElse(key, {
        if timeLeft <= 0 then 0
        else
            if !open.contains(current) then
                val ps = (timeLeft - 1) * layout.flow(current)
                if ps != 0 then
                    layout.neighbours(current)
                        .map(n => Math.max(
                            ps + maxFlow(layout, cache)(n, open + current, timeLeft - 2),
                            maxFlow(layout, cache)(n, open, timeLeft - 1))).max
                else
                    layout.neighbours(current)
                        .map(n => maxFlow(layout, cache)(n, open, timeLeft - 1)).max
            else 0
    })
    cache.put(key, res)
    res

def partI(layout: Layout) = maxFlow(layout, mutable.HashMap())("AA", Set(), 30)

@main def main(input: String) =
    val layout = Layout.from(valvesFromInput(input))

    println(f"Part I: ${partI(layout)}")
