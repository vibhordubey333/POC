package com.vibhordubey333.spring.enterprise.data;

import org.springframework.stereotype.Component;

import java.util.Arrays;
import java.util.List;

//Responsible for providing data.
@Component
public class DataService {
    public List<Integer> getData() {
        System.out.println("Hello");
        return Arrays.asList(10, 20, 30, 40);
    }
}