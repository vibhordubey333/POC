package com.vibhordubey333.spring.spring_boot_lessons.game;

import javax.annotation.processing.SupportedSourceVersion;

public class PacMan implements GamingConsole{

    @Override
    public void up() {
        System.out.println("PacMan up");
    }

    @Override
    public void down() {
        System.out.println("PacMan down");
    }

    @Override
    public void right() {
        System.out.println("PacMan right");
    }

    @Override
    public void left() {
        System.out.println("PacMan left");
    }
}
