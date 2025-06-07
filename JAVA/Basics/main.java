import java.util.Scanner;
import java.util.Random;


public class main{
		public static void main(String[] args) throws InterruptedException{
				/*
				   int a =6666;
				   int b =5555;
				   int tmp = a;
				   a=b;
				   b=tmp;
				   String foo="foo";
				   System.out.println("a: "+ a);
				   System.out.println("b: "+ b);
				   System.out.println("foo: "+ foo);
				   System.out.println();
				   Scanner sc = new Scanner(System.in);
				   System.out.print("Enter your fucking name age: ");
				   int age = sc.nextInt();
				   sc.nextLine();

				/*
				System.out.print("Enter your fucking name bitch: ");
				String name = sc.nextLine();
				System.out.print("Enter your fucking name height: ");
				double height= sc.nextDouble();
				System.out.println("Hello : "+ name);
				System.out.println("age : "+ age);
				System.out.println("height : "+ height);

				System.out.print("Are you happy with your life? (true/false): ");
				boolean happy= sc.nextBoolean();

				if (happy == true){
				System.out.println("I am very happy !!!!!");
				}else{
				System.out.println("I am very misrable !!!!!");
				}


				if (age ==18){
				System.out.println(" still loading");

				}else if(age <17){
				System.out.println("fucking child");
				}else{

				System.out.println("addult");
				}
				sc.nextLine();

				String adjective1, adjective2, noun1, noun2,verb1, adjective3;
				System.out.print("Enter an adjective (description): ");
				adjective1= sc.nextLine();

				System.out.print("Enter a noun (animal or person): ");
				noun1= sc.nextLine();

				System.out.print("Enter an adjective (description): ");
				adjective2= sc.nextLine();

				System.out.print("Enter a verb with -ing (action): ");
				verb1= sc.nextLine();

				System.out.print("Enter an adjective (description): ");
				adjective3= sc.nextLine();

				System.out.println("Today I went to a "+adjective1 + " zoo. \n In an exhibit, I saw a "+ noun1 +". \n "+ noun1 +" was "+ adjective2+ " and " +verb1 +"! \n I was "+ adjective3+"!" );

				sc.close();
				*/

				Random random = new Random();
				int num1 = random.nextInt(1,101);
				int num2 = random.nextInt(1,101);
				int num3 = random.nextInt(1,101);
				System.out.println("random: "+ num1);
				System.out.println("random: "+ num2);
				System.out.println("random: "+ num3);

				double doub = random.nextDouble(1,55);
				System.out.println("random double: "+ doub);

				boolean isHeadns = random.nextBoolean();
				System.out.println("isHead : "+ isHeadns);
				char currency ='$';

				System.out.println("currency : "+ currency	);


				String name1= "mahmoud";
				String name2= "mahmoud";

				System.out.println("equals : "+ (name1==name2)	);
				System.out.println("equals 2 : "+ (name1.equals(name2))	);
				System.out.printf("PI: %.2f ", Math.PI);
				System.out.println("E: "+ Math.E);

				double result;

				result= Math.pow(2, 5);
				int abs= Math.abs(-666);
				double sqrt = Math.sqrt(9);

				double	round= Math.round(3.14);
				double roundCeil= Math.ceil(3.14);
				double roundFloor= Math.floor(3.14);
				System.out.println("powser: "+ result);
				System.out.println("abs: "+ abs);
				System.out.println("sqrt: "+ sqrt);
				System.out.println("round: "+ round);
				System.out.println("roundCeil: "+ roundCeil);
				System.out.println("roundFloor: "+ roundFloor);
				System.out.println(Math.min(6, 6666));
				System.out.println(Math.max(6, 6666));

				String quote= "fuck life";
				System.out.println(quote.length());
				System.out.println(quote.charAt(0));
				System.out.println(quote.indexOf("f"));
				System.out.println(quote.lastIndexOf("f"));
				System.out.println(quote.toLowerCase());
				System.out.println(quote.toUpperCase());
				System.out.println(quote.trim());
				System.out.println(quote.replace("f", "$"));
				System.out.println(quote.isEmpty());
				System.out.println(quote.contains("r"));


				int dumpValue =99;


				String test=	(dumpValue > 100)  ? "PASS": "FAIL"		;
				System.out.println(test);


				String namef="Spongebob";
				char firstLetter = namef.charAt(0);

				int age =30;
				double heightf= 60.5;
				boolean isEmployed = true;

				
				System.out.printf("hello %s\n", namef);
				System.out.printf("your name starts with a %c\n ", firstLetter);
				System.out.printf("your age is %04d\n", age);
				System.out.printf("your height is %.2f\n", heightf);
				System.out.printf("Employed : %b\n", isEmployed);

				String email= "mahmoud.timoumi@gmail.com";
				System.out.println(email.substring(0,8));

				System.out.println(email.substring(0,(email.indexOf("@"))));
				System.out.println(email.substring((email.indexOf("@")+1)));


				System.out.println();
				System.out.println();
				System.out.println("Weight converter");
				

				String day = "Tuesday";
				switch(day){
						case "Monday", "Tuesday" -> System.out.println("Back to the shit");
						case "Wednesday" -> System.out.println("count day");
						case "Thursday" -> System.out.println("-1 count");
						case "Friday" -> System.out.println("happy day");

				}


				System.out.println();
				System.out.println();

				/*
				Scanner sc = new Scanner(System.in);
				String val ="";
				while (val.isEmpty()){
					System.out.println("Enter your fucking name Bitch !");
					val = sc.nextLine();
				}
					System.out.println("Hello fukcing "+ val);
					sc.close();

					*/
				System.out.println();
				System.out.println();



				Scanner sccc = new Scanner(System.in);
				String responseVal ="";
				while (!responseVal.equals("Q")){
					System.out.println("You ara playing a game");
					System.out.println("Press Q to quit");

					responseVal = sccc.next().toUpperCase();
				}

					System.out.println("responseVal: "+ responseVal );
					sccc.close();

					for (int i= 0;i<10; i++){
							if (i==5){
									continue;
							}

							Thread.sleep(1000);
					System.out.println("i: "+ i );
					}

		}
}


		/*
		 * JDK: Java Development Kit, 
		 * contains a compilter who will convert a source code into byte code
		 */



		/* import java.util.Scanner;
		 * Scanner  scanner = new Scanner(System.in)
		 * scanner.nextLine() ==> read a string of character including character
		 * scanner.next() ==> read a string of single realted character character
		 * scanner.NextInt()==> read an integer
		 * scanner.NextDouble()==> read an double
		 * scanner.NextBoolean()==> read a boolean
		 * scanner.next.charAt(0)==> read a single char
		 *
		 */

		/* import java.util.Random;
		 *
		 Random random = new Random();
		 int num = random.nextInt(1,6);
		 *
		 double doub = random.nextDouble(); [0,1]
		 * 
		 */



		/*
		 * Variable: a resuable container for a value
		 a variable behaves as if it was the value it contains.
		 *
		 *Primitive: simple values stored directly in memory (stack)
		 int, double, char, boolean
		 *Reference: memory address (stack) that points to the (heap)
		 string, array, object
		 *
		 *
		 *
		 *
		 *
		 */

