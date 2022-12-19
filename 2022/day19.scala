import scala.collection.mutable.HashMap
import scala.io.Source

enum Mineral:
    case Ore, Clay, Obsidian, Geode

case class Blueprint(id: Int, ore: Int, clay: Int, obsidian: (Int, Int), geode: (Int, Int)):
    def canAfford(state: (Int, Int, Int, Int)) = state match
        case (o, c, ob, g) =>
            var res = List.empty[Mineral]
            if o >= ore then res = Mineral.Ore :: res
            if o >= clay then res = Mineral.Clay :: res
            if o >= obsidian._1 && c >= obsidian._2 then res = Mineral.Obsidian :: res
            if o >= geode._1 && ob >= geode._2 then res = Mineral.Geode :: res
            res.toSet

object Blueprint:
    def from(line: String) =
        val split = line.split("\\.")
        val head = split(0).split(": ")
        val obs = split(2).substring(26).trim.split(" ")
        val geode = split(3).substring(23).trim.split(" ")
        Blueprint(
            id=head(0).split(" ")(1).toInt,
            ore=head(1).substring(21).trim.split(" ")(0).toInt,
            clay=split(1).substring(22).trim.split(" ")(0).toInt,
            obsidian=(obs(0).toInt, obs(3).toInt),
            geode=(geode(0).toInt, geode(3).toInt))

def maxGeodes(b: Blueprint, cache: HashMap[((Int, Int, Int, Int), (Int, Int, Int, Int), Int), Int])(
    state: (Int, Int, Int, Int),
    robots: (Int, Int, Int, Int),
    timeLeft: Int): Int = cache.get((state, robots, timeLeft)).getOrElse {
        val res = state match
            case (_, _, _, geo) if timeLeft <= 0 => geo
            case (ore, clay, obs, geo) =>
                val (rOre, rClay, rObs, rGeo) = robots
                val (nOre, nClay, nObs, nGeo) = (ore + rOre, clay + rClay, obs + rObs, geo + rGeo)
                val wait = () => maxGeodes(b, cache)((nOre, nClay, nObs, nGeo), robots, timeLeft - 1)
                val alternatives = b.canAfford(state)
                if alternatives.contains(Mineral.Geode) then
                    maxGeodes(b, cache)(
                        (nOre - b.geode._1, nClay, nObs - b.geode._2, nGeo),
                        (rOre, rClay, rObs, rGeo + 1),
                        timeLeft - 1)
                else
                    alternatives.flatMap {
                        case Mineral.Ore =>
                            if rOre >= b.ore && rOre >= b.clay && rOre >= b.obsidian._1 && rOre >= b.geode._1 then None
                            else Some(maxGeodes(b, cache)(
                                (nOre - b.ore, nClay, nObs, nGeo),
                                (rOre + 1, rClay, rObs, rGeo),
                                timeLeft - 1))
                        case Mineral.Clay =>
                            if rClay >= b.obsidian._2 then None
                            else Some(maxGeodes(b, cache)(
                                (nOre - b.clay, nClay, nObs, nGeo),
                                (rOre, rClay + 1, rObs, rGeo),
                                timeLeft - 1))
                        case Mineral.Obsidian =>
                            if rObs >= b.geode._2 then None
                            else Some(maxGeodes(b, cache)(
                                (nOre - b.obsidian._1, nClay - b.obsidian._2, nObs, nGeo),
                                (rOre, rClay, rObs + 1, rGeo),
                                timeLeft - 1))
                    }.maxOption.map(Math.max(_, wait())).getOrElse(wait())
        cache.put((state, robots, timeLeft), res)
        res
    }

def partI(blueprints: List[Blueprint]) =
    blueprints.iterator
        .map(b => (b.id, maxGeodes(b, HashMap())((0, 0, 0, 0), (1, 0, 0, 0), 24)))
        .map((id, max) => id * max).sum

@main def main(input: String) =
    val blueprints = Source.fromFile(input)
        .getLines.map(Blueprint.from)
        .toList

    println(s"Part I: ${partI(blueprints)}")
