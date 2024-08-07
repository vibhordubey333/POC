package com.vibhordubey333.spring.enterprise.business;

import com.vibhordubey333.spring.enterprise.data.DataService;
import org.springframework.beans.factory.annotation.Autowired;
import org.springframework.stereotype.Component;

import java.util.List;

//Responsible for performing computation.
@Component
public class BusinessService {
    // If using setDataService comment below  @Autowired and place it above setDataService.
//    @Autowired
    private DataService dataService;
    @Autowired
    public void setDataService(DataService dataService) {
        System.out.println("setDataService:");
        this.dataService = dataService;
    }

    public long calculateSum() {
        List<Integer> data = dataService.getData();
        //Utilizing functional programming to do sum of a list.
        return data.stream().reduce(Integer::sum).get();
    }
}
