import java.util.ArrayList;
import java.util.Arrays;
import java.util.LinkedList;

public class Main{
  public static void main(String[] args){
    System.out.println("Rolling ....");

    int[] tab= {1, 4,7, 55, 69 }; 
    System.out.println("5444: "+ linearSearch(54444, tab));
    System.out.println("69: "+ linearSearch(69, tab));
    for (int item: tab){
      System.out.println("  "+ item);
    }

    boolean tt= 'H' < 'I';
    System.out.println("H > I"+tt );

    boolean result=binarySearch(15,tab);
    System.out.println("result of 69: "+ result);
    System.out.println();
    System.out.println();

    int array[] = new int[100];
    int target= 69;
    for (int i=0; i<array.length; i++){
      array[i]=i;
    }

    int index = Arrays.binarySearch(array, 695);
    System.out.println("index: "+ index);
  }

  public static boolean linearSearch(int target,int[] array){
    for (int item: array){
      if(item == target){
        return true;
      }
    }
    return false;
  }

  public static boolean binarySearch(int target,int[] array){
    int left =0;
    int right=array.length -1;
    while(left <= right ){
      int middle = left +  (right - left) /2  ;
      if (array[middle] > target){
        right = middle -1 ;
      }else if (array[middle] <target){
        left = middle + 1  ;
      }else  if (array[middle] == target) {
        return true;
      }
    }
    return false;
  }
}
