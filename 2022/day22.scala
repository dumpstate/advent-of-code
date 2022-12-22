import scala.io.Source

type DirInstr = 'R' | 'L'
type BoardChar = ' ' | '.' | '#'
type Board = Vector[Vector[BoardChar]]
type Instr = List[Int | DirInstr]

def boardFromInput(input: String): (Board, Instr) =
    val lines = Source.fromFile(input)
        .getLines.toVector
    val board = lines.slice(0, lines.length - 2)
        .map(_.split("").map(_.charAt(0))
            .toVector.asInstanceOf[Vector[BoardChar]])
    val instruction = lines.last
        .split("(?<=\\d)(?=\\D)|(?=\\d)(?<=\\D)")
        .map(entry => entry.toIntOption.getOrElse(entry.charAt(0)))
        .toList.asInstanceOf[List[Int | DirInstr]]

    (board, instruction)

def startPos(board: Board) = (board(0).indexOf('.'), 0)

def rotate(dir: Int, by: DirInstr) =
    val next = by match
        case 'R' => (dir + 1) % 4
        case 'L' => (dir - 1) % 4

    if next < 0 then 4 + next else next

def nextRight(board: Board, pos: (Int, Int)) =
    var (x, y) = pos
    while
        x = (x + 1) % board(y).length
        board(y)(x) == ' '
    do ()
    if board(y)(x) == '#' then pos else (x, y)

def nextLeft(board: Board, pos: (Int, Int)) =
    var (x, y) = pos
    while
        x = if x == 0 then board(y).length - 1 else x - 1
        board(y)(x) == ' '
    do ()
    if board(y)(x) == '#' then pos else (x, y)

def nextDown(board: Board, pos: (Int, Int)) = 
    var (x, y) = pos
    while
        y = (y + 1) % board.length
        board(y).length < x || board(y)(x) == ' '
    do ()
    if board(y)(x) == '#' then pos else (x, y)

def nextUp(board: Board, pos: (Int, Int)) =
    var (x, y) = pos
    while
        y = if y == 0 then board.length - 1 else y - 1
        board(y).length < x || board(y)(x) == ' '
    do ()
    if board(y)(x) == '#' then pos else (x, y)

def step(board: Board, pos: (Int, Int), dir: Int) =
    dir match
        case 0 => nextRight(board, pos)
        case 1 => nextDown(board, pos)
        case 2 => nextLeft(board, pos)
        case 3 => nextUp(board, pos)
        case _ => throw Exception(s"Invalid direction: $dir")

def password(pos: (Int, Int), dir: Int) = pos match
    case (x, y) => 1000 * (y + 1) + 4 * (x + 1) + dir

def partI(board: Board, instrs: Instr) =
    var (pos, dir) = (startPos(board), 0)

    for instr <- instrs do instr match
        case in: Int =>
            for _ <- 1 to in do
                pos = step(board, pos, dir)
        case in: DirInstr => dir = rotate(dir, in)

    password(pos, dir)

@main def main(input: String) =
    val (board, instr) = boardFromInput(input)

    println(s"Part I: ${partI(board, instr)}")
