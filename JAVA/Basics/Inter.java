/*
 * Enums
 * HashMap
 * ArrayList
 * Timer / TimerTask
 * Interfaces
 * Anonymous function
 * Exception
 */



import java.util.Arrays;
import java.util.Scanner;
import java.util.Timer;
import java.util.TimerTask;
import java.util.ArrayList;
import java.util.Collections;
import java.util.HashMap;
import java.util.LocalDate;

public class Inter{

		public static void main(String[] args){

				Rabbit rb = new Rabbit();
				rb.flee();
				Hawk hk = new Hawk();
				hk.hunt();

				Fish shark = new Fish();
				Fish sardine = new Fish();

				shark.hunt();
				sardine.flee();

				Timer timer = new Timer();
				TimerTask task = new TimerTask(){
						int count =2;

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

				timer.schedule(task, 0, 1000);


				ArrayList<Integer> list= new ArrayList();
				list.add(99);
				list.add(69);
				list.add(96);

				System.out.println("list: "+ list);


				ArrayList<String> fruits= new ArrayList();
				fruits.add("Apple");
				fruits.add("Orange");
				fruits.add("Banana");
				fruits.add("Cocount");
				System.out.println("fruits: "+ fruits);
				fruits.set(3,"Kiwi");
				fruits.remove(0);
				System.out.println("fruits: "+ fruits);

				System.out.println("fruits: "+ fruits.get(fruits.size() -1));
				Collections.sort(fruits);

				System.out.println("fruits after sort: "+ fruits);


				HashMap<String, Double> hm= new HashMap();

				hm.put("apple", 0.95);
				hm.put("banana", 5.5);
				hm.put("orange", 0.65);
				hm.put("orange", 100000.555);
				hm.remove("apple");
				System.out.println("hm: keySet "+ hm.keySet());
				System.out.println("hm.get(key): "+ hm.get("orange"));
				System.out.println("hm.containsKey(key): "+ hm.containsKey("banana"));
				System.out.println("hm.containsValue(value): "+ hm.containsValue(5.5));
				for(String key : hm.keySet()){
						System.out.println(key + " : $" + hm.get(key));
				}



				System.out.println();
				System.out.println();

				Dog dog1 = new Dog();
				dog1.speek();

				Dog dog2 = new Dog(){
						@Override
						void speek(){
								System.out.println("Dog goes ***HAB HAB HABBBBB HABBB ***");
						}
				};

				dog2.speek();

				System.out.println();
				Scanner scDay = new Scanner(System.in);
				System.out.print("Enter a day man: ");
				String response = scDay.nextLine().toUpperCase();

				try{
						Day d = Day.valueOf(response);

						System.out.println(d);
						System.out.println(d.getDayNumber());

						switch(d){
								case SATURDAY, SUNDAY -> System.out.println("It is a weekend");
								case MONDAY, TUESDAY, THURSDAY, FRIDAY, WEDNESDAY ->  System.out.println("It is a work day");
						}
				}
				catch(IllegalArgumentException e){
						System.out.println("Please enter a valid day");
				}


				
				try{
				System.out.print("Sum " + (4/0));
				}
				catch(ArithmeticException e){
						System.out.println("unable to divide by zero");
				}

						System.out.println();
						LocalDate ld = LocalDate.now();
						System.out.println(ld);

		}

}

public enum Day{
		SUNDAY(1),
		MONDAY(2),
		TUESDAY(3),
		WEDNESDAY(4),
		THURSDAY(5),
		FRIDAY(6),
		SATURDAY(7);

		private final int dayNumber;

		Day(int dayNumber){
				this.dayNumber= dayNumber;
		}

		public int getDayNumber(){
				return this.dayNumber;
		}


}
public interface Predator{
		void hunt();
}

public interface Prey{
		void flee();
}

public class Rabbit implements Prey{
		@Override
		public void flee(){
				System.out.println("****The rabbit is running ");
		}
}
public class Hawk implements Predator{

		@Override
		public void hunt(){
				System.out.println("!!!!!! The hawk is chasing you *******");
		}
}

public class Fish implements Prey, Predator{
		@Override
		public void flee(){
				System.out.println("****The fish is running ");
		}
		@Override
		public void hunt(){
				System.out.println("!!!!!! The fish is chasing you *******");
		}

}

public class Dog{
		void speek(){
				System.out.println("Dog goes ***wooff****");
		}

}


