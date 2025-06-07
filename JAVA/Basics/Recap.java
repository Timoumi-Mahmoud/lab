/*
 * static
 * Arrays
 * Recap
 */  

import java.util.Arrays;
import java.util.Scanner;

public class Recap{
  static int x =99999;
  public static void main(String[] args){
    doSomething();
    System.out.println("hello there: x "+ x);
    System.out.println(sum(5,4));
    System.out.println(sum(5,4, 100));
    System.out.println();

    int[] tab ={1,2,3,4,5,6,7,8,9,10};

    tab[0]=9999;
    for (int i =0; i<tab.length; i++){
      System.out.println(tab[i]);
    }

    System.out.println();
    System.out.println();

    String[] fruits ={"banana","apple","orange","kiwi"};

    System.out.println("=================>"+ linearlinearSearch(fruits, "kiwi"));

    System.out.println(fruits.length);
    for (String fruit : fruits){
      System.out.println(fruit);
    }

    System.out.println();
    System.out.println();
    Arrays.sort(fruits);

    for (String fruit : fruits){
      System.out.println(fruit);
    }

    System.out.println();
    Arrays.fill(fruits, "cherry");

    for (String fruit : fruits){
      System.out.println(fruit);
    }

    Scanner scanner = new Scanner(System.in);
    System.out.println();
    System.out.println("What # of foods do you want ?: ");
    int size=scanner.nextInt();
    scanner.nextLine();

    String [] foods = new String[size];


    for (int i =0; i<foods.length; i++){
      System.out.println("Enter a food: ");
      foods[i]= scanner.nextLine();
    }


    scanner.close();


    System.out.println();
    System.out.println();



    for (String food : foods){
      System.out.println(food);
    }

    System.out.println(add(69,616,666));
    System.out.println(add(6,1,6));
    System.out.println(add());


    Car[] cars ={ new Car("Mustang","Red"), 
      new Car("BMW","Black"),
      new Car("Ferrari","Red")};
    //Car[] cars ={car1 , car2, car3};

    for (Car car: cars){
      car.color= "yellow";
      car.drive();
    }


    User user1= new User("SpongBob");
    User bara= new User("bara", "bara@gmail.com");
    User salim= new User("salim", "salim@gmail.com", 6);
    User u= new User();

    user1.printUser();
    bara.printUser();

    System.out.println();
    System.out.println(	bara.toString());


    Dog dog =new Dog();
    Cat cat = new Cat("alice");
    Fish fish = new Fish();
    Animal animal = new Animal();

    fish.move();
    dog.move();
    cat.move();
    animal.move();

    System.out.println();
    System.out.println();
    System.out.println(cat.toString());
    cat.move();
    cat.tt();

    System.out.println("is alive: "+ cat.isAlive);
    System.out.println("type: "+ cat.type);
  }


  static  int linearlinearSearch(String[]tab, String value){
    for (int i =0; i <tab.length;i++){
      if(tab[i].equals(value)){
        return i;
      }
    }

    return -1;
  }

  static int add(int... numbers){
    System.out.println(numbers);
    int sum=0;

    for (int num : numbers){
      sum += num;
    }
    return sum;
  }


  static int sum (int x ,int y ){
    return x+y;
  }


  static int sum (int x ,int y, int w ){
    return x+y +w ;
  }

  static void doSomething(){
    x++;
  }

}

public class Car{

  String model;
  String color;

  Car(String model, String  color){
    this.model= model;
    this.color= color;
  }

  void drive(){
    System.out.println("You drive the "+ this.color +" " + this.model);

  }
}

public class Organism{
  String type ="mamifaires";
  void tt(){
    System.out.println("living organism");
  }
}
public class Animal extends Organism{
  boolean isAlive;
  int age;
  Animal(){
    isAlive= true;
  }

  void eat(){
    System.out.println("The animal is eating");
  }

  void move(){
    System.out.println("this animal is on fire");
  }
}

public class Cat extends Animal{
  String name;
  Cat(String name){
    super();
    this.name= name;
  }
  @Override 
  public String toString(){
    return this.name + " is alive " +this.isAlive+" ";  
  }


  @Override
  void move(){
    System.out.println("this pussy is on fire");
  }
}

public class Dog extends Animal{
}

public class Fish extends Animal{

  @Override
  void move(){
    System.out.println("this fish is swimming");
  }

}

public class User{

  String username;
  String email;
  int age;

  User(String username){
    this.username= username;
    this.email= "not provided";
    this.age=0;
  }

  User(String username, String email){
    this.username= username;
    this.email= email;
    this.age=0;
  }

  User(String username, String email,int age){
    this.username= username;
    this.email= email;
    this.age=age;
  }
  User(){
    this.username= "UNKNOWN";
    this.email= "not provided";
    this.age=0;

  }
  public		void printUser(){
    System.out.println(this.username + " " +this.age+" "+ this.email );
  }

  @Override 
  public String toString(){
    return this.username + " " +this.age+" "+ this.email;  
  }


}


