package com.vibhordubey333.spring.spring_boot_lessons.game;

import org.springframework.context.annotation.Primary;
import org.springframework.stereotype.Component;

@Component
@Primary
public class MarioGame implements GamingConsole {
    public void up() {
        System.out.println("Jump");
    }

    public void down() {
        System.out.println("Down In The Hole");
    }

    public void left() {
        System.out.println("Stop");
    }

    public void right() {
        System.out.println("Accelerate");
    }
}
