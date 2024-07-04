package com.vibhordubey333.spring.enterprise.web;

import com.vibhordubey333.spring.enterprise.business.BusinessService;
import org.springframework.beans.factory.annotation.Autowired;
import org.springframework.stereotype.Component;

//Returning data.
@Component
public class MyWebController {
    @Autowired
    private BusinessService businessService;

    public long returnValueFromBusinessService() {
        return businessService.calculateSum();
    }
}


