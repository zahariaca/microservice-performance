package com.example.springbootrest.controllers;

import com.example.springbootrest.dto.UserDto;
import com.example.springbootrest.mappers.UserMapper;
import com.example.springbootrest.services.UserService;
import lombok.RequiredArgsConstructor;
import org.springframework.web.bind.annotation.PostMapping;
import org.springframework.web.bind.annotation.RequestBody;
import org.springframework.web.bind.annotation.RestController;

@RestController
@RequiredArgsConstructor
public class UsersController {
    private final UserMapper mapper;
    private final UserService userService;

    @PostMapping("/add")
    public UserDto add(@RequestBody UserDto dto) {
        var model = mapper.dtoToModel(dto);

        var savedModel = userService.save(model);

        return mapper.modelToDto(savedModel);
    }
}
