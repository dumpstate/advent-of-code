(define-module (aoc-common))

(use-modules (ice-9 textual-ports))

(define-public (read-file filename)
    (call-with-input-file
        filename
        (lambda (port) (get-string-all port))))

(define-public (read-lines filename)
    (string-split (read-file filename) #\newline))

(define-public (read-lines-as-numbers filename)
    (filter number?
        (map string->number (read-lines filename))))

(define-public (sum nums)
    (define (loop ns acc)
        (if (null? ns)
            acc
            (loop (cdr ns) (+ acc (car ns)))))
            
    (loop nums 0))

(define-public (l1 pt)
    (sum (map abs pt)))

(define-public (member? pt wire)
    (if (null? wire)
        #f
        (if (equal? pt (car wire))
            #t
            (member? pt (cdr wire)))))

(define-public (intersections w1 w2)
    (define (ints w1 w2 acc)
        (if (null? w1)
            acc
            (if (member? (car w1) w2)
                (ints (cdr w1) w2 (cons (car w1) acc))
                (ints (cdr w1) w2 acc))))
    (ints w1 w2 '()))

(define-public (minimum lst)
    (define (min lst acc)
        (if (null? lst)
            acc
            (if (< (car lst) acc)
                (min (cdr lst) (car lst))
                (min (cdr lst) acc))))
    (min lst (car lst)))
