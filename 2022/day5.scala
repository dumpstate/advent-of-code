import scala.io.Source
import scala.collection.mutable.Stack

val Command = raw"move (\d+) from (\d+) to (\d+)".r

def cmd(line: String) = line match
    case Command(moves, from, to) => (moves.toInt, from.toInt, to.toInt)

def vStacks(hStacks: List[String]) = hStacks.reverse match
    case ixs :: stacks =>
        for(ix <- ixs.split(" ").filter(_.nonEmpty).map(_.toInt - 1))
            yield stacks.reverse.map(_.charAt(4 * ix + 1)).filter(_ != ' ')

def stacksFromInput(input: String) =
    val (hStacks, cmds, _) = Source.fromFile(input)
        .getLines()
        .foldLeft((List[String](), List[(Int, Int, Int)](), true))((acc, line) => acc match
            case (hStacks, cmds, stacks) if line.isEmpty => (hStacks, cmds, false)
            case (hStacks, cmds, true) => (hStacks ++ Seq(line), cmds, true)
            case (hStacks, cmds, false) => (hStacks, cmds ++ Seq(cmd(line)), false)
        )

    (vStacks(hStacks).map(Stack(_: _*)).toVector, cmds)

def partI(input: String) =
    val (stacks, cmds) = stacksFromInput(input)

    cmds.foreach((count, from, to) =>
        for _ <- 1 to count do
            stacks(to - 1).push(stacks(from - 1).pop())
    )

    stacks.map(_.head).mkString

def partII(input: String) =
    val (stacks, cmds) = stacksFromInput(input)

    cmds.foreach((count, from, to) =>
        for item <- stacks(from - 1).take(count).reverse do
            stacks(from - 1).pop()
            stacks(to - 1).push(item)
    )

    stacks.map(_.head).mkString

@main def main(input: String) =
    println(s"Part I: ${partI(input)}")
    println(s"Part II: ${partII(input)}")
