package com.vibhordubey333.spring;

import com.vibhordubey333.spring.enterprise.web.MyWebController;
import com.vibhordubey333.spring.spring_boot_lessons.game.GameRunner;
import org.springframework.boot.SpringApplication;
import org.springframework.boot.autoconfigure.SpringBootApplication;
import org.springframework.context.ConfigurableApplicationContext;

@SpringBootApplication
public class SpringBootLessonsApplication {

    public static void main(String[] args) {
        // Tight Coupling: If we want to  execute Mario or SuperContra we need to make change in the below line. As we've implemented
        // the GamingConsole interface so we can uncomment the below code.
//		MarioGame marioGame = new MarioGame();
//		SuperContraGame superContraGame = new SuperContraGame();
//		PacMan pacMan = new PacMan();
//
//		GameRunner runnerMario = new GameRunner(marioGame);
//		runnerMario.run();
//
//		GameRunner runnerContra = new GameRunner(superContraGame);
//		runnerContra.run();
//
//		GameRunner runnerPacMan = new GameRunner(pacMan);
//		runnerPacMan.run();

        // Above code is working we've commented it as we're managing dependencies using SpringBoot
        //Visit each class to see @Component is added it denotes that dependency is being managed by Spring.
        ConfigurableApplicationContext context = SpringApplication.run(SpringBootLessonsApplication.class, args);
        GameRunner runner = context.getBean(GameRunner.class);
        runner.run();

        MyWebController controller = context.getBean(MyWebController.class);
        System.out.println(controller.returnValueFromBusinessService());

    }

}
