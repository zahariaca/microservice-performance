package com.example.springbootrest.model;


import lombok.Builder;

@Builder
public record UserModel(
        String username,
        String password) {
}
