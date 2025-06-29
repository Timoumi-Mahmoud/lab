## JAVA:

### Object
An entity that holds data(attributes) and can perform actions (methods) It is reference data type.
 
#### Constructor
A special method to initialize objects , you can pass arguments to a constructor and set up initial values.

#### Inheritance 
One class inherits the attributes and methods from another class. child <- parent.
super = Refers to the parent class (subclass <- superclass). Used in constructors and method overriding . Calls the parent constructor to initialize attributes.

####  Static 
A key word that Makes a variable or method belong to the class rather than to any object. Commonly used for utility methods or shared resources. For example Math.round(69.69)

#### method overriding 
When a subclass provides its own implementation of a method that is already defined.
Allows for code re usability and give specific implementations.

#### overloaded methods 
Methods that share the same name. But different parameters signature = name+ parameters.
Method can have the same name but not the same signature.

#### Overloaded constructors 
Allow a class to have multiple constructors with different parameter lists. Enable objects to be initialized in various ways.

#### Polymorphism 
POLY = MANY
MORPH = SHAPE 
Objects can identify as other objects.
Objects can be treated as objects of a common superclass.

Dog can be identify as Animal as well as Organism as well as Object
Dog <- Animal <- Organism <- Object
Can be used with class as well with interfaces.

#### Runtime Polymorphism (Dynamic Polymorphism)
When the method that gets executed is decided at runtime based on the actual type of the object.

#### abstract 
Used to define abstract classes and methods.
 * Abstraction is the process of hiding implementation details
 * and showing only the essential features;
 * Abstract classes CAN'T be iinstantiated directly
 * Can contain 'abstract' methods (which must be implemented) 
 * Can contain 'concrete' methods (which are inherited)
 
#### Interfaces
A blueprint for a class that specifies a set of abstract methods, that implementing classes Must define.
Supports multiple inhearitance-like behavior.

#### Getter Setters
They help protect object data and add rules for accessing or modifying them.
Getters: methods that make a field readable.
Setters: methods that make a field writable.

#### Composition
Represent a "part-of" relationship between objects.
For example, an Engine is "part of" a Car.
Allows complex objects to be constructed from smaller objects.


#### File reading
Three popular options:
1. BufferedReader + FileReader: Best for reading text files line-by-line.
2. FileInputStream: Best for binary files (eg., images, audio files).
3. RandomAccessFile: Best for read/write specific portions of a large file.

### Writting into files
Four popular options:
1. FileWriter : good for small or medium-sized text files
2. BufferedWriter : better performance for large amounts of text
3. PrintWriter: Best for structured data, like reports or logs
4. FileOutputStream: Best for binary files (images, audio files)

### Enum (Enumerations)
A special kind of class that represent a fixed set of constants.
They improve code readability and are easy to maintain.
More efficient with switches when comparing Strings.

### Exception
An event that interrupts the normal flow of a program (Dividing by zero, file not found, mismatch input type)
Surround any dangerous code with a try{} block 
try{}, catch{}, finally{}(will always execute weather there an exception or not)(for example scanner.close();, close a file)
Note: there is also try with resources: for example 
```Java
try(Scanner sc = new Scanner(System.in)){
....
} java will automaticlay close the scanner without closing it manualy in the finaly block 
```
ArithmeticException 
InputMismatchException
IOException
FileNotFoundException
-> Exception ```catch(Exception e){sout("somthing went wrong"); ```

-----------------------------------------------------------------------

#### Anonymous Class
A class that doesn't have a name. Cannot be resused.
Add custom behavior without having to create a new class.
Often used for one time uses (TimerStack, Runnable, callbacks)



#### Timer: 
class that schedules tasks at specific times or periodically 
Useful for: sending notifications, scheduled updated, repetitive actions

#### TimerTask: 
Represents the task that will be executed by the Timer 
You will extend the TimerTask class to define your task 
Create a subclass of TimerTask and @Override run()

```Java
				Timer timer = new Timer();
				TimerTask task = new TimerTask(){
						int count =5;

						@Override
						public void run(){
							System.out.println("rolling");
							count--;
							if (count <=0 ){
								System.out.println("Task complite");
								timer.cancel();
							}
						};

				};

				timer.schedule(task, 10000, 1000);
```

