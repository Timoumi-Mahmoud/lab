import java.util.LinkedList;
import java.util.Queue;

public class Main{
  public static void main(String[] args){
    System.out.println("Rolling ....");

    Queue<String>queue = new LinkedList<String>();
    //    Queue<String>queue = new Queue<String>(); //Queue is interface not a class so it should be either LinkedList or PriorityQueue

    queue.offer("Mahmoud");
    queue.add("Bara");
    queue.add("Khadija");

    System.out.println("isEmpty ...."+ queue.isEmpty());
    System.out.println("contains  ...."+ queue.contains("Bara"));
    printQueue(queue);
    String removedValue=   queue.poll();
    queue.remove();
    queue.poll();
    queue.poll();
    queue.poll();
    queue.poll();
    queue.poll();

    System.out.println("removedValue ...."+ removedValue);
    System.out.println("queue ...."+ queue);
    System.out.println("isEmpty ...."+ queue.isEmpty());
    System.out.println("peek ...."+ queue.peek());
//    System.out.println("peek ...."+ queue.element()); cause an excepiton
    System.out.println("size ...."+ queue.size());

  }
  public static void printQueue(Queue<String> queue){
    System.out.println();
    for (String item: queue){
      System.out.println("elements :"+ item);
    }
  }

}


