import java.util.LinkedList;
import java.util.Queue;

public class Main{
  public static void main(String[] args){
    System.out.println("Rolling ....");

    LinkedList<String>lList = new LinkedList<String>();

    System.out.println("isEmpty: "+ lList.isEmpty());
//    lList.add("mahmoud");
//    lList.add("Saitama");
//    lList.add("Bang");
//    lList.add("Garou");
//    lList.remove("Garou");
      lList.push("A");
      lList.push("B");
      lList.push("C");
      lList.push("D");

    System.out.println("lList: "+ lList);
lList.pop();

      lList.offer("D");
      lList.poll();
    System.out.println("lList: "+ lList);
    System.out.println("search for letter A: "+ lList.indexOf("A"));
    System.out.println("isEmpty: "+ lList.isEmpty());

    System.out.println();

    System.out.println("lList: "+ lList);
    System.out.println("lList peekFirst: "+ lList.peekFirst());
    System.out.println("lList peekLast: "+ lList.peekLast());
    lList.removeFirst();
    lList.removeLast();

    System.out.println("after remove first and last lList: "+ lList);
  }

///  public static void printLinkedList(LinkedList<String> lList){
///    System.out.println();
///    while(lList.next != null){
///      System.out.println(lList.next);
///    }
///  }

}


