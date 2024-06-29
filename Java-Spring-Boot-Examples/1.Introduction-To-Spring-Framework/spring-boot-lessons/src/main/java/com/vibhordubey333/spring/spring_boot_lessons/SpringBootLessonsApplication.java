package com.vibhordubey333.spring.spring_boot_lessons;

import com.vibhordubey333.spring.spring_boot_lessons.game.GameRunner;
import com.vibhordubey333.spring.spring_boot_lessons.game.MarioGame;
import com.vibhordubey333.spring.spring_boot_lessons.game.SuperContraGame;
import org.springframework.boot.autoconfigure.SpringBootApplication;

@SpringBootApplication
public class SpringBootLessonsApplication {

	public static void main(String[] args) {
		//SpringApplication.run(SpringBootLessonsApplication.class, args);
		// Tight Coupling: If we want to  execute Mario or SuperContra we need to make change in the below line.
		//MarioGame game = new MarioGame();
		SuperContraGame game = new SuperContraGame();
		GameRunner runner = new GameRunner(game);
		runner.run();
	}

}
