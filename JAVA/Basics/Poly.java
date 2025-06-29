/*
 * Arrays
 * Scanner
 * Polymorphism
 */
import java.util.Arrays;
import java.util.Scanner;

public class Poly{

  public static void main(String[] args){
    System.out.println("rolling");
    Friend f1 = new Friend("Bara");
    f1.printNumOfFriends();
    Friend f2 = new Friend("Salim");

    Friend f3 = new Friend("Khadija");
    f3.printNumOfFriends();

    System.out.println("numOfFriends from class: "+ Friend.numOfFriends);
    Friend.stupiedMessage();

    System.out.println();
    Person p1= new Person("ali", "klay");
    Person p2= new Person("mike", "tyson");
    p1.showName();
    p2.showName();

    Student s1 = new Student("bara", "barrrr", 10.5);
    Student s2 = new Student("salim", "salimmm", 15.5);
    s1.showName();
    s2.showName();

    System.out.println();
    Car c= new Car();
    c.go();

    System.out.println();
    Bike b= new Bike();
    b.go();

    System.out.println();
    Boat boat= new Boat();
    boat.go();

    System.out.println();
    Vehicle sisi= new Boat();
    sisi.go();

    System.out.println();
    System.out.println();
    System.out.println();
    Vehicle[] vehicleArray= {c, b, boat, sisi}; 
    for (Vehicle v : vehicleArray){
      v.go();
    }

    System.out.println();
    // Shape sh = new Shape(); 
    //Poly.java:56: error: Shape is abstract; cannot be instantiated
    Triangle tr = new Triangle(); 
    tr.print();
    tr.display();

    System.out.println();
    System.out.println();
    Animal animal;
    Scanner scanner = new Scanner(System.in);
    System.out.print("What you want (Dog/Cat): ");
    String choice = scanner.next();

    if (choice.equals("Cat")){
      animal = new Cat() ;
      animal.speek();
    }else if (choice.equals("Dog")){

      animal = new Dog() ;
      animal.speek();
    }else{

      System.out.print("verify you input man !!!!!");
    }

    scanner.close();



  }

}

public abstract class Animal{
  abstract void speek();
}
public class Dog extends Animal{
  @Override
  void speek(){

    System.out.println("The dog goes **woff**");
  }
}

public class Cat extends Animal{
  @Override
  void speek(){
    System.out.println("The cat goes **meow**");
  }
}









/*
   public abstract class Vehicle{
   abstract void go();
   }
   */
public interface Vehicle{
  void go();
}
public class Car implements Vehicle{
  @Override
  public void go(){
    System.out.println("You drive the car");
  }
}

public class Bike implements Vehicle{
  @Override
  public void go(){
    System.out.println("You drive a bike");
  }
}

public class Boat implements Vehicle{
  @Override
  public void go(){
    System.out.println("You drive the boat");
  }
}


public class Friend{
  static int numOfFriends= 0;
  String name;
  Friend(String name){
    this.name= name;
    numOfFriends ++;
  }
  void printNumOfFriends(){
    System.out.println("numOfFriends: "+ numOfFriends);
  }
  static void stupiedMessage(){
    System.out.println("I don't know shit !!!");
  }

}


public class Person{
  String first;
  String last;

  Person(String first, String last){
    this.first= first;
    this.last= last;
  }

  void showName(){
    System.out.println(this.first+ "  " + this.last);
  }

}


public class Student extends Person{

  double gpa;
  Student(String first, String last, double gpa){
    super(first, last);
    this.gpa= gpa;
  }


}


public abstract class Shape{
  abstract void print(); // ABSTRACT
  void display(){  //CONCRETE
    System.out.println("This is a shape !!!!!!!!!!!");
  }
}


public class Triangle extends Shape{
  @Override
  void print(){
    System.out.println("This a Triangle");
  }
}

public class Circle extends Shape{
  @Override
  void print(){
    System.out.println("This a Circle");
  }
}

public class Rectangle extends Shape{
  @Override
  void print(){
    System.out.println("This a Rectangle");
  }

}





