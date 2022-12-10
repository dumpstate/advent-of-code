import scala.io.Source

val AddXRx = raw"addx (-?)(\d+)".r

enum Instr:
    case Noop
    case AddX(x: Int)

def instructionsFromInput(input: String) =
    Source.fromFile(input)
        .getLines
        .map(line => line match
            case "noop" => Instr.Noop
            case AddXRx(n, x) => Instr.AddX(x.toInt * (if n == "-" then -1 else 1)))
        .toList

def signal(is: List[Instr], curr: Int, res: List[Int], delay: Int): Vector[Int] = is match
    case instr :: tail => instr match
        case Instr.Noop => signal(tail, curr, curr :: res, 0)
        case Instr.AddX(x) =>
            if delay < 1 then signal(instr :: tail, curr, curr :: res, delay + 1)
            else signal(tail, curr + x, curr + x :: res, 0)
    case Nil => res.reverse.toVector

def partI(sig: Vector[Int]) =
    List(20, 60, 100, 140, 180, 220)
        .map(ix => sig(ix - 1) * ix)
        .sum

def partII(sig: Vector[Int]) =
    (0 to 240).zip(sig).take(240)
        .map((ix, s) => (ix % 40, s))
        .map((ix, s) => if ix >= s - 1 && ix <= s + 1 then "#" else ".")
        .mkString
        .sliding(40, 40)
        .mkString("\n")

@main def main(input: String) =
    val sig = signal(instructionsFromInput(input), 1, List(1), 0)

    println(s"Part I: ${partI(sig)}")
    println(s"Part II:\n${partII(sig)}")
