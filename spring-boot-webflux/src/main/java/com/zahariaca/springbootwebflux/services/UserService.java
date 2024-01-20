package com.zahariaca.springbootwebflux.services;

import com.zahariaca.springbootwebflux.mappers.UserMapper;
import com.zahariaca.springbootwebflux.model.UserModel;
import com.zahariaca.springbootwebflux.repositories.UsersRepository;
import lombok.RequiredArgsConstructor;
import lombok.extern.slf4j.Slf4j;
import org.springframework.security.crypto.bcrypt.BCryptPasswordEncoder;
import org.springframework.stereotype.Service;
import reactor.core.publisher.Mono;

import java.security.SecureRandom;
import java.util.UUID;

import static java.lang.String.format;

@Slf4j
@Service
@RequiredArgsConstructor
public class UserService {
    private final UsersRepository repository;
    private final UserMapper mapper;

    public Mono<UserModel> save(UserModel userModel) {
        return Mono.just(userModel)
                .map(uM -> {
                    var bCryptPasswordEncoder =
                            new BCryptPasswordEncoder(12, new SecureRandom());
                    String encodedPassword = bCryptPasswordEncoder.encode(userModel.password());
                    var changedModel = UserModel.builder()
                            .username(format("%s_%s", userModel.username(), UUID.randomUUID()))
                            .password(encodedPassword)
                            .build();
                    return mapper.modelToEntity(changedModel);
                })
                .flatMap(repository::save)
                .map(mapper::entityToModel);
    }
}
