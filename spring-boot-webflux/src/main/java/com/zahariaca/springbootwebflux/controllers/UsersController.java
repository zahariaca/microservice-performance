package com.zahariaca.springbootwebflux.controllers;

import com.zahariaca.springbootwebflux.dto.UserDto;
import com.zahariaca.springbootwebflux.mappers.UserMapper;
import com.zahariaca.springbootwebflux.services.UserService;
import lombok.RequiredArgsConstructor;
import org.springframework.web.bind.annotation.PostMapping;
import org.springframework.web.bind.annotation.RequestBody;
import org.springframework.web.bind.annotation.RestController;
import reactor.core.publisher.Mono;

@RestController
@RequiredArgsConstructor
public class UsersController {
    private final UserMapper mapper;
    private final UserService userService;

    @PostMapping("/add")
    public Mono<UserDto> add(@RequestBody UserDto dto) {
        return Mono.just(dto)
                .map(mapper::dtoToModel)
                .flatMap(userService::save)
                .map(mapper::modelToDto);
    }
}
