package com.example.springbootrest.dto;


import lombok.Builder;

@Builder
public record UserDto(
        String username,
        String password) {
}
