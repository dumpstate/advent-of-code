import scala.io.Source

val SNAFU = Vector("0", "1", "2", "=", "-")

def snafuToDecimal(num: String) =
    num.reverse.split("").foldLeft((0L, 0L)) { case ((res, off), entry) =>
        val value = entry match
            case "-" => -1
            case "=" => -2
            case x => x.toLong
        
        (res + Math.pow(5L, off).toLong * value, off + 1) }._1

def decimalToSnafu(dec: Long): String =
    if dec == 0 then ""
    else
        val (rem, div) = (dec % 5, (dec + 2) / 5)

        decimalToSnafu(div) + SNAFU(rem.toInt)

def partI(numbers: List[String]) = decimalToSnafu(numbers.map(snafuToDecimal).sum)

@main def main(input: String) =
    val ns = Source.fromFile(input).getLines.toList

    println(s"Part I: ${partI(ns)}")
