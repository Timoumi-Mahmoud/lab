/*
 * Threads (with runnable)
 */

import java.util.Scanner;

public  class Th{
		public static void main(String[] args){
				System.out.println("Rolling ......");
				Scanner scanner = new Scanner(System.in);


				MyRunnable myRun = new MyRunnable();
				Thread thread = new Thread(myRun);
				thread.setDaemon(true);
				thread.start();

				System.out.println("You have 5 second to Enter your fucking name ");

				System.out.print("Enter your name: ");
				String name = scanner.nextLine();
				System.out.print("hello "+ name);

				scanner.close();
		}
}

public class MyRunnable implements Runnable{
		@Override
		public void run(){
				for(int i =1 ; i<=10; i++){
						try{
								Thread.sleep(1000);
						}
						catch(InterruptedException e){
								System.out.println("Thread was interrupted");
						}
						if (i ==10){
								System.out.print("Timeout sorry bitch");
								System.exit(0);
						}
				}
		}
}
