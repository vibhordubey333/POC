package com.vibhordubey333.spring.spring_boot_lessons.game;

public class GameRunner {
    private MarioGame game;
    public GameRunner(MarioGame game) {
        this.game = game;
    }

    public void run() {
        game.down();
        game.left();
        game.right();
        game.up();
    }
}
