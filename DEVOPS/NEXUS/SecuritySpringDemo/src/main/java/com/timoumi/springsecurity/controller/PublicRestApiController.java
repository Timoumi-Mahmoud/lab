package com.timoumi.springsecurity.controller;

import org.springframework.beans.factory.annotation.Autowired;
import org.springframework.web.bind.annotation.CrossOrigin;
import org.springframework.web.bind.annotation.GetMapping;
import org.springframework.web.bind.annotation.RequestMapping;
import org.springframework.web.bind.annotation.RestController;

import java.util.List;

@RestController
@RequestMapping("api/public")
@CrossOrigin
public class PublicRestApiController {
@Autowired
    public PublicRestApiController(){}
    @GetMapping("test")
    public String test1(){
        return "API Test ";
    }

    //available to managers
    @GetMapping("management/reports")
    public String test2(){
        return "Some report data ";
    }

}
