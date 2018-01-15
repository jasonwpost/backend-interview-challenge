Panono - www.panono.com
Backend programming challenge

In short: just perform the operation against the database, and save the result to a text file.

h3. General rules:
* You can use any programming language you’d like.
* You can use the Internet as much as you’d like, except don’t look at someone else’s solution to this same problem (this is a custom problem, so you might not find that anyways.)
* We don’t care much about performance, as long as you’re not doing anything really stupid.

You need two files:
* shapes.db
* operations.data

h3. [shapes.db](shapes.db)
It’s the database, stored as a file.

There’s two columns, separated by a pipe “|” symbol.
* The first column represents the id of the row.
* The second column represents points on a 2D coordinate plane (you can safely assume that each number in the point is less than 32bits in size.)
  * Points are separated by a semi-colon “;”
  * Where the x position is before the comma, y position is after it.
  * The points themselves having no real meaning or sort order.
  * There’s at most 7 points for a given shape, and minimum 3.


h3. [operations.data](operations.data)
It’s the operations to be performed against the database.

* The amount of columns are variable depending on the operation.
  * The second column is the id
    * For add and delete, this represents the shape to modify.
    * For create-shape, this represents what the id of the new shape.
  * The first column is the operation to be performed
    * Operation: delete-shape (delete the shape)
    * Operation: create-shape
      * Create a new shape, it’s points are listed in the last column.
  * Operation: add-point
    * Append the point (last column) onto the list of points for that shape.
  * Operation: delete-point
    * The last column is the index of the point to delete (zero indexed.) So, if the last 
  * To make things easier, you can safely assume that there’s no more than one operation per shape (meaning: ids are distinct within operations.data)

Your task
1. Read in the ‘shapes.db’ file.
2. Print (to the console or somewhere) the counts of: triangles, squares, pentagons, hexagons, and heptagons (nothing pretty needed, just output something).
Where a triangle has 3 points, square has 4 points, etc.
The lengths of the sides of the square are not equal, we’re just calling it that, you can call it what you want, but it’s really just a polygon with 4 vertices.
3. Perform the operations against the data.
4. Repeat step 2 – meaning reprint the count of each shape.
5. Print the updated database to a file named “updated.db”.
6. Send us your code and the updated db.

Tests are optional but correction will be taken into consideration. (Meaning, we will try not to judge you for not having tests or having bad tests.)

If by chance, you find anything wrong, bugs, mis-matched results or anything else, feel free to let us know, and of course this could be attributed to your favor (in our scoring of you.)