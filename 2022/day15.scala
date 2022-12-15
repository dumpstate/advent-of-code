import scala.collection.mutable
import scala.io.Source

def sensorsFromInput(input: String) = Source.fromFile(input).getLines
    .map(line =>
        val split = line.split(": ")
        val sensor = split(0).substring(12).split(", y=").map(_.toInt)
        val beacon = split(1).substring(23).split(", y=").map(_.toInt)
        ((sensor(0), sensor(1)), (beacon(0), beacon(1)))).toList

def dist(a: (Int, Int), b: (Int, Int)) = Math.abs(a._1 - b._1) + Math.abs(a._2 - b._2)

def freq(pos: (Int, Int)) = pos._1.toLong * 4000000 + pos._2

def iterPerimeter(s: (Int, Int), r: Int) = (0 to r).iterator
    .flatMap(off => Seq(
        (s._1 + r - off + 1, s._2 + off),
        (s._1 + r - off + 1, s._2 + off),
        (s._1 - r + off - 1, s._2 + off),
        (s._1 - r + off - 1, s._2 + off)))

class SensorMap(sensors: Map[(Int, Int), Char], radiuses: Map[(Int, Int), Int]):
    val xmin = radiuses.map((pos, r) => pos._1 - r).min
    val xmax = radiuses.map((pos, r) => pos._1 + r).max

    def isCovered(p: (Int, Int)) = radiuses.find((pos, r) => dist(pos, p) <= r).isDefined
    def isEmpty(p: (Int, Int)) = sensors.get(p).isEmpty
    def countForRow(y: Int) = (xmin to xmax).filter(x => isEmpty((x, y)) && isCovered((x, y))).size

    def find(limit: Int) = radiuses.iterator.flatMap(iterPerimeter)
        .filter { case (x, y) => x >= 0 && y >= 0 && x <= limit && y <= limit }
        .find(p => isEmpty(p) && !isCovered(p)).get

object SensorMap:
    def from(sensors: List[((Int, Int), (Int, Int))]): SensorMap =
        val map = mutable.HashMap[(Int, Int), Char]()
        val radiuses = mutable.HashMap[(Int, Int), Int]()

        for (sensor, beacon) <- sensors do
            map.put(sensor, 'S')
            map.put(beacon, 'B')
            radiuses.put(sensor, dist(sensor, beacon))

        new SensorMap(map.toMap, radiuses.toMap)

def partI(sensors: SensorMap, y: Int) = sensors.countForRow(y)
def partII(sensors: SensorMap, limit: Int) = freq(sensors.find(limit))

@main def main(input: String) =
    val sensors = SensorMap.from(sensorsFromInput(input))

    println(s"Part I: ${partI(sensors, 2000000)}")
    println(s"Part II: ${partII(sensors, 4000000)}")
