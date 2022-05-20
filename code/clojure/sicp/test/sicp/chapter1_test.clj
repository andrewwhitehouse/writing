(ns sicp.chapter1-test
  (:require [clojure.test :refer :all]
            [sicp.chapter1 :refer :all]))

(defn is-close [a b tolerance]
  (< (Math/abs (- a b)) tolerance))

(deftest test-square
  (testing "square integer."
    (is (= 9 (square 3))))
  (testing "square decimal"
    (is (is-close (square 2.2) 4.84 0.001))))
