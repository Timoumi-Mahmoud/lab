
import java.util.Scanner;
import java.util.Random;

public class Guess{
		public static void main(String[] args){
				Scanner scanner = new Scanner(System.in);
				Random random = new Random();

				int answer;
				int attempts=10;
				int randomNumber = random.nextInt(1,11);

				System.out.println("answer: " + randomNumber);

				while (attempts > 0) {
						System.out.println("your guess bitch: "+ attempts);
						answer= scanner.nextInt();
						if (answer == randomNumber){
						System.out.println("you got it bitch");
						break;
						}else if(answer<randomNumber){
						System.out.println("answer < randomNumbe");
						} else{
						System.out.println("answer > randomNumbe");
						}
						attempts --;

				}
		}

}
