(use-modules (aoc-common)
             (ice-9 format))

(define (fuel-req num) (- (floor/ num 3) 2))

(define (total-fuel-req num)
    (define (loop rem acc)
        (let ((req (fuel-req rem)))
            (if (<= req 0)
                acc
                (loop req (+ acc req)))))

    (loop num 0))

(define (main args)
    (define nums (read-lines-as-numbers (list-ref args 1)))

    (format #t "Part I: ~a\n" (sum (map fuel-req nums)))
    (format #t "Part II: ~a\n" (sum (map total-fuel-req nums))))
