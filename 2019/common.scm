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
