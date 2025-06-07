import java.util.ArrayList;
import java.util.LinkedList;

public class Main{
  public static void main(String[] args){
    System.out.println("Rolling ....");
    System.out.println("sum one: "+ addupOne(5));
    System.out.println("sum two: "+ addupTwo(5));
  }

  public static int addupOne(int n){  // o(n)
    int sum=0;
    for (int i=0; i<= n; i++){
      sum +=i;
    }
    return sum;
  }

  public static int addupTwo(int n){  // o(1) takes three steps
    int sum = n * (n+1) /2;
    return sum;
  }
}
