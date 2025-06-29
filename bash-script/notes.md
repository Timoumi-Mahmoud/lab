string          Numeric     True if
x = y           x -eq y		  x is equal to y
x != y			    x -ne y		  x is not equal to y
x < y		        x -lt y		  x is less than y
n/a				      x -le y		  x is less than or equal to y
x > y			      x -gt y		  x is greater than y
n/a			        x -ge y		  x is greater than or equal to y
-n x			      n/a			    x is not null
-z x			      n/a			    x is null

Operator
-d file				file exists and is a directory
-e file				file exists
-f file				file exists and is a regular file
-r file				User has read permission on file
-s file				file exists and is not empty
-w file				User has write permission on file
file1 -nt file2		file1 is newer than file2
file1 -ot file2		file1 is older than file2


Symbol			What it match or does
 .				  Matches any character
[chars]			Matches any character from a given set
[^chars]		Matches any character not in a given set
^				    Matche the beginning of a line
$				    Matches the end of a line
\w				  Matches any “word” character (same as [A-Za-z0-9_])
\s				  Matches any whitespace character (same as [ \f\t\n\r]) a
\d				  Matches any digit (same as [0-9])
|				    Matches either the element to its left or the one to its right
(expr)			Limits scope, groups elements, allows matches to be captured
?				    Allows zero or one match of the preceding element
*				    Allows zero, one, or many matches of the preceding element
+				    Allows one or more matches of the preceding element
{ n }			  Matches exactly n instances of the preceding element
{ min, }		Matches at least min instances (note the comma)
{ min,max }	Matches any number of instances from min to max



find . -type f -exec mv '{}' '{}'.jpg \;
