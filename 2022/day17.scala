import scala.collection.mutable
import scala.io.Source

val PatternOffset = 30

case class Rock(str: String):
    val lines = str.split("\n")
    val width = lines.map(_.length).max
    val height = lines.length

    def iterIx() =
        for(
            y <- 0 until lines.length;
            x <- 0 until lines(y).length
            if lines(y)(x) == '#'
        ) yield (x, y)

val Rocks = Stream.continually(List(
    Rock("####"),
    Rock(".#.\n###\n.#."),
    Rock("..#\n..#\n###"),
    Rock("#\n#\n#\n#"),
    Rock("##\n##"))).flatten

class Cave(map: mutable.ArrayBuffer[mutable.ArrayBuffer[Char]], jetPattern: Vector[Char]):
    var jetIx = 0

    def height() = map.length

    def nonEmptyHeight() =
        var y = 0
        while y < map.length && !isEmpty(map(y)) do y += 1
        y

    def isEmpty(line: Iterable[Char]) = line.find(_ != '.').isEmpty

    def nextJet() =
        if jetIx >= jetPattern.length then
            jetIx = 0

        val jet = jetPattern(jetIx)
        jetIx += 1
        jet

    def resize(units: Int) =
        for _ <- 1 to units do
            map.addOne(mutable.ArrayBuffer('.', '.', '.', '.', '.', '.', '.'))

    def initRock(rock: Rock) =
        var y = 0
        while y < map.length && !isEmpty(map(y)) do
            y += 1

        resize(y + 3 + rock.height - map.length)

        (2, y + 3 + rock.height - 1)

    def draw(rock: Rock, pos: (Int, Int)) =
        for (xr, yr) <- rock.iterIx() do
            map(pos._2 - yr)(pos._1 + xr) = '#'

    def hash() =
        var (res, yMax) = ("", map.length)
        while
            yMax -= 1
            isEmpty(map(yMax))
        do ()

        for y <- yMax to (yMax - PatternOffset) by -1 if y >= 0 do
            res += map(y).mkString

        res.hashCode()

    def canDraw(rock: Rock, pos: (Int, Int)): Boolean =
        for (xr, yr) <- rock.iterIx() do
            val (x, y) = (xr + pos._1, pos._2 - yr)
            if x >= 7 || x < 0 || y < 0 || map(y)(x) != '.' then
                return false
        true

    def drop(rock: Rock) =
        var (x, y) = initRock(rock)

        var canMove = true
        while canMove do
            nextJet() match
                case '<' => if canDraw(rock, (x - 1, y)) then x -= 1
                case '>' => if canDraw(rock, (x + 1, y)) then x += 1

            if canDraw(rock, (x, y - 1))
                then y -= 1
                else canMove = false

        draw(rock, (x, y))

def jetPatternFromInput(input: String) =
    Source.fromFile(input).getLines.map(_.split("").map(_.toCharArray()(0)).toVector).toList.head

def dropN(jetPattern: Vector[Char], n: Int) =
    val cave = new Cave(mutable.ArrayBuffer(), jetPattern)

    for rock <- Rocks.take(n) do
        cave.drop(rock)

    cave.nonEmptyHeight()

def checkCycle(items: mutable.ArrayBuffer[Int]): (Int, Boolean) =
    for p <- 0 until items.length / 2 do
        val off = p + ((items.length - p) / 2)
        val l = items.slice(p, off)
        val r = items.slice(off, items.length)
        if l == r then
            return (p, true)
    (-1, false)

def partI(pattern: Vector[Char]) = dropN(pattern, 2022)

def partII(pattern: Vector[Char]): Long =
    val cave = new Cave(mutable.ArrayBuffer(), pattern)
    val seen = new mutable.ArrayBuffer[Int]()
    var (ix, rem, res, resHeight) = (0, 0L, 0L, 0)
    var countRem = false

    for rock <- Rocks do
        cave.drop(rock)
        seen.addOne(cave.hash())
        ix += 1

        if countRem then
            rem -= 1

            if rem <= 0 then
                return res + cave.nonEmptyHeight() - resHeight
        else
            val (preambleEndIx, cycleFound) = checkCycle(seen)
            if cycleFound then
                val preambleHeight = dropN(pattern, preambleEndIx)
                val cycleIxSpan = (ix - preambleEndIx) / 2
                val cycleHeight = (cave.nonEmptyHeight() - preambleHeight) / 2
                resHeight = cave.nonEmptyHeight()
                rem = (1000000000000L - ix) % cycleIxSpan
                res = ((1000000000000L - ix) / cycleIxSpan) * cycleHeight + resHeight
                countRem = true

    throw new Exception("Cycle not found")

@main def main(input: String) =
    val pattern = jetPatternFromInput(input)

    println(s"Part I: ${partI(pattern)}")
    println(s"Part II: ${partII(pattern)}")
