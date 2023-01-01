import scala.collection.mutable.{ArrayBuffer, HashMap, HashSet}
import scala.io.Source

enum Dir:
    case Up, Down, Left, Right

def blizzardsFromInput(input: String) =
    val lines = Source.fromFile(input).getLines.toVector
    val map = lines.slice(1, lines.length - 1)
        .map(line => line.slice(1, line.length - 1).split("").map(_.charAt(0)).toVector)
    val blizzards = new HashMap[(Int, Int), ArrayBuffer[Dir]]()

    for
        y <- 0 until map.length
        x <- 0 until map(y).length
        if map(y)(x) != '.'
    do
        val blizzard = map(y)(x) match
            case '>' => Dir.Right
            case '<' => Dir.Left
            case '^' => Dir.Up
            case 'v' => Dir.Down
            case _ => throw Exception("Invalid input")
        if blizzards.contains((x, y)) then
            blizzards((x, y)).addOne(blizzard)
        else
            val arr = new ArrayBuffer[Dir]()
            arr.addOne(blizzard)
            blizzards.put((x, y), arr)

    (map, blizzards)

def simulate(map: Vector[Vector[Char]], blizzards: HashMap[(Int, Int), ArrayBuffer[Dir]], time: Int) =
    for t <- 1 to time do
        val toAdd = new ArrayBuffer[((Int, Int), Dir)]()
        for
            pos <- blizzards.keys
            dir <- blizzards(pos).clone()
        do
            val dirs = blizzards(pos)
            dirs.remove(dirs.indexOf(dir))

            val (x, y) = pos
            val nextPos = dir match
                case Dir.Right => ((x + 1) % map(y).length , y)
                case Dir.Left => (Math.floorMod(x - 1, map(y).length), y)
                case Dir.Down => (x, (y + 1) % map.length)
                case Dir.Up => (x, Math.floorMod(y - 1, map.length))

            toAdd.addOne((nextPos, dir))

        for (pos, dir) <- toAdd do
            if blizzards.contains(pos) then
                blizzards(pos).addOne(dir)
            else
                val arr = new ArrayBuffer[Dir]()
                arr.addOne(dir)
                blizzards.put(pos, arr)

    blizzards

def neighbours(map: Vector[Vector[Char]], pos: (Int, Int)) = pos match
    case (0, -1) => Iterator((0, 0))
    case (x, y) if x == map(0).length - 1 && y == map.length =>
        Iterator((map(0).length - 1, map.length - 1))
    case (x, y) =>
        val it = new ArrayBuffer[(Int, Int)]()

        if x > 0 then it.addOne((x - 1, y))
        if x < map(0).length - 1 then it.addOne((x + 1, y))
        if y > 0 then it.addOne((x, y - 1))
        if y < map.length - 1 then it.addOne((x, y + 1))
        if x == map(0).length - 1 && y == map.length - 1 then it.addOne((map(0).length - 1, map.length))
        if x == 0 && y == 0 then it.addOne((0, -1))

        it.iterator

def isEmpty(blizzards: HashMap[(Int, Int), ArrayBuffer[Dir]], pos: (Int, Int)) =
    blizzards.get(pos)
        .map(_.isEmpty)
        .getOrElse(true)

def shortestPath(
    map: Vector[Vector[Char]],
    blizzards: HashMap[(Int, Int), ArrayBuffer[Dir]],
    start: (Int, Int),
    end: (Int, Int),
): Int =
    var states = new HashSet[(Int, Int)]()

    states.add(start)

    for t <- LazyList.from(0) do
        simulate(map, blizzards, 1)

        var nextStates = new HashSet[(Int, Int)]()

        for state <- states do
            if state == end then
                return t

            for nextState <- neighbours(map, state) do
                if isEmpty(blizzards, nextState) then
                    nextStates.add(nextState)

            if isEmpty(blizzards, state) then
                nextStates.add(state)

        states = nextStates

    throw new Exception("Not found")

def partI(input: String): Int =
    val (map, blizzards) = blizzardsFromInput(input)
    val (start, end) = ((0, -1), (map(0).length - 1, map.length))

    shortestPath(map, blizzards, start, end)

def partII(input: String): Int =
    val (map, blizzards) = blizzardsFromInput(input)
    val (start, end) = ((0, -1), (map(0).length - 1, map.length))

    shortestPath(map, blizzards, start, end) + 1 +
        shortestPath(map, blizzards, end, start) + 1 +
        shortestPath(map, blizzards, start, end)

@main def main(input: String) =
    println(s"Part I: ${partI(input)}")
    println(s"Part II: ${partII(input)}")
