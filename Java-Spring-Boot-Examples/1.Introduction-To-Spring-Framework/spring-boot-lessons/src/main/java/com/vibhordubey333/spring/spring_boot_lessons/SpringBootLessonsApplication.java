package com.vibhordubey333.spring.spring_boot_lessons;

import com.vibhordubey333.spring.spring_boot_lessons.game.GameRunner;
import com.vibhordubey333.spring.spring_boot_lessons.game.MarioGame;
import org.springframework.boot.autoconfigure.SpringBootApplication;

@SpringBootApplication
public class SpringBootLessonsApplication {

	public static void main(String[] args) {
		//SpringApplication.run(SpringBootLessonsApplication.class, args);
		MarioGame game = new MarioGame();
		GameRunner runner = new GameRunner(game);
		runner.run();
	}

}
