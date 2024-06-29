package com.vibhordubey333.spring.spring_boot_lessons.game;


public class GameRunner {
    //private MarioGame game;
    //Below we're creating object of interface and games which are implementing it can be executed without making change to GameRunner class.
    private GamingConsole game;
    public GameRunner(GamingConsole game) {
        this.game = game;
    }

    public void run() {
        game.down();
        game.left();
        game.right();
        game.up();
    }
}
