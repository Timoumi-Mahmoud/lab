import java.util.ArrayList;
import java.util.LinkedList;

public class Main{
  public static void main(String[] args){
    System.out.println("Rolling ....");

    // default capcity is 10
    ArrayList<String>tab = new ArrayList<String>();

    System.out.println("isEmpty: "+ tab.isEmpty());
    tab.add("mahmoud");
    tab.add("Saitama");
    tab.add("Bang");
    tab.add("Garou");

    System.out.println("tab: "+ tab.isEmpty());

    tab.remove("Banig");

    System.out.println("tab: "+ tab);
    System.out.println();
    //    DynamicArray dynamicArray = new DynamicArray();
    //    System.out.println("dynamicArray: "+ dynamicArray.capacity);

    System.out.println();
    System.out.println();
    System.out.println();
    ArrayList<Integer> arrayList= new ArrayList<Integer>();
    LinkedList<Integer>linkedList = new LinkedList<Integer>();
    long startTime;
    long endTime;
    long elapsedTime;
    for(int i=0; i<1000000; i++){
      linkedList.add(i);
      arrayList.add(i);
    }
    startTime= System.nanoTime();
//    linkedList.get(0);
//    linkedList.get(50000);
//    linkedList.get(999999);
//    linkedList.remove(0);
    //linkedList.remove(50000);
    linkedList.remove(999999);
//    linkedList.add(1000001);
    endTime= System.nanoTime();
    elapsedTime = endTime - startTime;

    System.out.println("LinkedList : "+ elapsedTime +" ns");

    startTime= System.nanoTime();
//    arrayList.get(0);
//    arrayList.get(50000);
//    arrayList.get(999999);
    //arrayList.remove(0);
    //arrayList.remove(50000);
    arrayList.remove(999999);
    endTime= System.nanoTime();
    elapsedTime = endTime - startTime;

    System.out.println("ArrayList: "+ elapsedTime +" ns");

  }
}

//public class DynamicArray{
//  int size ;
//  int capacity=10;
//  Object[] array;
//
//  public DynamicArray(){
//    this.array= new Object[capacity];
//  }
//
//  public DynamicArray(int capacity){
//    this.capacity=capacity;
//    this.array= new Object[capacity];
//  }
//
//public void add(Object data){
//
//
//  }
//public void insert(int index,Object data){
//  if (size >= capacity){
//    grow();
//  }
//
//  }
//public void delete(Object data){
//  }
//
//public int search(Object data){
//  return -1;
//
//  }
//public void grow(Object data){
//
//  }
//public void shrink(Object data){
//
//  }
//
//public String toString(Object data){
//  return null;
//
//  }
//public boolean isEmpty(){
//  return size == 0;
//  }
//}
//
//
