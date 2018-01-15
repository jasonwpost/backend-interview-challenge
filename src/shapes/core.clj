(ns shapes.core)

(defn uuid [] (str (java.util.UUID/randomUUID)))
(defn third [coll] (nth coll 2))
(defn choose1 [n] (when (seq n) (rand-nth (seq n))))

(defn generate-point
  []
  [(- (rand-int 101) 50)
   (- (rand-int 101) 50)])

(defn generate-shape
  [points]
  (mapv (fn [_] (generate-point))
        (range points)))

(def shapes {:triangular 100
             :square 103
             :pentagon 109
             :hexagon 88
             :heptagon 99})

(defn generate-shapes []
  (let [triangulars (take (:triangular shapes) (repeatedly #(generate-shape 3)))
        squares (take (:square shapes) (repeatedly #(generate-shape 4)))
        pentagons (take (:pentagon shapes) (repeatedly #(generate-shape 5)))
        hexagon (take (:hexagon shapes) (repeatedly #(generate-shape 6)))
        octagon (take (:heptagon shapes) (repeatedly #(generate-shape 7)))]
    (shuffle (concat triangulars squares pentagons hexagon octagon))))

(defn serialize-point
  [point]
  (clojure.string/join "," point))

(defn serialize-shape
  [shape]
  (clojure.string/join ";" (map #(clojure.string/join "," %) shape)))

(defn serialize-db [ids shapes]
  (doall
    (->> (sort-by first (zipmap ids shapes))
         (into [])
         (map (fn [[id shape]] (str id "|" (serialize-shape shape)))))))

(defn write-to-file [filename data]
  (spit filename (clojure.string/join "\n" data)))

(defn create-op [id shape]
  (let [op (choose1 [:delete-shape :create-shape :delete-point :update-point :add-point])]
    (case op
      :delete-shape [op id]
      :create-shape [op (uuid) (generate-shape (+ 3 (rand-int (count shapes))))]
      :delete-point [op id (rand-int (count shape))]
      :update-point [op id (rand-int (count shape)) (choose1 ["+" "-"]) (generate-point)]
      :add-point [op id (generate-point)]
      nil)))

(defn create-ops [ids shapes]
  (->> (map (fn [_] (let [index (rand-int (count shapes))]
                     [(get ids index)
                      (get shapes index)]))
            (range 50))
       distinct
       (map (fn [[id shape]] (create-op id shape)))
       (remove nil?)))

(defn serialize-op [[op id & data]]
  (println "op: " op " data: "data)
  (case op
    :delete-shape (str (name op) "|" id)
    :create-shape (str (name op) "|" id "|" (serialize-shape (first data)))
    :delete-point (str (name op) "|" id "|" (first data))
    :update-point (str (name op) "|" id "|" (first data) "|" (second data) "|" (serialize-point (third data)))
    :add-point (str (name op) "|" id "|" (serialize-point (first data)))
    (throw (ex-info "Operation not found!" {:op op}))))

(defn main []
  (let [shapes (generate-shapes)
        ids (into [] (repeatedly (count shapes) uuid))
        ops (map serialize-op (create-ops ids shapes))
        db (serialize-db ids shapes)]
    (write-to-file "shapes.db" db)
    (write-to-file "operations.data" ops)))
