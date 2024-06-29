package com.vibhordubey333.spring.spring_boot_lessons;

import com.vibhordubey333.spring.spring_boot_lessons.game.GameRunner;
import com.vibhordubey333.spring.spring_boot_lessons.game.MarioGame;
import com.vibhordubey333.spring.spring_boot_lessons.game.PacMan;
import com.vibhordubey333.spring.spring_boot_lessons.game.SuperContraGame;
import org.springframework.boot.autoconfigure.SpringBootApplication;

@SpringBootApplication
public class SpringBootLessonsApplication {

	public static void main(String[] args) {
		//SpringApplication.run(SpringBootLessonsApplication.class, args);
		// Tight Coupling: If we want to  execute Mario or SuperContra we need to make change in the below line. As we've implemented
		// the GamingConsole interface so we can uncomment the below code.
		MarioGame marioGame = new MarioGame();
		SuperContraGame superContraGame = new SuperContraGame();
		PacMan pacMan = new PacMan();

		GameRunner runnerMario = new GameRunner(marioGame);
		runnerMario.run();

		GameRunner runnerContra = new GameRunner(superContraGame);
		runnerContra.run();

		GameRunner runnerPacMan = new GameRunner(pacMan);
		runnerPacMan.run();
	}

}
