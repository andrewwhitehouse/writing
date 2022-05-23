(ns sicp.chapter1)

(defn square [x]
      (* x x))

;; ----------------

(comment

  (defn p [] (p))

  (defn test [x y]
        (if (= x 0) 0 y))

  (test 0 (p))

  )

(comment

  (defn p [] (p))

  (defn test [x y]
        (if (= x 0) 0 (y)))

  (test 0 p)

  )

(comment defn two-largest [a b c]
         (let [first-pair (if (> a b) {:larger a :smaller b} {:larger b :smaller a})
               remaining-result (if (> c (:larger first-pair))
                                  {:first c :second (:larger first-pair)}
                                  (if (> c (:smaller first-pair))
                                    {:first (:larger first-pair) :second c}
                                    {:first (:larger first-pair) :second (:smaller first-pair)}))]
              [(:first remaining-result) (:second remaining-result)]))


(comment defn two-largest [a b c]
         (let [a-b-comparison (if (> a b) [a b] [b a])]
              (if (> c (first a-b-comparison))
                [c (first a-b-comparison)]
                (if (> c (second a-b-comparison))
                  [(first a-b-comparison) c]
                  a-b-comparison))))

(comment defn two-largest [a b c]
         (let [max-a-b (max a b)
               min-a-b (min a b)]
              (if (> c max-a-b)
                [c max-a-b]
                [max-a-b (max min-a-b c)])))

(comment defn two-largest [& nums]
         (let [sorted (sort > nums)]
              [(first sorted) (second sorted)]))

(comment defn two-largest [& nums]
         (let [first-two-values (fn [[a b]] [a b])]
              (first-two-values (sort > nums))))

(comment defn two-largest [& nums]
         (->> (sort > nums)
              ((fn [[a b]] [a b]))))

(comment defn two-largest [a b c]
  (let [max-a-b (max a b)]
     (if (> c max-a-b)
       [c max-a-b]
       [max-a-b (max (min a b) c)])))

(defn two-largest [a b c]
  (cond 
    (>= a b c) [a b]
    (>= b a c) [b a]
    (>= c b a) [c b]
    (>= a c b) [a c]
    (>= b c a) [b c]
    (>= c a b) [c a]))