#### ArrayList
A resizable array that stores objects (autoboxing)
Array are fixed in size but ArrayList can change.
[add, remove, set, size, get]
Collections.sort(ar);

#### HashMap
A data strucgture that stores key-value pairs keys 
keys are unique, but Values can be duplicated
Does not maintain any order, but is memory efficient
HashMap<Key, Value>
				HashMap<String, Double> hm= new HashMap();
				hm.put(key, value);
				hm.remove(key);
				hm.size();
				System.out.println("hm: keySet "+ hm.keySet());
				System.out.println("hm.get(key): "+ hm.get("orange"));
				System.out.println("hm.containsKey(key):"+ hm.containsKey("orange"));
				System.out.println("hm.containsValue(value):"+ hm.containsKey(69.69));
				for(String key : hm.keySet()){
						System.out.println(key + " : $" + hm.get(key));
				}

#### Threading
Allow a program to run multiple tasks simultaneously, Helps improve performance with time-consuming operations (File I/O, network communication, or
any background tasks).

How to create a Thread: 
Option 1. Extend the Thread class (simpler)
Option 2. Implement the Runnable interface (better) (see Th.java)

#### Generics
A concept where you can write a class, interface, or method that is compatable with different data types.
<T> type parameter (placeholder that gets replaced with a real type)
<String> type argument (specifies the type)


#### Dates & Times
LocalDate, LocalTime, LocalDateTime, UTC timestamp
```Java
	LocalDate date = LocalDate.now();
	System.out.println("data: "+ date);

	LocalTime time = LocalTime.now();
	System.out.println("time: "+ time);

	LocalDateTime dateTime = LocalDateTime.now();
	DateTimeFormatter formatter = DateTimeFormatter.ofPattern("MM-dd-yyyy HH:mm:ss");
	String newDateTime = dateTime.format(formatter);
	System.out.println("dateTime: "+ newDateTime);

	Instant instant = Instant.now();
	System.out.println("instant: "+ instant);


	LocalDate aDay = LocalDate.of(2024,12,25);
	LocalDate aDayTime = LocalDate.of(2024,12,25, 12,0,0);


	/// comparing :
	date1.isBefore(date2)
	date1.isAfter(date2)
	date1.isEqual(date2)
```

#### Wrapper classes
allow primitive values (int, char, double, boolean) to be used as objects.
Generallly, don't wrap primitives unless you need an object.
Allows use of Collections Framework and static Utility Methods.

Integer a = new Integer(123);
Double, Character, Boolean
// Autoboxing
Integer a = 123;
Double b= 3.14;
Character, Boolean
// Unboxing
int x = a;
char x= c;

String a= Integer.toString(123);
int a = Integer.parseInt("123");

#### Scanner
		 import java.util.Scanner;
		 * Scanner  scanner = new Scanner(System.in)
		 * scanner.nextLine() ==> read a string of character including character
		 * scanner.next() ==> read a string of single realted character character
		 * scanner.NextInt()==> read an integer
		 * scanner.NextDouble()==> read an double
		 * scanner.NextBoolean()==> read a boolean
		 * scanner.next.charAt(0)==> read a single char
#### Random

		 import java.util.Random;
		 Random random = new Random();
		 int num = random.nextInt(1,6);
		 double doub = random.nextDouble(); [0,1]



		 * Variable: a resuable container for a value
		 a variable behaves as if it was the value it contains.
		 *
		 *Primitive: simple values stored directly in memory (stack)
		 int, double, char, boolean
		 *Reference: memory address (stack) that points to the (heap)
		 string, array, object
 * condition ? isTrue : isFalse;

### Printf
 *%[flags][width][.precision][specifier-character]
 + = output a plus
 , = comma grouping sparator
 (= negative numbers are enclosed in ()
 space = dispaly a minus if negative, space if postive

 for int: 0= sero padding,
		  number= right justified padding
		  negative number = left justified padding


  .substring() = A method used to extract a portion of a string
  string.substring(start, end)
 case (oper) {
	case '' -> ...;
	case '' -> ...;
	case '' -> ...;
	default ->....;
 }

&& and
|| or
! not
 * Break: break out of a loop (Stop)
 * continuse: skip curreent iteration of a loop (Skip)

#### toString
* toString()= Method inherited from the Object class.
 * used to return a string representation of an object.
 * By default, it returns a hash code as a unique identifier. It can be overridden to provide meaningful details.

