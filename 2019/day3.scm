(use-modules (aoc-common))

(define (next-point head dir)
    (let ((x (car head))
          (y (car (cdr head))))
        (case (car (string->list dir))
            ((#\U) (cons x (cons (+ y 1) '())))
            ((#\D) (cons x (cons (- y 1) '())))
            ((#\L) (cons (- x 1) (cons y '())))
            ((#\R) (cons (+ x 1) (cons y '()))))))

(define (wire pts cmds)
    (if (null? cmds)
        pts
        (let ((dir (car (car cmds)))
              (steps (car (cdr (car cmds))))
              (head (car pts)))
            (wire (cons (next-point head dir) pts)
                  (if (> steps 1)
                      (cons (list dir (- steps 1)) (cdr cmds))
                      (cdr cmds))))))

(define (parse-wire-cmds line)
    (define (split-cmd cmd)
        (list
            (substring cmd 0 1)
            (string->number (substring cmd 1))))

    (map split-cmd (string-split line #\,)))

(define (main args)
    (define lines (read-lines (list-ref args 1)))
    (define w1 (wire (list '(0 0)) (parse-wire-cmds (list-ref lines 0))))
    (define w2 (wire (list '(0 0)) (parse-wire-cmds (list-ref lines 1))))
    (define intersection-pts (intersections (cdr (reverse w1)) (cdr (reverse w2))))
    (format #t "Part I: ~a\n" (minimum (map l1 intersection-pts))))
