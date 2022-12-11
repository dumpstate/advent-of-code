import scala.collection.mutable.Stack
import scala.io.Source

case class Expr(op: String, r: Option[Int]):
    def eval(level: Long) = (op, r) match
        case ("*", None) => level * level
        case ("*", Some(n)) => level * n
        case ("+", None) => level + level
        case ("+", Some(n)) => level + n

case class Monkey(items: Stack[Long], expr: Expr, testDiv: Int, trueCase: Int, falseCase: Int):
    var inspections = 0L

    def inspect(item: Long, div: Int) =
        inspections += 1

        val next = expr.eval(item) / div
        if next % testDiv == 0 then
            (trueCase, next)
        else
            (falseCase, next)

def monkeysFromInput(input: String) =
    Source.fromFile(input).getLines.sliding(7, 7)
        .map(lines => Monkey(
            items=Stack(lines(1).substring(18).split(", ").map(_.toLong): _*),
            expr=Expr(lines(2).substring(23, 24), lines(2).substring(25).toIntOption),
            testDiv=lines(3).substring(21).toInt,
            trueCase=lines(4).substring(29).toInt,
            falseCase=lines(5).substring(30).toInt)).toList

def round(monkeys: List[Monkey], div: Int) =
    val mod = monkeys.map(_.testDiv).product
    for monkey <- monkeys do
        while !monkey.items.isEmpty do
            val (nextMonkey, nextItem) = monkey.inspect(monkey.items.pop(), div)
            monkeys(nextMonkey).items.addOne(nextItem % mod)

def score(input: String, rounds: Int, div: Int): Long =
    val ms = monkeysFromInput(input)
    for _ <- 1 to rounds do round(ms, div)
    ms.map(_.inspections).sorted.reverse.take(2).product

@main def main(input: String) =
    println(s"Part I: ${score(input, 20, 3)}")
    println(s"Part II: ${score(input, 10000, 1)}")
