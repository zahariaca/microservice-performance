package com.zahariaca.springbootwebflux.dto;


import lombok.Builder;

@Builder
public record UserDto(
        String username,
        String password) {
}
