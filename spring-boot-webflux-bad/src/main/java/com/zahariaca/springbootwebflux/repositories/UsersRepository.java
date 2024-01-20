package com.zahariaca.springbootwebflux.repositories;

import com.zahariaca.springbootwebflux.entity.UserEntity;
import org.springframework.data.jpa.repository.JpaRepository;

public interface UsersRepository extends JpaRepository<UserEntity, Long> {
}
