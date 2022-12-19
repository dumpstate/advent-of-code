import scala.collection.mutable.HashMap
import scala.io.Source

enum Mineral:
    case Ore, Clay, Obsidian, Geode

case class Blueprint(id: Int, ore: Int, clay: Int, obsidian: (Int, Int), geode: (Int, Int)):
    def canAfford(state: (Int, Int, Int, Int), mineral: Mineral): Boolean = mineral match
        case Mineral.Ore => state._1 >= ore
        case Mineral.Clay => state._1 >= clay
        case Mineral.Obsidian => state._1 >= obsidian._1 && state._2 >= obsidian._2
        case Mineral.Geode => state._1 >= geode._1 && state._3 >= geode._2

    def canAfford(state: (Int, Int, Int, Int)): Set[Mineral] = state match
        case (o, c, ob, g) => Set(Mineral.Ore, Mineral.Clay, Mineral.Obsidian, Mineral.Geode)
            .filter(canAfford(state, _))

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
        val res = (state, robots) match
            case ((_, _, _, geo), _) if timeLeft <= 0 => geo
            case ((ore, clay, obs, geo), (rOre, rClay, rObs, rGeo)) if b.canAfford(state, Mineral.Geode) =>
                maxGeodes(b, cache)(
                    (ore + rOre - b.geode._1, clay + rClay, obs + rObs - b.geode._2, geo + rGeo),
                    (rOre, rClay, rObs, rGeo + 1),
                    timeLeft - 1)
            case ((ore, clay, obs, geo), (rOre, rClay, rObs, rGeo)) if (
                ore >= b.ore &&
                ore >= b.clay &&
                ore >= b.obsidian._1 &&
                ore >= b.geode._1 &&
                clay >= b.obsidian._2
             ) =>
                maxGeodes(b, cache)(
                    (ore + rOre - b.obsidian._1, clay + rClay - b.obsidian._2, obs + rObs, geo + rGeo),
                    (rOre, rClay, rObs + 1, rGeo),
                    timeLeft - 1)
            case ((ore, clay, obs, geo), (rOre, rClay, rObs, rGeo)) =>
                val (nOre, nClay, nObs, nGeo) = (ore + rOre, clay + rClay, obs + rObs, geo + rGeo)
                val wait = () => maxGeodes(b, cache)((nOre, nClay, nObs, nGeo), robots, timeLeft - 1)
                b.canAfford(state).flatMap {
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

def partII(blueprints: List[Blueprint]) =
    blueprints.take(3).iterator
        .map(maxGeodes(_, HashMap())((0, 0, 0, 0), (1, 0, 0, 0), 32))
        .product

@main def main(input: String) =
    val blueprints = Source.fromFile(input)
        .getLines.map(Blueprint.from)
        .toList

    println(s"Part I: ${partI(blueprints)}")
    println(s"Part II: ${partII(blueprints)}")
