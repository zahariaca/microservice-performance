package com.zahariaca.springbootwebflux.model;


import lombok.Builder;

@Builder
public record UserModel(
        String username,
        String password) {
}
