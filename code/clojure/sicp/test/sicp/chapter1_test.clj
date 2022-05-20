(ns sicp.chapter1-test
  (:require [clojure.test :refer :all]
            [sicp.chapter1 :as ch1]))

(defn is-close [a b tolerance]
  (< (Math/abs (- a b)) tolerance))

(deftest test-square
  (testing "square integer."
    (is (= 9 (ch1/square 3))))
  (testing "square decimal"
    (is (is-close (ch1/square 2.2) 4.84 0.001))))
