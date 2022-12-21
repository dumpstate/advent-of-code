import scala.io.Source

sealed abstract class Expression:
    def evaluate(ms: Map[String, Monkey]): Long

case class Unary(num: Long) extends Expression:
    def evaluate(ms: Map[String, Monkey]) = num

case class Binary(l: String, op: Char, r: String) extends Expression:
    def evaluate(ms: Map[String, Monkey]) =
        val lVal = ms(l).yell(ms)
        val rVal = ms(r).yell(ms)
        op match
            case '+' => lVal + rVal
            case '-' => lVal - rVal
            case '*' => lVal * rVal
            case '/' => lVal / rVal

case class Monkey(name: String, expr: Expression):
    def yell(ms: Map[String, Monkey]) = expr.evaluate(ms)

    def yellWithGuess(ms: Map[String, Monkey], guess: Long, name: String = "humn") =
        yell(ms.updated(name, Monkey(name, Unary(guess))))

    def dependsOn(ms: Map[String, Monkey], name: String): Boolean = expr match
        case Unary(_) => false
        case Binary(l, _, r) if l != name && r != name =>
            ms(l).dependsOn(ms, name) || ms(r).dependsOn(ms, name)
        case Binary(_, _, _) => true

def monkeysFromInput(input: String) = Source.fromFile(input).getLines
    .map(line =>
        val split = line.split(": ")
        (split(0), Monkey(split(0), split(1)
            .toLongOption
            .map(Unary(_))
            .getOrElse(Binary(
                split(1).substring(0, 4),
                split(1).charAt(5),
                split(1).substring(7)))))).toMap

def binSearch(ms: Map[String, Monkey], m: Monkey, min: Long, max: Long, target: Long): Long =
    if min == max then min
    else
        val mid = min + (max - min) / 2
        val minV = m.yellWithGuess(ms, min) - target
        val maxV = m.yellWithGuess(ms, max) - target
        val midV = m.yellWithGuess(ms, mid) - target
        if Math.min(minV, midV) <= 0 && Math.max(minV, midV) >= 0 then
            binSearch(ms, m, min, mid, target)
        else binSearch(ms, m, mid + 1, max, target)

def partI(ms: Map[String, Monkey]) = ms("root").yell(ms)

def partII(ms: Map[String, Monkey]) =
    val Monkey(_, Binary(l, _, r)) = ms("root")
    val (targetM, humnDependentM) = if ms(l).dependsOn(ms, "humn") then (ms(r), ms(l)) else (ms(l), ms(r))
    binSearch(ms, humnDependentM, 0, Math.pow(2, 42).toLong, targetM.yell(ms))

@main def main(input: String) =
    val monkeys = monkeysFromInput(input)

    println(s"Part I: ${partI(monkeys)}")
    println(s"Part II: ${partII(monkeys)}")
