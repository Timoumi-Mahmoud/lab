import java.util.Stack;
import java.util.EmptyStackException;

public class Main{
  public static void main(String[] args){
    System.out.println("Rolling ....");

    Stack<String>stack = new Stack<String>();
    stack.push("Mahmoud");
    stack.push("Bara");
    stack.push("Salim");

    printStack(stack);
    stack.pop();
    stack.pop();

    System.out.println("peek ...."+ stack.peek());
    System.out.println("search ...."+ stack.search("Bara"));

    System.out.println("is empty ...."+ stack.empty());

    stack.pop();
    try{
      stack.pop();
    }
    catch(EmptyStackException e ){
      System.out.println("Exception: trying to remove from an empty stack");
    }
    printStack(stack);
  }
  public static void printStack(Stack<String> stack){
    System.out.println();
    for (String item: stack){
      System.out.println("elements :"+ item);
    }
  }

}


