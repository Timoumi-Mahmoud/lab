/*
 * Read and Write files
 * Generics
 */

import java.io.*;

public  class Read{
	public static void main(String[] args){
		String filePath= "C:\\home\\TIMOUMI\\java\\test.txt";

		try(FileWriter writer= new FileWriter(filePath)){
			String  textContent= """
     	Hell is a lonely place man.
     	Hell is a lonely place man.
     	
					""";
			writer.write(textContent);
		}
		catch (FileNotFoundException e){
			System.out.println("Could not found the file location");
		}
		catch(IOException e){
			System.out.println("Clould not write file");
		}





		try(BufferedReader reader = new BufferedReader(new FileReader(filePath))){
			String line;
			while((line = reader.readLine()) != null){
				System.out.println(line);
			}

			System.out.println("that file exists");
		}
		catch(FileNotFoundException e){
			System.out.println("Could not found the file");
		}
		catch(IOException e){
			System.out.println("Somthing went wrong");
		}

		Box<String> box = new Box();
		box.setItem("banana");
		System.out.println("Box: "+ box.getItem());
	
		Box<Integer> box1 = new Box();
		box1.setItem(696999);
		System.out.println("Box: "+ box1.getItem());
		System.out.println();
		System.out.println();
		System.out.println();
		Product<String, Double> prod = new Product<>("Shoes", 80.00);

		System.out.println("product: "+ prod.item+ " :$"+prod.price);
		
		System.out.println();
		Product<String, Integer> prodTw = new Product<>("Shoes", 69);
		System.out.println("productTwo : "+ prodTw.item+ " :$"+prodTw.price);
	}

}

public class Box<T>{
		T item;
		public void setItem(T item){
				this.item= item;
		}
		public T getItem(){
				return this.item;
		}
}


public class Product<T, U>{
		T item;
		U price;

		Product(T item, U price){
				this.item = item;
				this.price= price;
		}
}
