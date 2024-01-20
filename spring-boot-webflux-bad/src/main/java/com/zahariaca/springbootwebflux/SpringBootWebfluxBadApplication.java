package com.zahariaca.springbootwebflux;

import org.springframework.boot.SpringApplication;
import org.springframework.boot.autoconfigure.SpringBootApplication;
import org.springframework.web.reactive.config.EnableWebFlux;

@EnableWebFlux
@SpringBootApplication
public class SpringBootWebfluxBadApplication {

    public static void main(String[] args) {
        SpringApplication.run(SpringBootWebfluxBadApplication.class, args);
    }

}