/*
 * condition ? isTrue : isFalse;
 */

/*
 *%[flags][width][.precision][specifier-character]
 + = output a plus
 , = comma grouping sparator
 (= negative numbers are enclosed in ()
 space = dispaly a minus if negative, space if postive

 for int: 0= sero padding,
		  number= right justified padding
		  negative number = left justified padding
 */
/*
 * .substring() = A method used to extract a portion of a string
 * string.substring(start, end)
 *
 */
/*
 *
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
 * /
*/

/*
 * Object= An entity that holds data(attributes) and can perform actions (methods) It is reference data type.
 *
 *
 * toString()= Method inherited from the Object class.
 * used to return a string representation of an object.
 * By default, it returns a has code as a unique identifier. It can be overridden to provide meaningful details.



 * a constructor a special method to initialize objects , you can pass arguments to a constructor and set up initial values
 *
 *
 *
 * Inheritence = One class inherits the attributes and methods from another class. child <- parent.
 * super = Refers to the parent class (subclass <- superclass). Used in constructors and method overriding . Calls the parrent constructor to initialize attributes.
 *
 * Static = a key word that Makes a variable or method belong to the class rather than to any object. Commonly used for utility methods or shared resources. for example Math.round(69.69)
 *
 */ 
 /* method overriding = when a subclass provides its own implementation of a method that is already defined.
  * Allows for code re usability and give specific implementations.
  *
  *
  *
  *
  */
 /* overloaded methods = methods that share the same name. But different parameters signature = name+ parameters
  * method can have the same name but not the same signature.
  *
    */
/* overloaded constructors = Allow a class to have multiple constructors with different parameter lists. Enable objects to be initialized in various ways.
 *
 */


/*
 *
 * Polymorphism = POLY = MANY
 *				  MORPH = SHAPE 
 *		Objects can identify as other objects.
 *		Objects can be treated as objects of a common superclass.
 *		*
 *		Dog can be indentify as Animal as well as Organism as well as Object
 *		Dog <- Animal <- Organism <- Object
 *		can be used with class as well with interfaces.
 *
 *
 *
 */


/*
 * abstract = used to define abstract classes and methods.
 * Abstraction is the process of hiding implementation details
 * and showing only the essential features;
 * Abstract classes CAN'T be iinstantiated directly
 * Can contain 'abstract' methods (which must be implemented) 
 * Can contain 'concrete' methods (which are inherited)
 *
 *
 */

 
