import scala.collection.mutable.Stack
import scala.collection.mutable.IndexedBuffer
import scala.io.Source

case class Packet(items: List[Int | Packet]):
    val size = items.length
    val isEmpty = size == 0
    def at(ix: Int) = items.lift(ix)

object Packet:
    def apply(value: Int | Packet): Packet = Packet(List(value))
    def apply(line: String): Packet =
        var ix = 0
        var digit = ""
        val stack = Stack[IndexedBuffer[Int | Packet]]()

        while ix < line.length do
            line(ix) match
                case '[' => stack.push(IndexedBuffer[Int | Packet]())
                case ']' =>
                    if digit.length > 0 then
                        stack.top.addOne(digit.toInt)
                        digit = ""
                    val curr = stack.pop()
                    if stack.isEmpty then
                        return Packet(curr.toList)
                    stack.top.addOne(Packet(curr.toList))
                case ',' if digit.length > 0 =>
                    stack.top.addOne(digit.toInt)
                    digit = ""
                case ',' =>
                case d => digit = digit + d

            ix += 1

        throw Exception("Invalid input")

def packetsFromInput(input: String) =
    Source.fromFile(input)
        .getLines.sliding(3, 3)
        .map(_.take(2).map(Packet.apply))
        .map(packets => (packets(0), packets(1)))
        .toList

def isValid(lPacket: Packet, rPacket: Packet, ix: Int = 0): Boolean =
    if ix >= Math.max(lPacket.size, rPacket.size) then
        true
    else
        (lPacket.at(ix), rPacket.at(ix)) match
            case (Some(l), Some(r)) => (l, r) match
                case (a: Packet, b: Packet) if a.isEmpty && b.isEmpty => isValid(lPacket, rPacket, ix + 1)
                case (a: Packet, b: Packet) => isValid(a, b)
                case (a: Int, b: Packet) => isValid(Packet(a), b)
                case (a: Packet, b: Int) => isValid(a, Packet(b))
                case (a: Int, b: Int) if a > b => false
                case (a: Int, b: Int) if a < b => true
                case (a: Int, b: Int) if a == b => isValid(lPacket, rPacket, ix + 1)
            case (None, Some(_)) => true
            case (Some(_), None) => false 

def partI(packets: List[(Packet, Packet)]) =
    packets.zipWithIndex
        .filter((pckts, _) => isValid(pckts._1, pckts._2))
        .map(_._2 + 1).sum

def partII(packets: List[(Packet, Packet)]) =
    val (d1, d2) = (Packet(Packet(2)), Packet(Packet(6)))
    val sorted = (d1 :: d2 :: packets.flatMap(_.toList))
        .sortWith((a, b) => isValid(b, a)).reverse
        .zipWithIndex

    (d1 :: d2 :: Nil).map(p => sorted.find(_._1 == p).get._2 + 1).product

@main def main(input: String) =
    val packets = packetsFromInput(input)

    println(s"Part I: ${partI(packets)}")
    println(s"Part II: ${partII(packets)}")
