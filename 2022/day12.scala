import scala.collection.mutable
import scala.io.Source

case class Node(pos: (Int, Int), height: Char):
    def isStart = height == 'S'
    def isEnd = height == 'E'

def mapFromInput(input: String) =
    Source.fromFile(input)
        .getLines
        .map(_.split("").map(_.charAt(0)).toVector)
        .toVector

def elevation(char: Char) = char match
    case 'S' => 'a'
    case 'E' => 'z'
    case _ => char

def graph(map: Vector[Vector[Char]], maxDiff: Int = 1) =
    val nodes = mutable.ListBuffer[Node]()
    val edges = mutable.ListBuffer[(Node, Node)]()

    for
        j <- 0 until map.length
        i <- 0 until map(j).length
    do
        val curr = Node((i, j), map(j)(i))
        nodes.append(curr)
        List((i, j + 1), (i, j - 1), (i + 1, j), (i - 1, j))
            .filter((x, y) => x >= 0 && y >= 0 && y < map.length && x < map(y).length)
            .filter((x, y) => elevation(map(y)(x)) - elevation(map(j)(i)) <= maxDiff)
            .foreach((x, y) => edges.append((curr, Node((x, y), map(y)(x)))))

    (nodes.toList, edges.toList)

def neighbours(node: Node, es: List[(Node, Node)]) = es.filter((from, _) => from == node).map(_._2)

def distances(start: Node, ns: List[Node], es: List[(Node, Node)]) =
    val visited = mutable.HashSet[Node]()
    val dist = mutable.HashMap[Node, Int]()
    val toVisit = mutable.Stack[(Node, Int)]()

    toVisit.push((start, -1))

    while toVisit.nonEmpty do
        val (node, prevDist) = toVisit.pop()
        if visited.contains(node) then
            if dist.get(node).getOrElse(Int.MinValue) > prevDist + 1 then
                dist.put(node, prevDist + 1)
                neighbours(node, es).foreach(n => toVisit.push((n, prevDist + 1)))
        else
            visited.add(node)
            val availNeighbours = neighbours(node, es)
                .filter(n => es.filter((from, to) => n == from && to == node).nonEmpty)
            if availNeighbours.nonEmpty then
                val dst = availNeighbours.map(dist.get).flatten.minOption.getOrElse(prevDist)
                dist.put(node, dst + 1)
                neighbours(node, es).foreach(n => toVisit.push((n, dst + 1)))

    dist

def partI(ns: List[Node], es: List[(Node, Node)]) = distances(ns.find(_.isStart).get, ns, es)(ns.find(_.isEnd).get)

def partII(ns: List[Node], es: List[(Node, Node)]) =
    val end = ns.find(_.isEnd).get
    ns.filter(node => elevation(node.height) == 'a')
         .map(distances(_, ns, es))
         .map(_.get(end)).flatten.min

@main def main(input: String) =
    val (nodes, edges) = graph(mapFromInput(input))

    println(s"Part I: ${partI(nodes, edges)}")
    println(s"Part II: ${partII(nodes, edges)}")
