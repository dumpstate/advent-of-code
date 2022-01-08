(use-modules (aoc-common)
             (ice-9 format)
             (ice-9 arrays))

(define (exec tape ix)
    (if (>= ix (array-length tape))
        (array-ref tape 0)
        (let ((op (array-ref tape ix))
              (l (array-ref tape (array-ref tape (+ ix 1))))
              (r (array-ref tape (array-ref tape (+ ix 2))))
              (t-ix (array-ref tape (+ ix 3))))
            (case op
                ((1)
                    (array-set! tape (+ l r) t-ix)
                    (exec tape (+ ix 4)))
                ((2)
                    (array-set! tape (* l r) t-ix)
                    (exec tape (+ ix 4)))
                ((99) (array-ref tape 0))))))

(define (exec-with-input tape noun verb)
    (let ((tape-cpy (array-copy tape)))
        (array-set! tape-cpy noun 1)
        (array-set! tape-cpy verb 2)
    
        (exec tape-cpy 0)))

(define (find-input tape target noun verb)
    (if (= target (exec-with-input tape noun verb))
        (+ (* 100 noun) verb)
        (if (<= verb 100)
            (find-input tape target noun (+ verb 1))
            (find-input tape target (+ noun 1) 0))))

(define (main args)
    (define nums
        (list->array 1
            (map string->number
                (string-split
                    (list-ref (read-lines (list-ref args 1)) 0)
                    #\,))))

    (format #t "Part I: ~a\n" (exec-with-input nums 12 2))
    (format #t "Part II: ~a\n" (find-input nums 19690720 0 0)))
