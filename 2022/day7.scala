import scala.io.Source
import java.nio.file.Paths

val CdRx = raw".+ cd (.*)".r
val DirRx = raw"dir (.*)".r
val FileRx = raw"(\d+) (.*)".r

case class File(name: String, dir: String = "", size: Int = 0):
    def path = Paths.get(dir, name).toString
    def du(fs: List[File]) = fs.filter(_.dir.startsWith(path)).map(_.size).sum

def filesFromInput(input: String) =
    Source.fromFile(input).getLines().foldLeft(("", List[File]())) {
        case ((path, files), next) => next match
            case "$ cd /" => ("/", File("/") :: files)
            case "$ cd .." => (Paths.get(path).getParent.toString, files)
            case CdRx(name) => (Paths.get(path, name).toString, files)
            case DirRx(name) => (path, File(name, path) :: files)
            case FileRx(size, name) => (path, File(name, path, size.toInt) :: files)
            case _ => (path, files)
    }._2

def partI(fs: List[File]) = fs.map(_.du(fs)).filter(_ <= 100000).sum

def partII(fs: List[File]) = fs.map(_.du(fs)).sorted.find(_ >= 30000000 - (70000000 - File("/").du(fs)))

@main def main(input: String) =
    val fs = filesFromInput(input)

    println(s"Part I: ${partI(fs)}")
    println(s"Part II: ${partII(fs).get}")
