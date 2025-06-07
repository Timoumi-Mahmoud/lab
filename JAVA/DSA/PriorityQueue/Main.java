import java.util.PriorityQueue;
import java.util.Queue;
import java.util.Collections;

public class Main{
  public static void main(String[] args){
    System.out.println("Rolling ....");

    Queue<Double>queue = new PriorityQueue<Double>(Collections.reverseOrder());

    queue.offer(10.5);
    queue.offer(18.5);
    queue.offer(5.69);

    System.out.println("isEmpty ...."+ queue.isEmpty());
    System.out.println("contains  ...."+ queue.contains(10.5));
    printQueue(queue);
    Double removedValue=   queue.poll();
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
  public static void printQueue(Queue<Double> queue){
    System.out.println();
    while (!queue.isEmpty()){
      System.out.println(queue.poll());
    }
    System.out.println();
    System.out.println();
  }

}


