

/* Basic sql query */

SELECT name AS "Name", population AS "Population", area AS "Area"
FROM World
WHERE area > 3000000 OR population > 25000000
