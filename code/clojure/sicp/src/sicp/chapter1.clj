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
