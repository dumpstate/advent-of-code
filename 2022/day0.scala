import scala.io.Source

@main def main(input: String) =
    println("Hello!")

    for (
        num <- Source.fromFile(input)
            .getLines
            .map(_.toInt)
    )
        println(s"number: $num")
