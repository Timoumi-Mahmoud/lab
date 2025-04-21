---- list of the start times for bookings by members named 'David Farrell'
--
select b.starttime from  cd.bookings b join  cd.members m on b.memid = m.memid where concat (m.firstname, ' ', m.surname)='David Farrell' ;

-- Another Approach
--select b.starttime 
--from cd.bookings b 
--inner join
--cd.members m 
--on b.memid= m.memid where m.firstname ='David' and m.surname='Farrell';
------------------------------------------------------------------------------------------------------------------------------------------

---- list of the start times for bookings for tennis courts, for the date '2012-09-21'
--
select b.starttime , f.name from cd.bookings b join cd.facilities f on b.facid= f.facid where f.name like 'Tennis Court%' and b.starttime >= '2012-09-21' and b.starttime < '2012-09-22' order by b.starttime;

------------------------------------------------------------------------------------------------------------------------------------------
---- output a list of all members who have recommended another member
--
select  distinct p.firstname,  p.surname from cd.members p full join cd.members c on p.memid=c.recommendedby order by p.firstname , p.surname desc;

------------------------------------------------------------------------------------------------------------------------------------------
----  list of all members, including the individual who recommended them (if any)
--

-- Inner join take a left and right table, and look for matching rows based on a join condition (on). When the condtion is satisfied, a joined row is producted.

-- A left Join operated similarly, except that if a given row on the left hand table does'nt match anything, it still produces an outputrow. That output row consists of the left hand table row and a bunch of NULLS in place of the right hand table row.
-- The right join is much like the left join, except that the left hand side of the expression is the on that contains the optional data.
-- Full join rarely-ised it treats both sides of the expression as optional.
-- select c.firstname, c.surname, p.firstname,  p.surname from cd.members c left join cd.members p on p.memid=c.recommendedby order by c.surname, c.firstname ;


-- list of all members who have used a tennis court
 select concat(m.firstname,' ', m.surname) as member, f.name as facility from cd.members m inner join cd.bookings b on m.memid=b.memid inner join cd.facilities f on b.facid=f.facid where f.name like 'Tennis%' group by m.firstname, m.surname, f.name order by m.firstname,m.surname;

select distinct  m.firstname||' '|| m.surname as member, f.name as facility 
from cd.members m inner join cd.bookings b on m.memid=b.memid 
inner join cd.facilities f on b.facid=f.facid 
where 
f.name  in ('Tennis Court 2','Tennis Court 1')
order by member, facility;
-- Note: concat strings two approach: || or concat
-- Note: SELECT DISTINCT a,b,c FROM t is roughly equivalent to: SELECT a,b,c FROM t GROUP BY a,b,c 

------------------------------------------------------------------------------------------------------------------------------------------
---- produce a list of bookings on the day of 2012-09-14 which will cost the member (or guest) more than $30
--
select m.firstname ||' '||m.surname as member, f.name as facility ,
case 
when m.memid =0 then
  b.slots * f.guestcost
else
  b.slots * f.membercost
end as cost
from cd.members m inner join cd.bookings b on m.memid=b.memid 
inner join cd.facilities f on b.facid=f.facid  
where b.starttime>='2012-09-14' and b.starttime <'2012-09-15' 
and (m.memid=0 and b.slots *f.guestcost > 30 or  
  m.memid !=0 and  b.slots * f.membercost > 30)
order by cost desc; 

------------------------------------------------------------------------------------------------------------------------------------------
---- list of all members, including the individual who recommended them (if any), without using any joins
--
select p.firstname ||' '||p.surname as member,(select firstname ||' '||surname  as recommender from cd.members where memid=p.recommendedby  ) from cd.members p order by member ;

-- Note: Subqueries are, queries within a query. They're commonly used with aggregates, to answer questions like 'get me all the details of the member

------------------------------------------------------------------------------------------------------------------------------------------
-- who has spent the most hours on Tennis Court1' . It can also uses inforamtion from the outer  query to emulate an outer join.

