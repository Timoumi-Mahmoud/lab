-- DML: data manipulation language operations that alter (inserting , updating, deleting)
-- SELECT * from cd.facilities;

-- adding a new facility - a spa. 
-- insert into cd.facilities (facid , Name , membercost, guestcost, initialoutlay , monthlymaintenance )values(9, 'Spa', 20,30, 100000, 800.);

--SELECT * from cd.facilities;

-- add multiple facilities in one command
-- insert into cd.facilities (facid , Name , membercost, guestcost, initialoutlay , monthlymaintenance )values(11, 'SpaTwo', 20,30, 100000, 800.),( 10,  'Squash Court 2',  3.5,  17.5,  5000,  80);


-- adding the spa to the facilities table again. This time, though, we want to automatically generate the value for the next facid, rather than  specifying it as a constant.
-- insert into cd.facilities (facid , Name , membercost, guestcost, initialoutlay , monthlymaintenance )values(((select max(facid) from cd.facilities)+1), 'Spa', 20,30, 100000, 400);


-- alter data for the second tennis court. The initial outlay was 10000 rather than 8000

-- update cd.facilities set initialoutlay = 8000 where name='Tennis Court 2';

-- increase the price of the tennis courts for both members and guests. Update the costs to be 6 for members, and 30 for guests.

-- update  cd.facilities set membercost=6 ,  guestcost=30 where Name in ('Tennis Court 1','Tennis Court 2');
-- SELECT * from cd.facilities where  Name in ('Tennis Court 1','Tennis Court 2');



--  alter the price of the second tennis court so that it costs 10% more than the first one. Try to do this without using constant values for the prices, so that we can reuse the statement if we want to.

-- update  cd.facilities set membercost=membercost + (((select membercost from cd.facilities where Name='Tennis Court 1')*10) /100 ),  guestcost=guestcost + (((select membercost from cd.facilities where Name='Tennis Court 1')*10) /100 )where Name='Tennis Court 2';

--SELECT * from cd.facilities where  Name in ('Tennis Court 1','Tennis Court 2');


-- Delete all bookings from the cd.bookings table
delete from cd.bookings;
select * from cd.bookings;
