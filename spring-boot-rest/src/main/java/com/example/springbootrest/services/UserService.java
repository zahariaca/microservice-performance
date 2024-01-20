package com.example.springbootrest.services;

import com.example.springbootrest.dto.UserDto;
import com.example.springbootrest.mappers.UserMapper;
import com.example.springbootrest.model.UserModel;
import com.example.springbootrest.repositories.UsersRepository;
import jakarta.transaction.Transactional;
import lombok.RequiredArgsConstructor;
import lombok.extern.slf4j.Slf4j;
import org.springframework.security.crypto.bcrypt.BCryptPasswordEncoder;
import org.springframework.stereotype.Service;

import java.security.SecureRandom;
import java.util.UUID;

import static java.lang.String.format;

@Slf4j
@Service
@RequiredArgsConstructor
public class UserService {
    private final UsersRepository repository;
    private final UserMapper mapper;

    @Transactional
    public UserModel save(UserModel userModel) {
//        log.info("UserModel: {}", userModel);
        var bCryptPasswordEncoder =
                new BCryptPasswordEncoder(12, new SecureRandom());
        String encodedPassword = bCryptPasswordEncoder.encode(userModel.password());
        var changedModel = UserModel.builder()
                .username(format("%s_%s", userModel.username(), UUID.randomUUID()))
                .password(encodedPassword)
                .build();
//        log.info("ChangedModel: {}", changedModel);

        var entity = mapper.modelToEntity(changedModel);
//        log.info("Entity: {}", entity);
        var saved = repository.save(entity);
//        log.info("Saved Entity: {}", saved);

        return mapper.entityToModel(saved);
    }
}
