package com.vibhordubey333.spring.spring_boot_lessons.game;

import org.springframework.stereotype.Component;

@Component
public class SuperContraGame implements GamingConsole{
    @Override
    public void up(){
        System.out.println("SuperContra Jump");
    }
    @Override
    public void down(){
        System.out.println("SuperContra Down In The Hole");
    }
    @Override
    public void left(){
        System.out.println("SuperContra Stop");
    }
    @Override
    public void right(){
        System.out.println("SuperContra Accelerate");
    }
}
