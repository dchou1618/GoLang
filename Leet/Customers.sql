/* query runs in 246 ms (faster than 84.12%)
# NOT IN
# SELECT Customers.Name as "customers"
# FROM Customers WHERE customers.id not in
# (
#     select CustomerId FROM Orders
# )

*/

/* JOIN - all customer names whose ids are not in orders (table2) but in table1
*/
SELECT Name AS "Customers" FROM Customers
LEFT JOIN Orders
ON Orders.CustomerId = Customers.Id
WHERE Orders.CustomerId IS NULL


/* Selecting pairs of people
"Deleting a person"
*/

DELETE p1 FROM Person p1, Person p2
WHERE p1.Email = p2.Email AND p1.Id > p2.Id

/* Deleting all duplicates

DELETE p1 FROM Person p1, Person p2
WHERE p1.Email = p2.Email

*/

DELETE FROM Person
WHERE Id NOT IN  /* if the id is not greater than the maximum per email*/
(
    SELECT MIN(Id) FROM (SELECT * FROM Person) p
    GROUP BY p.Email
)

/* Write your MySQL query statement below */

SELECT
    IFNULL(
        (SELECT DISTINCT Salary
         FROM Employee
         ORDER BY Salary DESC LIMIT 1,1
         ),
        NULL) AS "SecondHighestSalary"
/* select 1 starting from the second highest position
null otherwise if
*/


/* Obtaining duplicates */

SELECT DISTINCT p1.Email FROM Person p1, Person p2
WHERE p1.Id != p2.Id AND p1.Email = p2.Email
