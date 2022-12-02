import scala.io.Source


enum Shape(val score: Int):
    case Rock extends Shape(1)
    case Paper extends Shape(2)
    case Scissors extends Shape(3)

object Shape:
    def opponent(code: String) = code match {
        case "A" => Rock
        case "B" => Paper
        case "C" => Scissors
        case _ => throw Exception("Invalid code")
    }

    def mine(code: String) = code match {
        case "X" => Rock
        case "Y" => Paper
        case "Z" => Scissors
        case _ => throw Exception("Invalid code")
    }

val Rules = List(
    (Shape.Rock, Shape.Scissors),
    (Shape.Paper, Shape.Rock),
    (Shape.Scissors, Shape.Paper)
)

def turnsFromInput(input: String): List[(Shape, Shape)] =
    Source.fromFile(input)
        .getLines
        .map(_.split(" ").toList match {
            case List(opponent, mine) =>
                (Shape.opponent(opponent), Shape.mine(mine))
            case _ => throw Exception("Invalid input")
        })
        .toList

def evaluate(opponent: Shape, mine: Shape): Int =
    if opponent == mine then
        3
    else
        Rules.find((x, _) => x == mine)
            .filter((_, op) => opponent == op)
            .map(_ => 6)
            .getOrElse(0)

def partI(turns: List[(Shape, Shape)]) =
    turns.map {
        case (opponent, mine) => mine.score + evaluate(opponent, mine)
    }.sum

def partII(turns: List[(Shape, Shape)]) = turns.map {
    case (opponent, result) => result match {
        // should loose
        case Shape.Rock => Rules
            .find((op, _) => op == opponent)
            .map((_, mine) => mine.score)
            .get
        // should draw
        case Shape.Paper => 3 + opponent.score
        // should win
        case Shape.Scissors => 6 + Rules
            .find((_, op) => op == opponent)
            .map((mine, _) => mine.score)
            .get
    }
}.sum

@main def main(input: String) =
    val turns = turnsFromInput(input)

    println(s"Part I: ${partI(turns)}")
    println(s"Part II: ${partII(turns)}")
