package com.vibhordubey333.spring.spring_boot_lessons.game;

import org.springframework.beans.factory.annotation.Autowired;
import org.springframework.beans.factory.annotation.Qualifier;
import org.springframework.stereotype.Component;

@Component
public class GameRunner {
    //private MarioGame game;
    //Below we're creating object of interface and games which are implementing it can be executed without making change to GameRunner class.
   // @Qualifier("marioGame")
    @Autowired
    private GamingConsole game;

    public GameRunner(@Qualifier("marioGame") GamingConsole game) {
        this.game = game;
    }

    public void run() {
        game.down();
        game.left();
        game.right();
        game.up();
    }
}
