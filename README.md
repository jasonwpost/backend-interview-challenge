# Backend programming challenge

Tl;dr - just perform the [operation](operations.data) against the [database](shapes.db), and save the result to a text file.

### General rules:
* You can use any programming language you’d like.
* You can use the Internet as much as you’d like (without cheating.)
* We don’t care much about performance.
* Tests are optional (but correction will probably be taken into consideration.)

You need two files:
* [shapes.db](shapes.db)
* [operations.data](operations.data)

### [shapes.db](shapes.db)
It’s the database, stored as a file.

There’s two columns, separated by a pipe “|” symbol.
* The first column is the id of the row.
* The second column are points on a 2D coordinate plane
  * Points are separated by a semi-colon “;”
  * The x position is before the comma, y position is after it.
  * There’s at most 7 points for a given shape, and minimum 3.
  * You can safely assume that each number in the point is less than 32 bits in size.
  * The points themselves having no real meaning or sort order.

### [operations.data](operations.data)
It’s the operations to be performed against the database.

* The amount of columns are variable depending on the operation.
  * The first column is the operation to be performed
    * Operation: delete-shape (delete the shape)
    * Operation: create-shape
      * Create a new shape, it’s points are listed in the last column.
    * Operation: add-point
      * Append the point onto the list of points for that shape.
    * Operation: delete-point
      * The last column is the index of the point to delete (zero indexed.) So, if the last column is a 3, find the shape in the database, and delete the fourth point.
  * The second column is the id
    * For add and delete, this is the shape to modify.
    * For create-shape, this should be the id of the new shape.
 * To make things easier, you can safely assume that there’s no more than one operation per shape (meaning: ids are distinct within operations.data)

### Your task
1. Read in the database.
2. Print (to the console) the counts of: triangles, squares, pentagons, hexagons, and heptagons (nothing pretty needed, just output something).
    * Where a triangle has 3 points, square has 4 points, pemtagon has 5 points, hexagon - 6 points, heptagon - 7 points.
3. Perform the operations against the data.
4. Repeat step 2 – meaning reprint the count of each shape.
5. Print the updated database to a file like “updated.db”.
6. Send us your code and the updated db.
