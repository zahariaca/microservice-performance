package com.zahariaca.springbootwebflux.repositories;

import com.zahariaca.springbootwebflux.entity.UserEntity;
import org.springframework.data.r2dbc.repository.R2dbcRepository;

public interface UsersRepository extends R2dbcRepository<UserEntity, Long> {
}
