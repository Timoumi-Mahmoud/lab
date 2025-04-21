--- Select all rows of cd.facilities
--
SELECT * from cd.facilities;

------------------------------------------------------------------------------------------------------------------------------------------
---- list of facilities that charge a fee to members
--
select facid, name, membercost , guestcost, initialoutlay from cd.facilities where membercost > 0;

------------------------------------------------------------------------------------------------------------------------------------------
---- list of facilities that charge a fee to members, and that fee is less than 1/50th of the monthly maintenance cost
--
select facid, name, membercost , monthlymaintenance from cd.facilities where membercost > 0 and membercost <(monthlymaintenance/50) ;

------------------------------------------------------------------------------------------------------------------------------------------
---- list of all facilities with the word 'Tennis' in their name 
--
select facid, name, membercost ,guestcost , initialoutlay,monthlymaintenance from cd.facilities where name like '%Tennis%';

------------------------------------------------------------------------------------------------------------------------------------------
---- retrieve the details of facilities with ID 1 and 5
--
select facid, name, membercost ,guestcost , initialoutlay,monthlymaintenance from cd.facilities where facid in(1,5);

------------------------------------------------------------------------------------------------------------------------------------------
---- list of facilities, with each labelled as 'cheap' or 'expensive' depending on if their monthly maintenance cost is more than $100
--
select name ,monthlymaintenance ,
case 
when monthlymaintenance >100 then 'expensive'
ELSE 'cheap'
END as cost
from cd.facilities;

------------------------------------------------------------------------------------------------------------------------------------------
---- list of members who joined after the start of September 2012---
--
select memid, surname, firstname, joindate from cd.members where joindate > '2012-09-01';

------------------------------------------------------------------------------------------------------------------------------------------
---- produce an ordered list of the first 10 surnames in the members table? The list must not contain duplicates. 
--

select distinct surname from cd.members  order by surname limit 10;
 
------------------------------------------------------------------------------------------------------------------------------------------
---- combined list of all surnames and all facility names
--
select surname from cd.members union select name from cd.facilities;

------------------------------------------------------------------------------------------------------------------------------------------
----  get the signup date of your last member
--
select joindate as latest from cd.members order by joindate desc limit 1;
select firstname, surname, joindate  from cd.members order by joindate desc limit 1;

