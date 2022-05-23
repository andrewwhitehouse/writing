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

(deftest test-two-largest
  (testing "all different"
    (is (= [3 2] (ch1/two-largest 1 2 3))))
  (testing "two equal"
    (is (= [2 1] (ch1/two-largest 1 2 1)))
    (is (= [2 1] (ch1/two-largest 1 1 2))))
  (testing "all equal"
    (is (= [3 3] (ch1/two-largest 3 3 3))))
  (testing "reverse order"
    (is (= [10 9] (ch1/two-largest 10 9 8)))))
